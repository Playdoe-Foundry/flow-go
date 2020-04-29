package eventhandler_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/dapperlabs/flow-go/consensus/hotstuff"
	"github.com/dapperlabs/flow-go/consensus/hotstuff/eventhandler"
	"github.com/dapperlabs/flow-go/consensus/hotstuff/mocks"
	"github.com/dapperlabs/flow-go/consensus/hotstuff/model"
	"github.com/dapperlabs/flow-go/consensus/hotstuff/notifications"
	"github.com/dapperlabs/flow-go/consensus/hotstuff/pacemaker"
	"github.com/dapperlabs/flow-go/consensus/hotstuff/pacemaker/timeout"
	"github.com/dapperlabs/flow-go/model/flow"
	module "github.com/dapperlabs/flow-go/module/mock"
)

const (
	startRepTimeout        float64 = 400.0 // Milliseconds
	minRepTimeout          float64 = 100.0 // Milliseconds
	voteTimeoutFraction    float64 = 0.5   // multiplicative factor
	multiplicativeIncrease float64 = 1.5   // multiplicative factor
	additiveDecrease       float64 = 50    // Milliseconds
)

// TestPaceMaker is a real pacemaker module with logging for view changes
type TestPaceMaker struct {
	*pacemaker.FlowPaceMaker
	t *testing.T
}

func NewTestPaceMaker(t *testing.T, startView uint64, timeoutController *timeout.Controller, notifier hotstuff.Consumer) *TestPaceMaker {
	p, err := pacemaker.New(startView, timeoutController, notifier)
	if err != nil {
		t.Fatal(err)
	}
	return &TestPaceMaker{
		FlowPaceMaker: p,
		t:             t,
	}
}

func (p *TestPaceMaker) UpdateCurViewWithQC(qc *model.QuorumCertificate) (*model.NewViewEvent, bool) {
	oldView := p.CurView()
	newView, changed := p.FlowPaceMaker.UpdateCurViewWithQC(qc)
	p.t.Logf("pacemaker.UpdateCurViewWithQC old view: %v, new view: %v\n", oldView, p.CurView())
	return newView, changed
}

func (p *TestPaceMaker) UpdateCurViewWithBlock(block *model.Block, isLeaderForNextView bool) (*model.NewViewEvent, bool) {
	oldView := p.CurView()
	newView, changed := p.FlowPaceMaker.UpdateCurViewWithBlock(block, isLeaderForNextView)
	p.t.Logf("pacemaker.UpdateCurViewWithBlock old view: %v, new view: %v\n", oldView, p.CurView())
	return newView, changed
}

func (p *TestPaceMaker) OnTimeout() *model.NewViewEvent {
	oldView := p.CurView()
	newView := p.FlowPaceMaker.OnTimeout()
	p.t.Logf("pacemaker.OnTimeout old view: %v, new view: %v\n", oldView, p.CurView())
	return newView
}

// using a real pacemaker for testing event handler
func initPaceMaker(t *testing.T, view uint64) hotstuff.PaceMaker {
	notifier := &mocks.Consumer{}
	tc, err := timeout.NewConfig(
		time.Duration(startRepTimeout*1e6),
		time.Duration(minRepTimeout*1e6),
		voteTimeoutFraction,
		multiplicativeIncrease,
		time.Duration(additiveDecrease*1e6))
	if err != nil {
		t.Fail()
	}
	pm := NewTestPaceMaker(t, view, timeout.NewController(tc), notifier)
	notifier.On("OnStartingTimeout", mock.Anything).Return()
	notifier.On("OnSkippedAhead", mock.Anything).Return()
	notifier.On("OnReachedTimeout", mock.Anything).Return()
	pm.Start()
	return pm
}

// VoteAggregator is a mock for testing eventhandler
type VoteAggregator struct {
	// if a blockID exists in qcs field, then a vote can be made into a QC
	qcs map[flow.Identifier]*model.QuorumCertificate
	t   *testing.T
}

