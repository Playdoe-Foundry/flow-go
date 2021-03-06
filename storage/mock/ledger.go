// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"
)

// Ledger is an autogenerated mock type for the Ledger type
type Ledger struct {
	mock.Mock
}

// DiskSize provides a mock function with given fields:
func (_m *Ledger) DiskSize() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Done provides a mock function with given fields:
func (_m *Ledger) Done() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// EmptyStateCommitment provides a mock function with given fields:
func (_m *Ledger) EmptyStateCommitment() []byte {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// GetRegisters provides a mock function with given fields: registerIDs, stateCommitment
func (_m *Ledger) GetRegisters(registerIDs []flow.RegisterID, stateCommitment []byte) ([][]byte, error) {
	ret := _m.Called(registerIDs, stateCommitment)

	var r0 [][]byte
	if rf, ok := ret.Get(0).(func([]flow.RegisterID, []byte) [][]byte); ok {
		r0 = rf(registerIDs, stateCommitment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]flow.RegisterID, []byte) error); ok {
		r1 = rf(registerIDs, stateCommitment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRegistersWithProof provides a mock function with given fields: registerIDs, stateCommitment
func (_m *Ledger) GetRegistersWithProof(registerIDs []flow.RegisterID, stateCommitment []byte) ([][]byte, [][]byte, error) {
	ret := _m.Called(registerIDs, stateCommitment)

	var r0 [][]byte
	if rf, ok := ret.Get(0).(func([]flow.RegisterID, []byte) [][]byte); ok {
		r0 = rf(registerIDs, stateCommitment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	var r1 [][]byte
	if rf, ok := ret.Get(1).(func([]flow.RegisterID, []byte) [][]byte); ok {
		r1 = rf(registerIDs, stateCommitment)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([][]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func([]flow.RegisterID, []byte) error); ok {
		r2 = rf(registerIDs, stateCommitment)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Ready provides a mock function with given fields:
func (_m *Ledger) Ready() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// UpdateRegisters provides a mock function with given fields: registerIDs, values, stateCommitment
func (_m *Ledger) UpdateRegisters(registerIDs []flow.RegisterID, values [][]byte, stateCommitment []byte) ([]byte, error) {
	ret := _m.Called(registerIDs, values, stateCommitment)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]flow.RegisterID, [][]byte, []byte) []byte); ok {
		r0 = rf(registerIDs, values, stateCommitment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]flow.RegisterID, [][]byte, []byte) error); ok {
		r1 = rf(registerIDs, values, stateCommitment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRegistersWithProof provides a mock function with given fields: registerIDs, values, stateCommitment
func (_m *Ledger) UpdateRegistersWithProof(registerIDs []flow.RegisterID, values [][]byte, stateCommitment []byte) ([]byte, [][]byte, error) {
	ret := _m.Called(registerIDs, values, stateCommitment)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]flow.RegisterID, [][]byte, []byte) []byte); ok {
		r0 = rf(registerIDs, values, stateCommitment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 [][]byte
	if rf, ok := ret.Get(1).(func([]flow.RegisterID, [][]byte, []byte) [][]byte); ok {
		r1 = rf(registerIDs, values, stateCommitment)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([][]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func([]flow.RegisterID, [][]byte, []byte) error); ok {
		r2 = rf(registerIDs, values, stateCommitment)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
