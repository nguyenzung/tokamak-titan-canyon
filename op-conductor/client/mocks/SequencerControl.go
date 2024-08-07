// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	eth "github.com/tokamak-network/tokamak-thanos/op-service/eth"

	mock "github.com/stretchr/testify/mock"
)

// SequencerControl is an autogenerated mock type for the SequencerControl type
type SequencerControl struct {
	mock.Mock
}

type SequencerControl_Expecter struct {
	mock *mock.Mock
}

func (_m *SequencerControl) EXPECT() *SequencerControl_Expecter {
	return &SequencerControl_Expecter{mock: &_m.Mock}
}

// LatestUnsafeBlock provides a mock function with given fields: ctx
func (_m *SequencerControl) LatestUnsafeBlock(ctx context.Context) (eth.BlockInfo, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for LatestUnsafeBlock")
	}

	var r0 eth.BlockInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (eth.BlockInfo, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) eth.BlockInfo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(eth.BlockInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SequencerControl_LatestUnsafeBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LatestUnsafeBlock'
type SequencerControl_LatestUnsafeBlock_Call struct {
	*mock.Call
}

// LatestUnsafeBlock is a helper method to define mock.On call
//   - ctx context.Context
func (_e *SequencerControl_Expecter) LatestUnsafeBlock(ctx interface{}) *SequencerControl_LatestUnsafeBlock_Call {
	return &SequencerControl_LatestUnsafeBlock_Call{Call: _e.mock.On("LatestUnsafeBlock", ctx)}
}

func (_c *SequencerControl_LatestUnsafeBlock_Call) Run(run func(ctx context.Context)) *SequencerControl_LatestUnsafeBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *SequencerControl_LatestUnsafeBlock_Call) Return(_a0 eth.BlockInfo, _a1 error) *SequencerControl_LatestUnsafeBlock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SequencerControl_LatestUnsafeBlock_Call) RunAndReturn(run func(context.Context) (eth.BlockInfo, error)) *SequencerControl_LatestUnsafeBlock_Call {
	_c.Call.Return(run)
	return _c
}

// PostUnsafePayload provides a mock function with given fields: ctx, payload
func (_m *SequencerControl) PostUnsafePayload(ctx context.Context, payload *eth.ExecutionPayloadEnvelope) error {
	ret := _m.Called(ctx, payload)

	if len(ret) == 0 {
		panic("no return value specified for PostUnsafePayload")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *eth.ExecutionPayloadEnvelope) error); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SequencerControl_PostUnsafePayload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PostUnsafePayload'
type SequencerControl_PostUnsafePayload_Call struct {
	*mock.Call
}

// PostUnsafePayload is a helper method to define mock.On call
//   - ctx context.Context
//   - payload *eth.ExecutionPayloadEnvelope
func (_e *SequencerControl_Expecter) PostUnsafePayload(ctx interface{}, payload interface{}) *SequencerControl_PostUnsafePayload_Call {
	return &SequencerControl_PostUnsafePayload_Call{Call: _e.mock.On("PostUnsafePayload", ctx, payload)}
}

func (_c *SequencerControl_PostUnsafePayload_Call) Run(run func(ctx context.Context, payload *eth.ExecutionPayloadEnvelope)) *SequencerControl_PostUnsafePayload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*eth.ExecutionPayloadEnvelope))
	})
	return _c
}

func (_c *SequencerControl_PostUnsafePayload_Call) Return(_a0 error) *SequencerControl_PostUnsafePayload_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SequencerControl_PostUnsafePayload_Call) RunAndReturn(run func(context.Context, *eth.ExecutionPayloadEnvelope) error) *SequencerControl_PostUnsafePayload_Call {
	_c.Call.Return(run)
	return _c
}

// SequencerActive provides a mock function with given fields: ctx
func (_m *SequencerControl) SequencerActive(ctx context.Context) (bool, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for SequencerActive")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (bool, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SequencerControl_SequencerActive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SequencerActive'
type SequencerControl_SequencerActive_Call struct {
	*mock.Call
}

// SequencerActive is a helper method to define mock.On call
//   - ctx context.Context
func (_e *SequencerControl_Expecter) SequencerActive(ctx interface{}) *SequencerControl_SequencerActive_Call {
	return &SequencerControl_SequencerActive_Call{Call: _e.mock.On("SequencerActive", ctx)}
}

func (_c *SequencerControl_SequencerActive_Call) Run(run func(ctx context.Context)) *SequencerControl_SequencerActive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *SequencerControl_SequencerActive_Call) Return(_a0 bool, _a1 error) *SequencerControl_SequencerActive_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SequencerControl_SequencerActive_Call) RunAndReturn(run func(context.Context) (bool, error)) *SequencerControl_SequencerActive_Call {
	_c.Call.Return(run)
	return _c
}

// StartSequencer provides a mock function with given fields: ctx, hash
func (_m *SequencerControl) StartSequencer(ctx context.Context, hash common.Hash) error {
	ret := _m.Called(ctx, hash)

	if len(ret) == 0 {
		panic("no return value specified for StartSequencer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) error); ok {
		r0 = rf(ctx, hash)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SequencerControl_StartSequencer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StartSequencer'
type SequencerControl_StartSequencer_Call struct {
	*mock.Call
}

// StartSequencer is a helper method to define mock.On call
//   - ctx context.Context
//   - hash common.Hash
func (_e *SequencerControl_Expecter) StartSequencer(ctx interface{}, hash interface{}) *SequencerControl_StartSequencer_Call {
	return &SequencerControl_StartSequencer_Call{Call: _e.mock.On("StartSequencer", ctx, hash)}
}

func (_c *SequencerControl_StartSequencer_Call) Run(run func(ctx context.Context, hash common.Hash)) *SequencerControl_StartSequencer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Hash))
	})
	return _c
}

func (_c *SequencerControl_StartSequencer_Call) Return(_a0 error) *SequencerControl_StartSequencer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SequencerControl_StartSequencer_Call) RunAndReturn(run func(context.Context, common.Hash) error) *SequencerControl_StartSequencer_Call {
	_c.Call.Return(run)
	return _c
}

// StopSequencer provides a mock function with given fields: ctx
func (_m *SequencerControl) StopSequencer(ctx context.Context) (common.Hash, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for StopSequencer")
	}

	var r0 common.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (common.Hash, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) common.Hash); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SequencerControl_StopSequencer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StopSequencer'
type SequencerControl_StopSequencer_Call struct {
	*mock.Call
}

// StopSequencer is a helper method to define mock.On call
//   - ctx context.Context
func (_e *SequencerControl_Expecter) StopSequencer(ctx interface{}) *SequencerControl_StopSequencer_Call {
	return &SequencerControl_StopSequencer_Call{Call: _e.mock.On("StopSequencer", ctx)}
}

func (_c *SequencerControl_StopSequencer_Call) Run(run func(ctx context.Context)) *SequencerControl_StopSequencer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *SequencerControl_StopSequencer_Call) Return(_a0 common.Hash, _a1 error) *SequencerControl_StopSequencer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SequencerControl_StopSequencer_Call) RunAndReturn(run func(context.Context) (common.Hash, error)) *SequencerControl_StopSequencer_Call {
	_c.Call.Return(run)
	return _c
}

// NewSequencerControl creates a new instance of SequencerControl. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSequencerControl(t interface {
	mock.TestingT
	Cleanup(func())
}) *SequencerControl {
	mock := &SequencerControl{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}