func NewVoteAggregator(t *testing.T) *VoteAggregator {
	return &VoteAggregator{
		qcs: make(map[flow.Identifier]*model.QuorumCertificate),
		t:   t,
	}
}

func (v *VoteAggregator) StoreVoteAndBuildQC(vote *model.Vote, block *model.Block) (*model.QuorumCertificate, bool, error) {
	qc, ok := v.qcs[block.BlockID]
	v.t.Logf("voteaggregator.StoreVoteAndBuildQC, qc built: %v, for view: %x, blockID: %v\n", ok, block.View, block.BlockID)

	return qc, ok, nil
}

func (v *VoteAggregator) StorePendingVote(vote *model.Vote) bool {
	return false
}

func (v *VoteAggregator) StoreProposerVote(vote *model.Vote) bool {
	return true
}

func (v *VoteAggregator) BuildQCOnReceivedBlock(block *model.Block) (*model.QuorumCertificate, bool, error) {
	qc, ok := v.qcs[block.BlockID]
	v.t.Logf("voteaggregator.BuildQCOnReceivedBlock, qc built: %v, for view: %x, blockID: %v\n", ok, block.View, block.BlockID)

	return qc, ok, nil
}

func (v *VoteAggregator) PruneByView(view uint64) {
	v.t.Logf("pruned at view:%v\n", view)
}

type Committee struct {
	mocks.Committee
	// to mock I'm the leader of a certain view, add the view into the keys of leaders field
	leaders map[uint64]struct{}
}

func NewCommittee() *Committee {
	return &Committee{
		leaders: make(map[uint64]struct{}),
	}
}

func (c *Committee) LeaderForView(view uint64) (flow.Identifier, error) {
	_, isLeader := c.leaders[view]
	if isLeader {
		return flow.Identifier{0x01}, nil
	}
	return flow.Identifier{0x00}, nil
}

func (c *Committee) Self() flow.Identifier {
	return flow.Identifier{0x01}
}

// The Voter mock will not vote for any block unless the block's ID exists in votable field's key
type Voter struct {
	votable       map[flow.Identifier]struct{}
	lastVotedView uint64
	t             *testing.T
}

func NewVoter(t *testing.T, lastVotedView uint64) *Voter {
	return &Voter{
		votable:       make(map[flow.Identifier]struct{}),
		lastVotedView: lastVotedView,
		t:             t,
	}
}

// voter will not vote for any block, unless the blockID exists in votable map
func (v *Voter) ProduceVoteIfVotable(block *model.Block, curView uint64) (*model.Vote, error) {
	_, ok := v.votable[block.BlockID]
	if !ok {
		return nil, &model.NoVoteError{}
	}
	return createVote(block), nil
}

// Forks mock allows to customize the Add QC and AddBlock function by specifying the addQC and addBlock callbacks
type Forks struct {
	mocks.Forks
	blocks    map[flow.Identifier]*model.Block
	finalized uint64
	t         *testing.T
	qc        *model.QuorumCertificate
	// addQC is to customize the logic to change finalized view
	addQC func(qc *model.QuorumCertificate) error
	// addBlock is to customize the logic to change finalized view
	addBlock func(block *model.Block) error
}

func NewForks(t *testing.T, finalized uint64) *Forks {
	f := &Forks{
		blocks:    make(map[flow.Identifier]*model.Block),
		finalized: finalized,
		t:         t,
	}

	f.addQC = func(qc *model.QuorumCertificate) error {
		if f.qc == nil || qc.View > f.qc.View {
			f.qc = qc
		}
		return nil
	}
	f.addBlock = func(block *model.Block) error {
		f.blocks[block.BlockID] = block
		_ = f.addQC(block.QC)
		return nil
	}

	return f
}

func (f *Forks) AddBlock(block *model.Block) error {
	f.t.Logf("forks.AddBlock received Block for view: %v, qc: %v\n", block.View, block.QC.View)
	return f.addBlock(block)
}

