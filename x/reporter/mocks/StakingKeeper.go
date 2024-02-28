// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"

	address "cosmossdk.io/core/address"

	cosmos_sdktypes "github.com/cosmos/cosmos-sdk/types"

	math "cosmossdk.io/math"

	mock "github.com/stretchr/testify/mock"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// StakingKeeper is an autogenerated mock type for the StakingKeeper type
type StakingKeeper struct {
	mock.Mock
}

// ConsensusAddressCodec provides a mock function with given fields:
func (_m *StakingKeeper) ConsensusAddressCodec() address.Codec {
	ret := _m.Called()

	var r0 address.Codec
	if rf, ok := ret.Get(0).(func() address.Codec); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(address.Codec)
		}
	}

	return r0
}

// Delegate provides a mock function with given fields: ctx, delAddr, bondAmt, tokenSrc, validator, subtractAccount
func (_m *StakingKeeper) Delegate(ctx context.Context, delAddr cosmos_sdktypes.AccAddress, bondAmt math.Int, tokenSrc stakingtypes.BondStatus, validator stakingtypes.Validator, subtractAccount bool) (math.LegacyDec, error) {
	ret := _m.Called(ctx, delAddr, bondAmt, tokenSrc, validator, subtractAccount)

	var r0 math.LegacyDec
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, cosmos_sdktypes.AccAddress, math.Int, stakingtypes.BondStatus, stakingtypes.Validator, bool) (math.LegacyDec, error)); ok {
		return rf(ctx, delAddr, bondAmt, tokenSrc, validator, subtractAccount)
	}
	if rf, ok := ret.Get(0).(func(context.Context, cosmos_sdktypes.AccAddress, math.Int, stakingtypes.BondStatus, stakingtypes.Validator, bool) math.LegacyDec); ok {
		r0 = rf(ctx, delAddr, bondAmt, tokenSrc, validator, subtractAccount)
	} else {
		r0 = ret.Get(0).(math.LegacyDec)
	}

	if rf, ok := ret.Get(1).(func(context.Context, cosmos_sdktypes.AccAddress, math.Int, stakingtypes.BondStatus, stakingtypes.Validator, bool) error); ok {
		r1 = rf(ctx, delAddr, bondAmt, tokenSrc, validator, subtractAccount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delegation provides a mock function with given fields: _a0, _a1, _a2
func (_m *StakingKeeper) Delegation(_a0 context.Context, _a1 cosmos_sdktypes.AccAddress, _a2 cosmos_sdktypes.ValAddress) (stakingtypes.DelegationI, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 stakingtypes.DelegationI
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, cosmos_sdktypes.AccAddress, cosmos_sdktypes.ValAddress) (stakingtypes.DelegationI, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, cosmos_sdktypes.AccAddress, cosmos_sdktypes.ValAddress) stakingtypes.DelegationI); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(stakingtypes.DelegationI)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, cosmos_sdktypes.AccAddress, cosmos_sdktypes.ValAddress) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetValidator provides a mock function with given fields: ctx, addr
func (_m *StakingKeeper) GetValidator(ctx context.Context, addr cosmos_sdktypes.ValAddress) (stakingtypes.Validator, error) {
	ret := _m.Called(ctx, addr)

	var r0 stakingtypes.Validator
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, cosmos_sdktypes.ValAddress) (stakingtypes.Validator, error)); ok {
		return rf(ctx, addr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, cosmos_sdktypes.ValAddress) stakingtypes.Validator); ok {
		r0 = rf(ctx, addr)
	} else {
		r0 = ret.Get(0).(stakingtypes.Validator)
	}

	if rf, ok := ret.Get(1).(func(context.Context, cosmos_sdktypes.ValAddress) error); ok {
		r1 = rf(ctx, addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatorByConsAddr provides a mock function with given fields: _a0, _a1
func (_m *StakingKeeper) ValidatorByConsAddr(_a0 context.Context, _a1 cosmos_sdktypes.ConsAddress) (stakingtypes.ValidatorI, error) {
	ret := _m.Called(_a0, _a1)

	var r0 stakingtypes.ValidatorI
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, cosmos_sdktypes.ConsAddress) (stakingtypes.ValidatorI, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, cosmos_sdktypes.ConsAddress) stakingtypes.ValidatorI); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(stakingtypes.ValidatorI)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, cosmos_sdktypes.ConsAddress) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewStakingKeeper interface {
	mock.TestingT
	Cleanup(func())
}

// NewStakingKeeper creates a new instance of StakingKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStakingKeeper(t mockConstructorTestingTNewStakingKeeper) *StakingKeeper {
	mock := &StakingKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
