// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	flow "github.com/onflow/flow-go/model/flow"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Communicator is an autogenerated mock type for the Communicator type
type Communicator struct {
	mock.Mock
}

// BroadcastProposal provides a mock function with given fields: proposal
func (_m *Communicator) BroadcastProposal(proposal *flow.Header) error {
	ret := _m.Called(proposal)

	var r0 error
	if rf, ok := ret.Get(0).(func(*flow.Header) error); ok {
		r0 = rf(proposal)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BroadcastProposalWithDelay provides a mock function with given fields: proposal, delay
func (_m *Communicator) BroadcastProposalWithDelay(proposal *flow.Header, delay time.Duration) error {
	ret := _m.Called(proposal, delay)

	var r0 error
	if rf, ok := ret.Get(0).(func(*flow.Header, time.Duration) error); ok {
		r0 = rf(proposal, delay)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendVote provides a mock function with given fields: blockID, view, sigData, recipientID
func (_m *Communicator) SendVote(blockID flow.Identifier, view uint64, sigData []byte, recipientID flow.Identifier) error {
	ret := _m.Called(blockID, view, sigData, recipientID)

	var r0 error
	if rf, ok := ret.Get(0).(func(flow.Identifier, uint64, []byte, flow.Identifier) error); ok {
		r0 = rf(blockID, view, sigData, recipientID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
