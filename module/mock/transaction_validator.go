// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"
)

// TransactionValidator is an autogenerated mock type for the TransactionValidator type
type TransactionValidator struct {
	mock.Mock
}

// ValidateTransaction provides a mock function with given fields: tx
func (_m *TransactionValidator) ValidateTransaction(tx *flow.TransactionBody) error {
	ret := _m.Called(tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(*flow.TransactionBody) error); ok {
		r0 = rf(tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}