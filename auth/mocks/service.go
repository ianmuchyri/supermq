// Code generated by mockery v2.46.3. DO NOT EDIT.

// Copyright (c) Abstract Machines

package mocks

import (
	context "context"

	auth "github.com/absmach/supermq/auth"

	mock "github.com/stretchr/testify/mock"

	policies "github.com/absmach/supermq/pkg/policies"

	time "time"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// AddScope provides a mock function with given fields: ctx, token, patID, scopes
func (_m *Service) AddScope(ctx context.Context, token string, patID string, scopes []auth.Scope) error {
	ret := _m.Called(ctx, token, patID, scopes)

	if len(ret) == 0 {
		panic("no return value specified for AddScope")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []auth.Scope) error); ok {
		r0 = rf(ctx, token, patID, scopes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Authorize provides a mock function with given fields: ctx, pr
func (_m *Service) Authorize(ctx context.Context, pr policies.Policy) error {
	ret := _m.Called(ctx, pr)

	if len(ret) == 0 {
		panic("no return value specified for Authorize")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, policies.Policy) error); ok {
		r0 = rf(ctx, pr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuthorizePAT provides a mock function with given fields: ctx, userID, patID, entityType, optionalDomainID, operation, entityID
func (_m *Service) AuthorizePAT(ctx context.Context, userID string, patID string, entityType auth.EntityType, optionalDomainID string, operation auth.Operation, entityID string) error {
	ret := _m.Called(ctx, userID, patID, entityType, optionalDomainID, operation, entityID)

	if len(ret) == 0 {
		panic("no return value specified for AuthorizePAT")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, auth.EntityType, string, auth.Operation, string) error); ok {
		r0 = rf(ctx, userID, patID, entityType, optionalDomainID, operation, entityID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreatePAT provides a mock function with given fields: ctx, token, name, description, duration
func (_m *Service) CreatePAT(ctx context.Context, token string, name string, description string, duration time.Duration) (auth.PAT, error) {
	ret := _m.Called(ctx, token, name, description, duration)

	if len(ret) == 0 {
		panic("no return value specified for CreatePAT")
	}

	var r0 auth.PAT
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, time.Duration) (auth.PAT, error)); ok {
		return rf(ctx, token, name, description, duration)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, time.Duration) auth.PAT); ok {
		r0 = rf(ctx, token, name, description, duration)
	} else {
		r0 = ret.Get(0).(auth.PAT)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, time.Duration) error); ok {
		r1 = rf(ctx, token, name, description, duration)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePAT provides a mock function with given fields: ctx, token, patID
func (_m *Service) DeletePAT(ctx context.Context, token string, patID string) error {
	ret := _m.Called(ctx, token, patID)

	if len(ret) == 0 {
		panic("no return value specified for DeletePAT")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, token, patID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Identify provides a mock function with given fields: ctx, token
func (_m *Service) Identify(ctx context.Context, token string) (auth.Key, error) {
	ret := _m.Called(ctx, token)

	if len(ret) == 0 {
		panic("no return value specified for Identify")
	}

	var r0 auth.Key
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (auth.Key, error)); ok {
		return rf(ctx, token)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) auth.Key); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Get(0).(auth.Key)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IdentifyPAT provides a mock function with given fields: ctx, paToken
func (_m *Service) IdentifyPAT(ctx context.Context, paToken string) (auth.PAT, error) {
	ret := _m.Called(ctx, paToken)

	if len(ret) == 0 {
		panic("no return value specified for IdentifyPAT")
	}

	var r0 auth.PAT
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (auth.PAT, error)); ok {
		return rf(ctx, paToken)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) auth.PAT); ok {
		r0 = rf(ctx, paToken)
	} else {
		r0 = ret.Get(0).(auth.PAT)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, paToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Issue provides a mock function with given fields: ctx, token, key
func (_m *Service) Issue(ctx context.Context, token string, key auth.Key) (auth.Token, error) {
	ret := _m.Called(ctx, token, key)

	if len(ret) == 0 {
		panic("no return value specified for Issue")
	}

	var r0 auth.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, auth.Key) (auth.Token, error)); ok {
		return rf(ctx, token, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, auth.Key) auth.Token); ok {
		r0 = rf(ctx, token, key)
	} else {
		r0 = ret.Get(0).(auth.Token)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, auth.Key) error); ok {
		r1 = rf(ctx, token, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPATS provides a mock function with given fields: ctx, token, pm
func (_m *Service) ListPATS(ctx context.Context, token string, pm auth.PATSPageMeta) (auth.PATSPage, error) {
	ret := _m.Called(ctx, token, pm)

	if len(ret) == 0 {
		panic("no return value specified for ListPATS")
	}

	var r0 auth.PATSPage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, auth.PATSPageMeta) (auth.PATSPage, error)); ok {
		return rf(ctx, token, pm)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, auth.PATSPageMeta) auth.PATSPage); ok {
		r0 = rf(ctx, token, pm)
	} else {
		r0 = ret.Get(0).(auth.PATSPage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, auth.PATSPageMeta) error); ok {
		r1 = rf(ctx, token, pm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListScopes provides a mock function with given fields: ctx, token, pm
func (_m *Service) ListScopes(ctx context.Context, token string, pm auth.ScopesPageMeta) (auth.ScopesPage, error) {
	ret := _m.Called(ctx, token, pm)

	if len(ret) == 0 {
		panic("no return value specified for ListScopes")
	}

	var r0 auth.ScopesPage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, auth.ScopesPageMeta) (auth.ScopesPage, error)); ok {
		return rf(ctx, token, pm)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, auth.ScopesPageMeta) auth.ScopesPage); ok {
		r0 = rf(ctx, token, pm)
	} else {
		r0 = ret.Get(0).(auth.ScopesPage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, auth.ScopesPageMeta) error); ok {
		r1 = rf(ctx, token, pm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveAllPAT provides a mock function with given fields: ctx, token
func (_m *Service) RemoveAllPAT(ctx context.Context, token string) error {
	ret := _m.Called(ctx, token)

	if len(ret) == 0 {
		panic("no return value specified for RemoveAllPAT")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemovePATAllScope provides a mock function with given fields: ctx, token, patID
func (_m *Service) RemovePATAllScope(ctx context.Context, token string, patID string) error {
	ret := _m.Called(ctx, token, patID)

	if len(ret) == 0 {
		panic("no return value specified for RemovePATAllScope")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, token, patID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveScope provides a mock function with given fields: ctx, token, patID, scopeIDs
func (_m *Service) RemoveScope(ctx context.Context, token string, patID string, scopeIDs ...string) error {
	_va := make([]interface{}, len(scopeIDs))
	for _i := range scopeIDs {
		_va[_i] = scopeIDs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, token, patID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RemoveScope")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...string) error); ok {
		r0 = rf(ctx, token, patID, scopeIDs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResetPATSecret provides a mock function with given fields: ctx, token, patID, duration
func (_m *Service) ResetPATSecret(ctx context.Context, token string, patID string, duration time.Duration) (auth.PAT, error) {
	ret := _m.Called(ctx, token, patID, duration)

	if len(ret) == 0 {
		panic("no return value specified for ResetPATSecret")
	}

	var r0 auth.PAT
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, time.Duration) (auth.PAT, error)); ok {
		return rf(ctx, token, patID, duration)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, time.Duration) auth.PAT); ok {
		r0 = rf(ctx, token, patID, duration)
	} else {
		r0 = ret.Get(0).(auth.PAT)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, time.Duration) error); ok {
		r1 = rf(ctx, token, patID, duration)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveKey provides a mock function with given fields: ctx, token, id
func (_m *Service) RetrieveKey(ctx context.Context, token string, id string) (auth.Key, error) {
	ret := _m.Called(ctx, token, id)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveKey")
	}

	var r0 auth.Key
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (auth.Key, error)); ok {
		return rf(ctx, token, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) auth.Key); ok {
		r0 = rf(ctx, token, id)
	} else {
		r0 = ret.Get(0).(auth.Key)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, token, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrievePAT provides a mock function with given fields: ctx, userID, patID
func (_m *Service) RetrievePAT(ctx context.Context, userID string, patID string) (auth.PAT, error) {
	ret := _m.Called(ctx, userID, patID)

	if len(ret) == 0 {
		panic("no return value specified for RetrievePAT")
	}

	var r0 auth.PAT
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (auth.PAT, error)); ok {
		return rf(ctx, userID, patID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) auth.PAT); ok {
		r0 = rf(ctx, userID, patID)
	} else {
		r0 = ret.Get(0).(auth.PAT)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, userID, patID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Revoke provides a mock function with given fields: ctx, token, id
func (_m *Service) Revoke(ctx context.Context, token string, id string) error {
	ret := _m.Called(ctx, token, id)

	if len(ret) == 0 {
		panic("no return value specified for Revoke")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, token, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RevokePATSecret provides a mock function with given fields: ctx, token, patID
func (_m *Service) RevokePATSecret(ctx context.Context, token string, patID string) error {
	ret := _m.Called(ctx, token, patID)

	if len(ret) == 0 {
		panic("no return value specified for RevokePATSecret")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, token, patID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePATDescription provides a mock function with given fields: ctx, token, patID, description
func (_m *Service) UpdatePATDescription(ctx context.Context, token string, patID string, description string) (auth.PAT, error) {
	ret := _m.Called(ctx, token, patID, description)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePATDescription")
	}

	var r0 auth.PAT
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (auth.PAT, error)); ok {
		return rf(ctx, token, patID, description)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) auth.PAT); ok {
		r0 = rf(ctx, token, patID, description)
	} else {
		r0 = ret.Get(0).(auth.PAT)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, token, patID, description)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePATName provides a mock function with given fields: ctx, token, patID, name
func (_m *Service) UpdatePATName(ctx context.Context, token string, patID string, name string) (auth.PAT, error) {
	ret := _m.Called(ctx, token, patID, name)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePATName")
	}

	var r0 auth.PAT
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (auth.PAT, error)); ok {
		return rf(ctx, token, patID, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) auth.PAT); ok {
		r0 = rf(ctx, token, patID, name)
	} else {
		r0 = ret.Get(0).(auth.PAT)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, token, patID, name)
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
