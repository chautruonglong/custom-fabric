// Code generated by mockery v2.7.4. DO NOT EDIT.

package mocks

import (
	math "github.com/IBM/mathlib"
	mock "github.com/stretchr/testify/mock"
)

// NonceManager is an autogenerated mock type for the NonceManager type
type NonceManager struct {
	mock.Mock
}

// CheckNonce provides a mock function with given fields: nonce
func (_m *NonceManager) CheckNonce(nonce *math.Zr) error {
	ret := _m.Called(nonce)

	var r0 error
	if rf, ok := ret.Get(0).(func(*math.Zr) error); ok {
		r0 = rf(nonce)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetNonce provides a mock function with given fields:
func (_m *NonceManager) GetNonce() (*math.Zr, error) {
	ret := _m.Called()

	var r0 *math.Zr
	if rf, ok := ret.Get(0).(func() *math.Zr); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*math.Zr)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SweepExpiredNonces provides a mock function with given fields:
func (_m *NonceManager) SweepExpiredNonces() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
