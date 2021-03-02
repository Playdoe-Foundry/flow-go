package trace

// Span names
const (

	// Common
	//

	// State
	// mutator.Extend - full payload check
	ProtoStateMutatorExtend                SpanName = "common.state.proto.mutator.extend"
	ProtoStateMutatorExtendCheckHeader     SpanName = "common.state.proto.mutator.extend.checkHeader"
	ProtoStateMutatorExtendCheckGuarantees SpanName = "common.state.proto.mutator.extend.checkGuarantees"
	ProtoStateMutatorExtendCheckSeals      SpanName = "common.state.proto.mutator.extend.checkSeals"
	ProtoStateMutatorExtendCheckReceipts   SpanName = "common.state.proto.mutator.extend.checkReceipts"
	ProtoStateMutatorExtendDBInsert        SpanName = "common.state.proto.mutator.extend.dbInsert"

	// mutator.HeaderExtend - header-only check
	ProtoStateMutatorHeaderExtend              SpanName = "common.state.proto.mutator.headerExtend"
	ProtoStateMutatorHeaderExtendGetLastSealed SpanName = "common.state.proto.mutator.headerExtend.lastSealed"

	// mutator.Finalize
	ProtoStateMutatorFinalize SpanName = "common.state.proto.mutator.finalize"

	// Consensus Node
	//

	CONProcessCollection SpanName = "con.processCollection"
	CONProcessBlock      SpanName = "con.processBlock"

	// Hotstuff
	CONHotFinalizeCollection SpanName = "con.hotstuff.finalizeCollection"
	CONHotFinalizeBlock      SpanName = "con.hotstuff.finalizeBlock"

	// Ingestion
	CONIngOnCollectionGuarantee SpanName = "con.ingestion.onCollectionGuarantee"

	// Propagation
	CONPropOnGuarantee SpanName = "con.propagation.onGuarantee"

	// Provider
	CONProvOnBlockProposal SpanName = "con.provider.onBlockProposal"

	// Compliance
	CONCompBroadcastProposalWithDelay      SpanName = "con.compliance.BroadcastProposalWithDelay"
	CONCompOnBlockProposal                 SpanName = "con.compliance.onBlockProposal"
	CONCompOnBlockProposalProcessRecursive SpanName = "con.compliance.onBlockProposal.processBlockProposal.recursive"
	CONCompOnBlockProposalProcessSingle    SpanName = "con.compliance.onBlockProposal.processBlockProposal.single"

	// Matching
	CONMatchCheckSealing                        SpanName = "con.matching.checkSealing"
	CONMatchCheckSealingSealableResults         SpanName = "con.matching.checkSealing.sealableResults"
	CONMatchCheckSealingClearPools              SpanName = "con.matching.checkSealing.clearPools"
	CONMatchCheckSealingRequestPendingReceipts  SpanName = "con.matching.checkSealing.requestPendingReceipts"
	CONMatchCheckSealingRequestPendingApprovals SpanName = "con.matching.checkSealing.requestPendingApprovals"
	CONMatchOnReceipt                           SpanName = "con.matching.onReceipt"
	CONMatchOnReceiptVal                        SpanName = "con.matching.onReceipt.validation"
	CONMatchOnApproval                          SpanName = "con.matching.onApproval"

	// Builder
	CONBuildOn                        SpanName = "con.builder"
	CONBuildOnCreatePayloadGuarantees SpanName = "con.builder.createPayload.guarantees"
	CONBuildOnCreatePayloadSeals      SpanName = "con.builder.createPayload.seals"
	CONBuildOnCreatePayloadReceipts   SpanName = "con.builder.createPayload.receipts"
	CONBuildOnCreateHeader            SpanName = "con.builder.createHeader"
	CONBuildOnDBInsert                SpanName = "con.builder.dbInsert"

	// Collection Node
	//

	// Builder
	COLBuildOn                  SpanName = "col.builder"
	COLBuildOnSetup             SpanName = "col.builder.setup"
	COLBuildOnUnfinalizedLookup SpanName = "col.builder.unfinalizedLookup"
	COLBuildOnFinalizedLookup   SpanName = "col.builder.finalizedLookup"
	COLBuildOnCreatePayload     SpanName = "col.builder.createPayload"
	COLBuildOnCreateHeader      SpanName = "col.builder.createHeader"
	COLBuildOnDBInsert          SpanName = "col.builder.dbInsert"

	// Cluster State
	COLClusterStateMutatorExtend                       SpanName = "col.state.mutator.extend"
	COLClusterStateMutatorExtendSetup                  SpanName = "col.state.mutator.extend.setup"
	COLClusterStateMutatorExtendCheckAncestry          SpanName = "col.state.mutator.extend.ancestry"
	COLClusterStateMutatorExtendCheckTransactionsValid SpanName = "col.state.mutator.extend.transactions.validity"
	COLClusterStateMutatorExtendCheckTransactionsDupes SpanName = "col.state.mutator.extend.transactions.dupes"
	COLClusterStateMutatorExtendDBInsert               SpanName = "col.state.mutator.extend.dbInsert"

	// Execution Node
	//

	EXEExecuteBlock           SpanName = "exe.ingestion.executeBlock"
	EXESaveExecutionResults   SpanName = "exe.ingestion.saveExecutionResults"
	EXESaveExecutionReceipt   SpanName = "exe.ingestion.saveExecutionReceipt"
	EXESaveTransactionResults SpanName = "exe.ingestion.saveTransactionResults"
	EXESaveTransactionEvents  SpanName = "exe.ingestion.saveTransactionEvents"

	EXEBroadcastExecutionReceipt SpanName = "exe.provider.broadcastExecutionReceipt"

	EXEComputeBlock            SpanName = "exe.computer.computeBlock"
	EXEComputeCollection       SpanName = "exe.computer.computeCollection"
	EXEComputeSystemCollection SpanName = "exe.computer.computeSystemCollection"
	EXEComputeTransaction      SpanName = "exe.computer.computeTransaction"

	EXECommitDelta                        SpanName = "exe.state.commitDelta"
	EXEGetRegisters                       SpanName = "exe.state.getRegisters"
	EXEGetRegistersWithProofs             SpanName = "exe.state.getRegistersWithProofs"
	EXEPersistStateCommitment             SpanName = "exe.state.persistStateCommitment"
	EXEPersistChunkDataPack               SpanName = "exe.state.persistChunkDataPack"
	EXEGetExecutionResultID               SpanName = "exe.state.getExecutionResultID"
	EXEPersistExecutionResult             SpanName = "exe.state.persistExecutionResult"
	EXEPersistStateInteractions           SpanName = "exe.state.persistStateInteractions"
	EXERetrieveStateDelta                 SpanName = "exe.state.retrieveStateDelta"
	EXEUpdateHighestExecutedBlockIfHigher SpanName = "exe.state.updateHighestExecutedBlockIfHigher"
	EXEGetHighestExecutedBlockID          SpanName = "exe.state.getHighestExecutedBlockID"

	// Verification node
	//
	// assigner engine
	VERProcessFinalizedBlock SpanName = "ver.processFinalizedBlock"
	// children of VERProcessFinalizedBlock
	VERAssignerHandleFinalizedBlock   SpanName = "ver.assigner.handleFinalizedBlock"
	VERAssignerHandleExecutionReceipt SpanName = "ver.assigner.handleExecutionReceipt"
	VERAssignerChunkAssignment        SpanName = "ver.assigner.chunkAssignment"
	VERAssignerProcessChunk           SpanName = "ver.assigner.processChunk"

	VERProcessExecutionReceipt SpanName = "ver.processExecutionReceipt"
	// children of VERProcessExecutionReceipt
	VERFindHandleExecutionReceipt SpanName = "ver.find.handleExecutionReceipt"
	VERFindOnFinalizedBlock       SpanName = "ver.finder.OnFinalizedBlock"
	VERFindCheckCachedReceipts    SpanName = "ver.finder.checkCachedReceipts"
	VERFindCheckPendingReceipts   SpanName = "ver.finder.checkPendingReceipts"
	VERFindCheckReadyReceipts     SpanName = "ver.finder.checkReadyReceipts"
	VERFindProcessResult          SpanName = "ver.finder.processResult"
	VERFindOnResultProcessed      SpanName = "ver.finder.onResultProcessed"

	VERProcessExecutionResult SpanName = "ver.processExecutionResult"
	// children of VERProcessExecutionResult
	VERMatchHandleExecutionResult SpanName = "ver.match.handleExecutionResult"
	VERMatchHandleChunkDataPack   SpanName = "ver.match.handleChunkDataPack"
	VERMatchMyChunkAssignments    SpanName = "ver.match.myChunkAssignments"
	VERVerVerifyWithMetrics       SpanName = "ver.verify.verifyWithMetrics"
	VERVerChunkVerify             SpanName = "ver.verify.ChunkVerifier.Verify"
	VERVerGenerateResultApproval  SpanName = "ver.verify.GenerateResultApproval"
)

// Tag names
const (
	EXEParseDurationTag         = "runtime.parseTransactionDuration"
	EXECheckDurationTag         = "runtime.checkTransactionDuration"
	EXEInterpretDurationTag     = "runtime.interpretTransactionDuration"
	EXEValueEncodingDurationTag = "runtime.encodingValueDuration"
	EXEValueDecodingDurationTag = "runtime.decodingValueDuration"
)
