// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Home is an autogenerated mock type for the Home type
type Home struct {
	mock.Mock
}

// GetMessage provides a mock function with given fields:
func (_m *Home) GetMessage() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
