// Code generated by mockery v1.0.0. DO NOT EDIT.

package mempool

import flow "github.com/dapperlabs/flow-go/model/flow"

import mock "github.com/stretchr/testify/mock"

// IdentifierMap is an autogenerated mock type for the IdentifierMap type
type IdentifierMap struct {
	mock.Mock
}

// Append provides a mock function with given fields: key, id
func (_m *IdentifierMap) Append(key flow.Identifier, id flow.Identifier) error {
	ret := _m.Called(key, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(flow.Identifier, flow.Identifier) error); ok {
		r0 = rf(key, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: key
func (_m *IdentifierMap) Get(key flow.Identifier) ([]flow.Identifier, bool) {
	ret := _m.Called(key)

	var r0 []flow.Identifier
	if rf, ok := ret.Get(0).(func(flow.Identifier) []flow.Identifier); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]flow.Identifier)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(flow.Identifier) bool); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Rem provides a mock function with given fields: key
func (_m *IdentifierMap) Rem(key flow.Identifier) bool {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(flow.Identifier) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