func (f *Forks) AddQC(qc *model.QuorumCertificate) error {
	f.t.Logf("forks.AddQC received QC for view: %v\n", qc.View)
	return f.addQC(qc)
}

func (f *Forks) FinalizedView() uint64 {
	return f.finalized
}

func (f *Forks) GetBlock(blockID flow.Identifier) (*model.Block, bool) {
	b, ok := f.blocks[blockID]
	var view uint64
	if ok {
		view = b.View
	}
	f.t.Logf("forks.GetBlock found: %v, view: %v\n", ok, view)
	return b, ok
}

func (f *Forks) GetBlocksForView(view uint64) []*model.Block {
	blocks := make([]*model.Block, 0)
	for _, b := range f.blocks {
		if b.View == view {
			blocks = append(blocks, b)
		}
	}
	f.t.Logf("forks.GetBlocksForView found %v block(s) for view %v\n", len(blocks), view)
	return blocks
}

func (f *Forks) MakeForkChoice(curView uint64) (*model.QuorumCertificate, *model.Block, error) {
	if f.qc == nil {
		f.t.Fatalf("cannot make fork choice for curview: %v", curView)
	}

	block, ok := f.blocks[f.qc.BlockID]
	if !ok {
		return nil, nil, fmt.Errorf("cannot block %V for fork choice qc", f.qc.BlockID)
	}
	f.t.Logf("forks.MakeForkChoice for view: %v, qc view: %v\n", curView, f.qc.View)
	return f.qc, block, nil
}

// BlockProducer mock will always make a valid block
type BlockProducer struct{}

func (b *BlockProducer) MakeBlockProposal(qc *model.QuorumCertificate, view uint64) (*model.Proposal, error) {
	return createProposal(view, qc.View), nil
}

// BlacklistValidator is Validator mock that consider all proposals are valid unless the proposal's BlockID exists
// in the invalidProposals key or unverifiable key
type BlacklistValidator struct {
	mocks.Validator
	invalidProposals map[flow.Identifier]struct{}
	unverifiable     map[flow.Identifier]struct{}
	t                *testing.T
}

func NewBlacklistValidator(t *testing.T) *BlacklistValidator {
	return &BlacklistValidator{
		invalidProposals: make(map[flow.Identifier]struct{}),
		unverifiable:     make(map[flow.Identifier]struct{}),
		t:                t,
	}
}

func (v *BlacklistValidator) ValidateProposal(proposal *model.Proposal) error {
	// check if is invalid
	_, ok := v.invalidProposals[proposal.Block.BlockID]
	if ok {
		v.t.Logf("invalid proposal: %v\n", proposal.Block.View)
		return &model.ErrorInvalidBlock{
			BlockID: proposal.Block.BlockID,
			View:    proposal.Block.View,
			Err:     fmt.Errorf("some error"),
		}
	}

	// check if is unverifiable
	_, ok = v.unverifiable[proposal.Block.BlockID]
	if ok {
		v.t.Logf("unverifiable proposal: %v\n", proposal.Block.View)
		return model.ErrUnverifiableBlock
	}

	return nil
}

func TestEventHandler(t *testing.T) {
	suite.Run(t, new(EventHandlerSuite))
}

type EventHandlerSuite struct {
	suite.Suite

	eventhandler *eventhandler.EventHandler

	paceMaker      hotstuff.PaceMaker
	forks          *Forks
	persist        *mocks.Persister
	blockProducer  *BlockProducer
	communicator   *mocks.Communicator
	committee      *Committee
	voteAggregator *VoteAggregator
	voter          *Voter
	validator      *BlacklistValidator
	notifier       hotstuff.Consumer

	initView    uint64
	endView     uint64
	vote        *model.Vote
	votingBlock *model.Block
	qc          *model.QuorumCertificate
	newview     *model.NewViewEvent
}

