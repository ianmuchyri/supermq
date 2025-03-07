// Code generated by mockery v2.46.3. DO NOT EDIT.

// Copyright (c) Abstract Machines

package mocks

import (
	context "context"

	messaging "github.com/absmach/supermq/pkg/messaging"
	mock "github.com/stretchr/testify/mock"
)

// PubSub is an autogenerated mock type for the PubSub type
type PubSub struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *PubSub) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Publish provides a mock function with given fields: ctx, topic, msg
func (_m *PubSub) Publish(ctx context.Context, topic string, msg *messaging.Message) error {
	ret := _m.Called(ctx, topic, msg)

	if len(ret) == 0 {
		panic("no return value specified for Publish")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *messaging.Message) error); ok {
		r0 = rf(ctx, topic, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Subscribe provides a mock function with given fields: ctx, cfg
func (_m *PubSub) Subscribe(ctx context.Context, cfg messaging.SubscriberConfig) error {
	ret := _m.Called(ctx, cfg)

	if len(ret) == 0 {
		panic("no return value specified for Subscribe")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, messaging.SubscriberConfig) error); ok {
		r0 = rf(ctx, cfg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unsubscribe provides a mock function with given fields: ctx, id, topic
func (_m *PubSub) Unsubscribe(ctx context.Context, id string, topic string) error {
	ret := _m.Called(ctx, id, topic)

	if len(ret) == 0 {
		panic("no return value specified for Unsubscribe")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, id, topic)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPubSub creates a new instance of PubSub. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPubSub(t interface {
	mock.TestingT
	Cleanup(func())
}) *PubSub {
	mock := &PubSub{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
