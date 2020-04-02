// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import flow "github.com/dapperlabs/flow-go/model/flow"
import mock "github.com/stretchr/testify/mock"
import state "github.com/dapperlabs/flow-go/engine/execution/state"

// ExecutionState is an autogenerated mock type for the ExecutionState type
type ExecutionState struct {
	mock.Mock
}

// ChunkDataPackByChunkID provides a mock function with given fields: _a0
func (_m *ExecutionState) ChunkDataPackByChunkID(_a0 flow.Identifier) (*flow.ChunkDataPack, error) {
	ret := _m.Called(_a0)

	var r0 *flow.ChunkDataPack
	if rf, ok := ret.Get(0).(func(flow.Identifier) *flow.ChunkDataPack); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.ChunkDataPack)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(flow.Identifier) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChunkHeaderByChunkID provides a mock function with given fields: _a0
func (_m *ExecutionState) ChunkHeaderByChunkID(_a0 flow.Identifier) (*flow.ChunkHeader, error) {
	ret := _m.Called(_a0)

	var r0 *flow.ChunkHeader
	if rf, ok := ret.Get(0).(func(flow.Identifier) *flow.ChunkHeader); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.ChunkHeader)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(flow.Identifier) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CommitDelta provides a mock function with given fields: _a0, _a1
func (_m *ExecutionState) CommitDelta(_a0 state.Delta, _a1 []byte) ([]byte, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(state.Delta, []byte) []byte); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(state.Delta, []byte) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChunkRegisters provides a mock function with given fields: _a0
func (_m *ExecutionState) GetChunkRegisters(_a0 flow.Identifier) (flow.Ledger, error) {
	ret := _m.Called(_a0)

	var r0 flow.Ledger
	if rf, ok := ret.Get(0).(func(flow.Identifier) flow.Ledger); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.Ledger)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(flow.Identifier) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExecutionResultID provides a mock function with given fields: blockID
func (_m *ExecutionState) GetExecutionResultID(blockID flow.Identifier) (flow.Identifier, error) {
	ret := _m.Called(blockID)

	var r0 flow.Identifier
	if rf, ok := ret.Get(0).(func(flow.Identifier) flow.Identifier); ok {
		r0 = rf(blockID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.Identifier)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(flow.Identifier) error); ok {
		r1 = rf(blockID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRegisters provides a mock function with given fields: _a0, _a1
func (_m *ExecutionState) GetRegisters(_a0 []byte, _a1 [][]byte) ([][]byte, error) {
	ret := _m.Called(_a0, _a1)

	var r0 [][]byte
	if rf, ok := ret.Get(0).(func([]byte, [][]byte) [][]byte); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte, [][]byte) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRegistersWithProofs provides a mock function with given fields: _a0, _a1
func (_m *ExecutionState) GetRegistersWithProofs(_a0 []byte, _a1 [][]byte) ([][]byte, [][]byte, error) {
	ret := _m.Called(_a0, _a1)

	var r0 [][]byte
	if rf, ok := ret.Get(0).(func([]byte, [][]byte) [][]byte); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	var r1 [][]byte
	if rf, ok := ret.Get(1).(func([]byte, [][]byte) [][]byte); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([][]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func([]byte, [][]byte) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewView provides a mock function with given fields: _a0
func (_m *ExecutionState) NewView(_a0 []byte) *state.View {
	ret := _m.Called(_a0)

	var r0 *state.View
	if rf, ok := ret.Get(0).(func([]byte) *state.View); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.View)
		}
	}

	return r0
}

// PersistChunkDataPack provides a mock function with given fields: _a0
func (_m *ExecutionState) PersistChunkDataPack(_a0 *flow.ChunkDataPack) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*flow.ChunkDataPack) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PersistChunkHeader provides a mock function with given fields: _a0
func (_m *ExecutionState) PersistChunkHeader(_a0 *flow.ChunkHeader) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*flow.ChunkHeader) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PersistExecutionResult provides a mock function with given fields: blockID, result
func (_m *ExecutionState) PersistExecutionResult(blockID flow.Identifier, result flow.ExecutionResult) error {
	ret := _m.Called(blockID, result)

	var r0 error
	if rf, ok := ret.Get(0).(func(flow.Identifier, flow.ExecutionResult) error); ok {
		r0 = rf(blockID, result)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PersistStateCommitment provides a mock function with given fields: _a0, _a1
func (_m *ExecutionState) PersistStateCommitment(_a0 flow.Identifier, _a1 []byte) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(flow.Identifier, []byte) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StateCommitmentByBlockID provides a mock function with given fields: _a0
func (_m *ExecutionState) StateCommitmentByBlockID(_a0 flow.Identifier) ([]byte, error) {
	ret := _m.Called(_a0)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(flow.Identifier) []byte); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(flow.Identifier) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