func (es *EventHandlerSuite) SetupTest() {
	finalized, curView := uint64(3), uint64(6)

	es.paceMaker = initPaceMaker(es.T(), curView)
	es.forks = NewForks(es.T(), finalized)
	es.persist = &mocks.Persister{}
	es.persist.On("StartedView", mock.Anything).Return(nil)
	es.persist.On("VotedView", mock.Anything).Return(nil)
	es.blockProducer = &BlockProducer{}
	es.communicator = &mocks.Communicator{}
	es.communicator.On("BroadcastProposal", mock.Anything).Return(nil)
	es.communicator.On("SendVote", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	es.committee = NewCommittee()
	es.voteAggregator = NewVoteAggregator(es.T())
	es.voter = NewVoter(es.T(), finalized)
	es.validator = NewBlacklistValidator(es.T())
	es.notifier = &notifications.NoopConsumer{}

	// initialize no-op metrics mock
	metrics := &module.Metrics{}
	metrics.On("MadeBlockProposal")

	eventhandler, err := eventhandler.New(
		zerolog.New(os.Stderr),
		metrics,
		es.paceMaker,
		es.blockProducer,
		es.forks,
		es.persist,
		es.communicator,
		es.committee,
		es.voteAggregator,
		es.voter,
		es.validator,
		es.notifier)
	require.NoError(es.T(), err)

	es.eventhandler = eventhandler

	es.initView = curView
	es.endView = curView
	// voting block is a block for the current view, which will trigger view change
	es.votingBlock = createBlockWithQC(es.paceMaker.CurView(), es.paceMaker.CurView()-1)
	es.vote = &model.Vote{
		BlockID:  es.votingBlock.BlockID,
		View:     es.votingBlock.View,
		SignerID: flow.ZeroID,
		SigData:  nil,
	}
	es.qc = &model.QuorumCertificate{
		BlockID:   es.votingBlock.BlockID,
		View:      es.votingBlock.View,
		SignerIDs: nil,
		SigData:   nil,
	}
	es.newview = &model.NewViewEvent{
		View: es.votingBlock.View + 1, // the vote for the voting blocks will trigger a view change to the next view
	}
}

func (es *EventHandlerSuite) TestVoteLowerFinalView() {
	es.vote.View = uint64(es.forks.finalized - 1)

	err := es.eventhandler.OnReceiveVote(es.vote)
	require.NoError(es.T(), err, "if a vote's view is lower than the finalized view, "+
		"it should be ignored")
	require.Equal(es.T(), es.endView, es.paceMaker.CurView(), "incorrect view change")
}

func (es *EventHandlerSuite) TestVoteEqualFinalView() {
	es.vote.View = uint64(es.forks.finalized)

	err := es.eventhandler.OnReceiveVote(es.vote)
	require.NoError(es.T(), err, "if a vote's view equals to the finalized view,"+
		"it should be ignored")
	require.Equal(es.T(), es.endView, es.paceMaker.CurView(), "incorrect view change")
}

func (es *EventHandlerSuite) TestMissingVotingBlock() {
	// voting block doesn't exist

	err := es.eventhandler.OnReceiveVote(es.vote)
	require.NoError(es.T(), err, "if voting block is missing, the pending vote will be stored,"+
		"but not processed")
	require.Equal(es.T(), es.endView, es.paceMaker.CurView(), "incorrect view change")
}

func (es *EventHandlerSuite) TestNoQCBuilt() {
	// voting block exists
	es.forks.blocks[es.vote.BlockID] = es.votingBlock

	// no qc is built
	err := es.eventhandler.OnReceiveVote(es.vote)
	require.NoError(es.T(), err, "if voting block exists, the vote will be stored,"+
		"if a QC can not be built, then no QC will be processed")
	require.Equal(es.T(), es.endView, es.paceMaker.CurView(), "incorrect view change")
}

func (es *EventHandlerSuite) TestQCBuiltNoViewChange() {
	// voting block exists
	oldBlock := createBlockWithQC(es.paceMaker.CurView()-1, es.paceMaker.CurView()-2)
	es.forks.blocks[oldBlock.BlockID] = oldBlock

	// create an old vote
	oldVote := createVote(oldBlock)

	// a qc is built
	es.voteAggregator.qcs[oldBlock.BlockID] = createQC(oldBlock)

	// new qc is added to forks

	err := es.eventhandler.OnReceiveVote(oldVote)
	require.NoError(es.T(), err, "if a vote can trigger a QC to be built,"+
		"but the QC didn't trigger view change, then it won't start new view")
	require.Equal(es.T(), es.endView, es.paceMaker.CurView(), "incorrect view change")
}

func (es *EventHandlerSuite) TestQCBuiltViewChanged() {
	// voting block exists
	es.forks.blocks[es.vote.BlockID] = es.votingBlock

	// a qc is built
	es.voteAggregator.qcs[es.vote.BlockID] = createQC(es.votingBlock)

	// new qc is added to forks
	// view changed
	// I'm not the next leader
	// haven't received block for next view
	// goes to the new view
	es.endView++
	// not the leader of the newview
	// don't have block for the newview
	// over

	err := es.eventhandler.OnReceiveVote(es.vote)
	require.NoError(es.T(), err, "if a vote can trigger a QC to be built,"+
		"and the QC triggered a view change, then start new view")
	require.Equal(es.T(), es.endView, es.paceMaker.CurView(), "incorrect view change")
}

// in the newview, I'm not the leader, and I have the cur block,
// and the block is not a safe node, and I'm the next leader, and no qc built for this block.
func (es *EventHandlerSuite) TestInNewView_NotLeader_HasBlock_NoVote_IsNextLeader_NoQC() {
	// voting block exists
	es.forks.blocks[es.vote.BlockID] = es.votingBlock
	// a qc is built
	es.voteAggregator.qcs[es.vote.BlockID] = createQC(es.votingBlock)
	// viewchanged
	es.endView++
	// not leader for newview

	// has block for newview
	newviewblock := createBlockWithQC(es.newview.View, es.newview.View-1)
	es.forks.blocks[newviewblock.BlockID] = newviewblock

	// not to vote for the new view block

	// I'm the next leader
	es.committee.leaders[es.newview.View+1] = struct{}{}

	// no QC for the new view
	err := es.eventhandler.OnReceiveVote(es.vote)
	require.NoError(es.T(), err)
	require.Equal(es.T(), es.endView, es.paceMaker.CurView(), "incorrect view change")
}

// TestInNewView_NotLeader_HasBlock_NoVote_IsNextLeader_QCBuilt_NoViewChange doesn't exist

// in the newview, I'm not the leader, and I have the cur block,
// and the block is not a safe node, and I'm the next leader, and a qc is built for this block,
// and the qc triggered view change.
func (es *EventHandlerSuite) TestInNewView_NotLeader_HasBlock_NoVote_IsNextLeader_QCBuilt_ViewChanged() {
	// voting block exists
	es.forks.blocks[es.vote.BlockID] = es.votingBlock
	// a qc is built
	es.voteAggregator.qcs[es.vote.BlockID] = createQC(es.votingBlock)
	// viewchanged
	es.endView++
	// not leader for newview

	// has block for newview
	newviewblock := createBlockWithQC(es.newview.View, es.newview.View-1)
	es.forks.blocks[newviewblock.BlockID] = newviewblock

	// not to vote for the new view block

	// I'm the next leader
	es.committee.leaders[es.newview.View+1] = struct{}{}

	// qc built for the new view block
	es.voteAggregator.qcs[newviewblock.BlockID] = createQC(newviewblock)
	// view change by this qc
	es.endView++
	// not leader in next view, viewchange
	es.endView++

	err := es.eventhandler.OnReceiveVote(es.vote)
	require.NoError(es.T(), err)
	require.Equal(es.T(), es.endView, es.paceMaker.CurView(), "incorrect view change")
}

// in the newview, I'm not the leader, and I have the cur block,
// and the block is a safe node to vote, and I'm the next leader, and no qc is built for this block.
func (es *EventHandlerSuite) TestInNewView_NotLeader_HasBlock_NotSafeNode_IsNextLeader_Voted_NoQC() {
	// voting block exists
	es.forks.blocks[es.vote.BlockID] = es.votingBlock
	// a qc is built
	es.voteAggregator.qcs[es.vote.BlockID] = createQC(es.votingBlock)
	// viewchanged by new qc
	es.endView++
	// not leader for newview

	// has block for newview
	newviewblock := createBlockWithQC(es.newview.View, es.newview.View-1)
	es.forks.blocks[newviewblock.BlockID] = newviewblock

	// not to vote for the new view block

	// I'm the next leader
	es.committee.leaders[es.newview.View+1] = struct{}{}

	// no qc for the newview block

	// should not trigger view change
	err := es.eventhandler.OnReceiveVote(es.vote)
	require.NoError(es.T(), err)
	require.Equal(es.T(), es.endView, es.paceMaker.CurView(), "incorrect view change")
}

// in the newview, I'm not the leader, and I have the cur block,
// and the block is not a safe node to vote, and I'm not the next leader
func (es *EventHandlerSuite) TestInNewView_NotLeader_HasBlock_NotSafeNode_NotNextLeader() {
	// voting block exists
	es.forks.blocks[es.vote.BlockID] = es.votingBlock
	// a qc is built
	es.voteAggregator.qcs[es.vote.BlockID] = createQC(es.votingBlock)
	// viewchanged by new qc
	es.endView++

	// view changed to newview
	// I'm not the leader for newview

	// have received block for cur view
	newviewblock := createBlockWithQC(es.newview.View, es.newview.View-1)
	es.forks.blocks[newviewblock.BlockID] = newviewblock

	// I'm not the next leader
	// no vote for this block
	// goes to the next view
	es.endView++
	// not leader for next view

	err := es.eventhandler.OnReceiveVote(es.vote)
	require.NoError(es.T(), err, "if a vote can trigger a QC to be built,"+
		"and the QC triggered a view change, then start new view")
	require.Equal(es.T(), es.endView, es.paceMaker.CurView(), "incorrect view change")
}

// receiving an invalid proposal should not trigger view change
func (es *EventHandlerSuite) TestOnReceiveProposal_InvalidProposal_NoViewChange() {
	proposal := createProposal(es.initView, es.initView-1)
	// invalid proposal
	es.validator.invalidProposals[proposal.Block.BlockID] = struct{}{}

	err := es.eventhandler.OnReceiveProposal(proposal)
	require.NoError(es.T(), err)
	require.Equal(es.T(), es.endView, es.paceMaker.CurView(), "incorrect view change")
}

// received a valid proposal that has older view, and cannot build qc from votes for this block,
// the proposal's QC didn't trigger view change
func (es *EventHandlerSuite) TestOnReceiveProposal_OlderThanCurView_CannotBuildQCFromVotes_NoViewChange() {
	proposal := createProposal(es.initView-1, es.initView-2)

	// can not build qc from votes for block
	// should not trigger view change
	err := es.eventhandler.OnReceiveProposal(proposal)
	require.NoError(es.T(), err)
	require.Equal(es.T(), es.endView, es.paceMaker.CurView(), "incorrect view change")
}

// received a valid proposal that has older view, and can built a qc from votes for this block,
// the proposal's QC didn't trigger view change
func (es *EventHandlerSuite) TestOnReceiveProposal_OlderThanCurView_CanBuildQCFromVotes_NoViewChange() {
	proposal := createProposal(es.initView-1, es.initView-2)

	// a qc is built
	es.voteAggregator.qcs[proposal.Block.BlockID] = createQC(proposal.Block)
	// should not trigger view change
	err := es.eventhandler.OnReceiveProposal(proposal)
	require.NoError(es.T(), err)
	require.Equal(es.T(), es.endView, es.paceMaker.CurView(), "incorrect view change")
}

// received a valid proposal that has newer view, and cannot build qc from votes for this block,
// the proposal's QC triggered view change
func (es *EventHandlerSuite) TestOnReceiveProposal_NewerThanCurView_CannotBuildQCFromVotes_ViewChange() {
	proposal := createProposal(es.initView+1, es.initView)

	// can not build qc from votes for block
	// block 7 triggered view change
	es.endView++

	// not leader of view 7, go to view 8
	es.endView++
	err := es.eventhandler.OnReceiveProposal(proposal)
	require.NoError(es.T(), err)
	require.Equal(es.T(), es.endView, es.paceMaker.CurView(), "incorrect view change")
}

// received a valid proposal that has newer view, and can build qc from votes for this block,
// the proposal's QC triggered view change
func (es *EventHandlerSuite) TestOnReceiveProposal_NewerThanCurView_CanBuildQCFromVotes_ViewChange() {
	proposal := createProposal(es.initView+1, es.initView)

	es.forks.blocks[proposal.Block.BlockID] = proposal.Block
	// a qc is built
	es.voteAggregator.qcs[proposal.Block.BlockID] = createQC(proposal.Block)
	// trigged view change
	es.endView++
	// the proposal is for next view, has block for next view, no vote, trigger view change
	es.endView++

	err := es.eventhandler.OnReceiveProposal(proposal)
	require.NoError(es.T(), err)
	require.Equal(es.T(), es.endView, es.paceMaker.CurView(), "incorrect view change")
}

// received a valid proposal whose QC that has newer view, and cannot build qc from votes for this block,
// the proposal's QC triggered view change
func (es *EventHandlerSuite) TestOnReceiveProposal_QCNewerThanCurView_CannotBuildQCFromVotes_ViewChanged() {
	proposal := createProposal(es.initView+2, es.initView+1)

	// can not build qc from votes for block
	// block 8 triggered view change
	es.endView = es.endView + 2

	// not leader of view 8, go to view 9
	es.endView++
	err := es.eventhandler.OnReceiveProposal(proposal)
	require.NoError(es.T(), err)
	require.Equal(es.T(), es.endView, es.paceMaker.CurView(), "incorrect view change")
	require.Contains(es.T(), es.forks.blocks, proposal.Block.BlockID, "proposal block should be stored")
}

// received a valid proposal for cur view, but not a safe node to vote, and I'm the next leader,
// no qc for the block
func (es *EventHandlerSuite) TestOnReceiveProposal_ForCurView_NoVote_IsNextLeader_NoQC() {
	proposal := createProposal(es.initView, es.initView-1)
	// I'm the next leader
	es.committee.leaders[es.initView+1] = struct{}{}
	// no qc can be built for this block
	err := es.eventhandler.OnReceiveProposal(proposal)
	require.NoError(es.T(), err)
	require.Equal(es.T(), es.endView, es.paceMaker.CurView(), "incorrect view change")
}

// received a valid proposal for cur view, but not a safe node to vote, and I'm the next leader,
// a qc can be built for the block, trigged view change
func (es *EventHandlerSuite) TestOnReceiveProposal_ForCurView_NoVote_IsNextLeader_QCBuilt_ViewChange() {
	proposal := createProposal(es.initView, es.initView-1)
	// I'm the next leader
	es.committee.leaders[es.initView+1] = struct{}{}
	// a qc can be built for this block
	es.voteAggregator.qcs[proposal.Block.BlockID] = createQC(proposal.Block)
	// qc triggered view change
	es.endView++
	// I'm the leader of cur view (7)
	// I'm not the leader of next view (8), trigger view change
	es.endView++
	// I'm not the leader of cur view(8)
	// no block for curView, over

	err := es.eventhandler.OnReceiveProposal(proposal)
	require.NoError(es.T(), err)
	require.Equal(es.T(), es.endView, es.paceMaker.CurView(), "incorrect view change")
}

// received a unverifiable proposal for future view, no view change
func (es *EventHandlerSuite) TestOnReceiveProposal_Unverifiable() {
	// qc.View is below the finalized view
	proposal := createProposal(es.forks.finalized+2, es.forks.finalized-1)

	// proposal is unverifiable
	es.validator.unverifiable[proposal.Block.BlockID] = struct{}{}

	err := es.eventhandler.OnReceiveProposal(proposal)
	require.NoError(es.T(), err)
	require.Equal(es.T(), es.endView, es.paceMaker.CurView(), "incorrect view change")
}

func (es *EventHandlerSuite) TestOnTimeout() {
	err := es.eventhandler.OnLocalTimeout()
	// timeout will trigger viewchange
	es.endView++
	require.NoError(es.T(), err)
	require.Equal(es.T(), es.endView, es.paceMaker.CurView(), "incorrect view change")
}

func (es *EventHandlerSuite) Test100Timeout() {
	for i := 0; i < 100; i++ {
		err := es.eventhandler.OnLocalTimeout()
		es.endView++
		require.NoError(es.T(), err)
	}
	require.Equal(es.T(), es.endView, es.paceMaker.CurView(), "incorrect view change")
}

// a leader builds 100 blocks one after another
func (es *EventHandlerSuite) TestLeaderBuild100Blocks() {
	proposal := createProposal(es.initView, es.initView-1)

	for i := 0; i < 100; i++ {
		// I'm the leader for 100 views
		es.committee.leaders[es.initView+uint64(i)] = struct{}{}
		// I can build qc for all 100 views
		proposal := createProposal(es.initView+uint64(i), es.initView+uint64(i)-1)
		es.voteAggregator.qcs[proposal.Block.BlockID] = createQC(proposal.Block)
		es.voter.votable[proposal.Block.BlockID] = struct{}{}
		// should trigger 100 view change
		es.endView++
	}

	err := es.eventhandler.OnReceiveProposal(proposal)
	require.NoError(es.T(), err)
	require.Equal(es.T(), es.endView, es.paceMaker.CurView(), "incorrect view change")
	require.Equal(es.T(), 100, len(es.forks.blocks))
}

// a follower receives 100 blocks
func (es *EventHandlerSuite) TestFollowerFollows100Blocks() {
	for i := 0; i < 100; i++ {
		proposal := createProposal(es.initView+uint64(i), es.initView+uint64(i)-1)
		err := es.eventhandler.OnReceiveProposal(proposal)
		require.NoError(es.T(), err)
		es.endView++
	}
	require.Equal(es.T(), es.endView, es.paceMaker.CurView(), "incorrect view change")
	require.Equal(es.T(), 100, len(es.forks.blocks))
}

// a follower receives 100 forks built on top of the same block
func (es *EventHandlerSuite) TestFollowerReceives100Forks() {
	for i := 0; i < 100; i++ {
		proposal := createProposal(es.initView+uint64(i)+1, es.initView-1)
		err := es.eventhandler.OnReceiveProposal(proposal)
		require.NoError(es.T(), err)
	}
	require.Equal(es.T(), es.endView, es.paceMaker.CurView(), "incorrect view change")
	require.Equal(es.T(), 100, len(es.forks.blocks))
}

func createBlock(view uint64) *model.Block {
	blockID := flow.MakeID(struct {
		BlockID uint64
	}{
		BlockID: view,
	})
	return &model.Block{
		BlockID: blockID,
		View:    uint64(view),
	}
}

func createBlockWithQC(view uint64, qcview uint64) *model.Block {
	block := createBlock(view)
	parent := createBlock(qcview)
	block.QC = createQC(parent)
	return block
}

func createQC(parent *model.Block) *model.QuorumCertificate {
	qc := &model.QuorumCertificate{
		BlockID:   parent.BlockID,
		View:      parent.View,
		SignerIDs: nil,
		SigData:   nil,
	}
	return qc
}

func createVote(block *model.Block) *model.Vote {
	return &model.Vote{
		View:     block.View,
		BlockID:  block.BlockID,
		SignerID: flow.ZeroID,
		SigData:  nil,
	}
}

func createProposal(view uint64, qcview uint64) *model.Proposal {
	block := createBlockWithQC(view, qcview)
	return &model.Proposal{
		Block:   block,
		SigData: nil,
	}
}
