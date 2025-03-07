// Code generated by mockery v2.46.3. DO NOT EDIT.

// Copyright (c) Abstract Machines

package mocks

import (
	messaging "github.com/absmach/supermq/pkg/messaging"
	mock "github.com/stretchr/testify/mock"
)

// Notifier is an autogenerated mock type for the Notifier type
type Notifier struct {
	mock.Mock
}

// Notify provides a mock function with given fields: from, to, msg
func (_m *Notifier) Notify(from string, to []string, msg *messaging.Message) error {
	ret := _m.Called(from, to, msg)

	if len(ret) == 0 {
		panic("no return value specified for Notify")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string, *messaging.Message) error); ok {
		r0 = rf(from, to, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewNotifier creates a new instance of Notifier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNotifier(t interface {
	mock.TestingT
	Cleanup(func())
}) *Notifier {
	mock := &Notifier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
