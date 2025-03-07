// Code generated by mockery v2.46.3. DO NOT EDIT.

// Copyright (c) Abstract Machines

package mocks

import (
	context "context"

	policies "github.com/absmach/supermq/pkg/policies"
	mock "github.com/stretchr/testify/mock"

	roles "github.com/absmach/supermq/pkg/roles"
)

// Provisioner is an autogenerated mock type for the Provisioner type
type Provisioner struct {
	mock.Mock
}

// AddNewEntitiesRoles provides a mock function with given fields: ctx, domainID, userID, entityIDs, optionalEntityPolicies, newBuiltInRoleMembers
func (_m *Provisioner) AddNewEntitiesRoles(ctx context.Context, domainID string, userID string, entityIDs []string, optionalEntityPolicies []policies.Policy, newBuiltInRoleMembers map[roles.BuiltInRoleName][]roles.Member) ([]roles.RoleProvision, error) {
	ret := _m.Called(ctx, domainID, userID, entityIDs, optionalEntityPolicies, newBuiltInRoleMembers)

	if len(ret) == 0 {
		panic("no return value specified for AddNewEntitiesRoles")
	}

	var r0 []roles.RoleProvision
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []string, []policies.Policy, map[roles.BuiltInRoleName][]roles.Member) ([]roles.RoleProvision, error)); ok {
		return rf(ctx, domainID, userID, entityIDs, optionalEntityPolicies, newBuiltInRoleMembers)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []string, []policies.Policy, map[roles.BuiltInRoleName][]roles.Member) []roles.RoleProvision); ok {
		r0 = rf(ctx, domainID, userID, entityIDs, optionalEntityPolicies, newBuiltInRoleMembers)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]roles.RoleProvision)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, []string, []policies.Policy, map[roles.BuiltInRoleName][]roles.Member) error); ok {
		r1 = rf(ctx, domainID, userID, entityIDs, optionalEntityPolicies, newBuiltInRoleMembers)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveEntitiesRoles provides a mock function with given fields: ctx, domainID, userID, entityIDs, optionalFilterDeletePolicies, optionalDeletePolicies
func (_m *Provisioner) RemoveEntitiesRoles(ctx context.Context, domainID string, userID string, entityIDs []string, optionalFilterDeletePolicies []policies.Policy, optionalDeletePolicies []policies.Policy) error {
	ret := _m.Called(ctx, domainID, userID, entityIDs, optionalFilterDeletePolicies, optionalDeletePolicies)

	if len(ret) == 0 {
		panic("no return value specified for RemoveEntitiesRoles")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []string, []policies.Policy, []policies.Policy) error); ok {
		r0 = rf(ctx, domainID, userID, entityIDs, optionalFilterDeletePolicies, optionalDeletePolicies)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewProvisioner creates a new instance of Provisioner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProvisioner(t interface {
	mock.TestingT
	Cleanup(func())
}) *Provisioner {
	mock := &Provisioner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
