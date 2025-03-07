// Code generated by mockery v2.46.3. DO NOT EDIT.

// Copyright (c) Abstract Machines

package mocks

import (
	context "context"

	policies "github.com/absmach/supermq/pkg/policies"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// AddPolicies provides a mock function with given fields: ctx, prs
func (_m *Service) AddPolicies(ctx context.Context, prs []policies.Policy) error {
	ret := _m.Called(ctx, prs)

	if len(ret) == 0 {
		panic("no return value specified for AddPolicies")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []policies.Policy) error); ok {
		r0 = rf(ctx, prs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddPolicy provides a mock function with given fields: ctx, pr
func (_m *Service) AddPolicy(ctx context.Context, pr policies.Policy) error {
	ret := _m.Called(ctx, pr)

	if len(ret) == 0 {
		panic("no return value specified for AddPolicy")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, policies.Policy) error); ok {
		r0 = rf(ctx, pr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CountObjects provides a mock function with given fields: ctx, pr
func (_m *Service) CountObjects(ctx context.Context, pr policies.Policy) (uint64, error) {
	ret := _m.Called(ctx, pr)

	if len(ret) == 0 {
		panic("no return value specified for CountObjects")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, policies.Policy) (uint64, error)); ok {
		return rf(ctx, pr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, policies.Policy) uint64); ok {
		r0 = rf(ctx, pr)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, policies.Policy) error); ok {
		r1 = rf(ctx, pr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountSubjects provides a mock function with given fields: ctx, pr
func (_m *Service) CountSubjects(ctx context.Context, pr policies.Policy) (uint64, error) {
	ret := _m.Called(ctx, pr)

	if len(ret) == 0 {
		panic("no return value specified for CountSubjects")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, policies.Policy) (uint64, error)); ok {
		return rf(ctx, pr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, policies.Policy) uint64); ok {
		r0 = rf(ctx, pr)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, policies.Policy) error); ok {
		r1 = rf(ctx, pr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePolicies provides a mock function with given fields: ctx, prs
func (_m *Service) DeletePolicies(ctx context.Context, prs []policies.Policy) error {
	ret := _m.Called(ctx, prs)

	if len(ret) == 0 {
		panic("no return value specified for DeletePolicies")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []policies.Policy) error); ok {
		r0 = rf(ctx, prs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePolicyFilter provides a mock function with given fields: ctx, pr
func (_m *Service) DeletePolicyFilter(ctx context.Context, pr policies.Policy) error {
	ret := _m.Called(ctx, pr)

	if len(ret) == 0 {
		panic("no return value specified for DeletePolicyFilter")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, policies.Policy) error); ok {
		r0 = rf(ctx, pr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListAllObjects provides a mock function with given fields: ctx, pr
func (_m *Service) ListAllObjects(ctx context.Context, pr policies.Policy) (policies.PolicyPage, error) {
	ret := _m.Called(ctx, pr)

	if len(ret) == 0 {
		panic("no return value specified for ListAllObjects")
	}

	var r0 policies.PolicyPage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, policies.Policy) (policies.PolicyPage, error)); ok {
		return rf(ctx, pr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, policies.Policy) policies.PolicyPage); ok {
		r0 = rf(ctx, pr)
	} else {
		r0 = ret.Get(0).(policies.PolicyPage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, policies.Policy) error); ok {
		r1 = rf(ctx, pr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllSubjects provides a mock function with given fields: ctx, pr
func (_m *Service) ListAllSubjects(ctx context.Context, pr policies.Policy) (policies.PolicyPage, error) {
	ret := _m.Called(ctx, pr)

	if len(ret) == 0 {
		panic("no return value specified for ListAllSubjects")
	}

	var r0 policies.PolicyPage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, policies.Policy) (policies.PolicyPage, error)); ok {
		return rf(ctx, pr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, policies.Policy) policies.PolicyPage); ok {
		r0 = rf(ctx, pr)
	} else {
		r0 = ret.Get(0).(policies.PolicyPage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, policies.Policy) error); ok {
		r1 = rf(ctx, pr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListObjects provides a mock function with given fields: ctx, pr, nextPageToken, limit
func (_m *Service) ListObjects(ctx context.Context, pr policies.Policy, nextPageToken string, limit uint64) (policies.PolicyPage, error) {
	ret := _m.Called(ctx, pr, nextPageToken, limit)

	if len(ret) == 0 {
		panic("no return value specified for ListObjects")
	}

	var r0 policies.PolicyPage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, policies.Policy, string, uint64) (policies.PolicyPage, error)); ok {
		return rf(ctx, pr, nextPageToken, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, policies.Policy, string, uint64) policies.PolicyPage); ok {
		r0 = rf(ctx, pr, nextPageToken, limit)
	} else {
		r0 = ret.Get(0).(policies.PolicyPage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, policies.Policy, string, uint64) error); ok {
		r1 = rf(ctx, pr, nextPageToken, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPermissions provides a mock function with given fields: ctx, pr, permissionsFilter
func (_m *Service) ListPermissions(ctx context.Context, pr policies.Policy, permissionsFilter []string) (policies.Permissions, error) {
	ret := _m.Called(ctx, pr, permissionsFilter)

	if len(ret) == 0 {
		panic("no return value specified for ListPermissions")
	}

	var r0 policies.Permissions
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, policies.Policy, []string) (policies.Permissions, error)); ok {
		return rf(ctx, pr, permissionsFilter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, policies.Policy, []string) policies.Permissions); ok {
		r0 = rf(ctx, pr, permissionsFilter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(policies.Permissions)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, policies.Policy, []string) error); ok {
		r1 = rf(ctx, pr, permissionsFilter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSubjects provides a mock function with given fields: ctx, pr, nextPageToken, limit
func (_m *Service) ListSubjects(ctx context.Context, pr policies.Policy, nextPageToken string, limit uint64) (policies.PolicyPage, error) {
	ret := _m.Called(ctx, pr, nextPageToken, limit)

	if len(ret) == 0 {
		panic("no return value specified for ListSubjects")
	}

	var r0 policies.PolicyPage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, policies.Policy, string, uint64) (policies.PolicyPage, error)); ok {
		return rf(ctx, pr, nextPageToken, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, policies.Policy, string, uint64) policies.PolicyPage); ok {
		r0 = rf(ctx, pr, nextPageToken, limit)
	} else {
		r0 = ret.Get(0).(policies.PolicyPage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, policies.Policy, string, uint64) error); ok {
		r1 = rf(ctx, pr, nextPageToken, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
