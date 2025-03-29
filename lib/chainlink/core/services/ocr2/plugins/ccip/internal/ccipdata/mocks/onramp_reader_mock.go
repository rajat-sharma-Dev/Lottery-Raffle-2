// Code generated by mockery v2.53.0. DO NOT EDIT.

package mocks

import (
	ccip "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// OnRampReader is an autogenerated mock type for the OnRampReader type
type OnRampReader struct {
	mock.Mock
}

type OnRampReader_Expecter struct {
	mock *mock.Mock
}

func (_m *OnRampReader) EXPECT() *OnRampReader_Expecter {
	return &OnRampReader_Expecter{mock: &_m.Mock}
}

// Address provides a mock function with given fields: ctx
func (_m *OnRampReader) Address(ctx context.Context) (ccip.Address, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Address")
	}

	var r0 ccip.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (ccip.Address, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) ccip.Address); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(ccip.Address)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OnRampReader_Address_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Address'
type OnRampReader_Address_Call struct {
	*mock.Call
}

// Address is a helper method to define mock.On call
//   - ctx context.Context
func (_e *OnRampReader_Expecter) Address(ctx interface{}) *OnRampReader_Address_Call {
	return &OnRampReader_Address_Call{Call: _e.mock.On("Address", ctx)}
}

func (_c *OnRampReader_Address_Call) Run(run func(ctx context.Context)) *OnRampReader_Address_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *OnRampReader_Address_Call) Return(_a0 ccip.Address, _a1 error) *OnRampReader_Address_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OnRampReader_Address_Call) RunAndReturn(run func(context.Context) (ccip.Address, error)) *OnRampReader_Address_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with no fields
func (_m *OnRampReader) Close() error {
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

// OnRampReader_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type OnRampReader_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *OnRampReader_Expecter) Close() *OnRampReader_Close_Call {
	return &OnRampReader_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *OnRampReader_Close_Call) Run(run func()) *OnRampReader_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *OnRampReader_Close_Call) Return(_a0 error) *OnRampReader_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OnRampReader_Close_Call) RunAndReturn(run func() error) *OnRampReader_Close_Call {
	_c.Call.Return(run)
	return _c
}

// GetDynamicConfig provides a mock function with given fields: ctx
func (_m *OnRampReader) GetDynamicConfig(ctx context.Context) (ccip.OnRampDynamicConfig, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetDynamicConfig")
	}

	var r0 ccip.OnRampDynamicConfig
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (ccip.OnRampDynamicConfig, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) ccip.OnRampDynamicConfig); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(ccip.OnRampDynamicConfig)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OnRampReader_GetDynamicConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDynamicConfig'
type OnRampReader_GetDynamicConfig_Call struct {
	*mock.Call
}

// GetDynamicConfig is a helper method to define mock.On call
//   - ctx context.Context
func (_e *OnRampReader_Expecter) GetDynamicConfig(ctx interface{}) *OnRampReader_GetDynamicConfig_Call {
	return &OnRampReader_GetDynamicConfig_Call{Call: _e.mock.On("GetDynamicConfig", ctx)}
}

func (_c *OnRampReader_GetDynamicConfig_Call) Run(run func(ctx context.Context)) *OnRampReader_GetDynamicConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *OnRampReader_GetDynamicConfig_Call) Return(_a0 ccip.OnRampDynamicConfig, _a1 error) *OnRampReader_GetDynamicConfig_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OnRampReader_GetDynamicConfig_Call) RunAndReturn(run func(context.Context) (ccip.OnRampDynamicConfig, error)) *OnRampReader_GetDynamicConfig_Call {
	_c.Call.Return(run)
	return _c
}

// GetSendRequestsBetweenSeqNums provides a mock function with given fields: ctx, seqNumMin, seqNumMax, finalized
func (_m *OnRampReader) GetSendRequestsBetweenSeqNums(ctx context.Context, seqNumMin uint64, seqNumMax uint64, finalized bool) ([]ccip.EVM2EVMMessageWithTxMeta, error) {
	ret := _m.Called(ctx, seqNumMin, seqNumMax, finalized)

	if len(ret) == 0 {
		panic("no return value specified for GetSendRequestsBetweenSeqNums")
	}

	var r0 []ccip.EVM2EVMMessageWithTxMeta
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, bool) ([]ccip.EVM2EVMMessageWithTxMeta, error)); ok {
		return rf(ctx, seqNumMin, seqNumMax, finalized)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, bool) []ccip.EVM2EVMMessageWithTxMeta); ok {
		r0 = rf(ctx, seqNumMin, seqNumMax, finalized)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ccip.EVM2EVMMessageWithTxMeta)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, bool) error); ok {
		r1 = rf(ctx, seqNumMin, seqNumMax, finalized)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OnRampReader_GetSendRequestsBetweenSeqNums_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSendRequestsBetweenSeqNums'
type OnRampReader_GetSendRequestsBetweenSeqNums_Call struct {
	*mock.Call
}

// GetSendRequestsBetweenSeqNums is a helper method to define mock.On call
//   - ctx context.Context
//   - seqNumMin uint64
//   - seqNumMax uint64
//   - finalized bool
func (_e *OnRampReader_Expecter) GetSendRequestsBetweenSeqNums(ctx interface{}, seqNumMin interface{}, seqNumMax interface{}, finalized interface{}) *OnRampReader_GetSendRequestsBetweenSeqNums_Call {
	return &OnRampReader_GetSendRequestsBetweenSeqNums_Call{Call: _e.mock.On("GetSendRequestsBetweenSeqNums", ctx, seqNumMin, seqNumMax, finalized)}
}

func (_c *OnRampReader_GetSendRequestsBetweenSeqNums_Call) Run(run func(ctx context.Context, seqNumMin uint64, seqNumMax uint64, finalized bool)) *OnRampReader_GetSendRequestsBetweenSeqNums_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(uint64), args[3].(bool))
	})
	return _c
}

func (_c *OnRampReader_GetSendRequestsBetweenSeqNums_Call) Return(_a0 []ccip.EVM2EVMMessageWithTxMeta, _a1 error) *OnRampReader_GetSendRequestsBetweenSeqNums_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OnRampReader_GetSendRequestsBetweenSeqNums_Call) RunAndReturn(run func(context.Context, uint64, uint64, bool) ([]ccip.EVM2EVMMessageWithTxMeta, error)) *OnRampReader_GetSendRequestsBetweenSeqNums_Call {
	_c.Call.Return(run)
	return _c
}

// IsSourceChainHealthy provides a mock function with given fields: ctx
func (_m *OnRampReader) IsSourceChainHealthy(ctx context.Context) (bool, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for IsSourceChainHealthy")
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

// OnRampReader_IsSourceChainHealthy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsSourceChainHealthy'
type OnRampReader_IsSourceChainHealthy_Call struct {
	*mock.Call
}

// IsSourceChainHealthy is a helper method to define mock.On call
//   - ctx context.Context
func (_e *OnRampReader_Expecter) IsSourceChainHealthy(ctx interface{}) *OnRampReader_IsSourceChainHealthy_Call {
	return &OnRampReader_IsSourceChainHealthy_Call{Call: _e.mock.On("IsSourceChainHealthy", ctx)}
}

func (_c *OnRampReader_IsSourceChainHealthy_Call) Run(run func(ctx context.Context)) *OnRampReader_IsSourceChainHealthy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *OnRampReader_IsSourceChainHealthy_Call) Return(_a0 bool, _a1 error) *OnRampReader_IsSourceChainHealthy_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OnRampReader_IsSourceChainHealthy_Call) RunAndReturn(run func(context.Context) (bool, error)) *OnRampReader_IsSourceChainHealthy_Call {
	_c.Call.Return(run)
	return _c
}

// IsSourceCursed provides a mock function with given fields: ctx
func (_m *OnRampReader) IsSourceCursed(ctx context.Context) (bool, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for IsSourceCursed")
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

// OnRampReader_IsSourceCursed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsSourceCursed'
type OnRampReader_IsSourceCursed_Call struct {
	*mock.Call
}

// IsSourceCursed is a helper method to define mock.On call
//   - ctx context.Context
func (_e *OnRampReader_Expecter) IsSourceCursed(ctx interface{}) *OnRampReader_IsSourceCursed_Call {
	return &OnRampReader_IsSourceCursed_Call{Call: _e.mock.On("IsSourceCursed", ctx)}
}

func (_c *OnRampReader_IsSourceCursed_Call) Run(run func(ctx context.Context)) *OnRampReader_IsSourceCursed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *OnRampReader_IsSourceCursed_Call) Return(_a0 bool, _a1 error) *OnRampReader_IsSourceCursed_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OnRampReader_IsSourceCursed_Call) RunAndReturn(run func(context.Context) (bool, error)) *OnRampReader_IsSourceCursed_Call {
	_c.Call.Return(run)
	return _c
}

// RouterAddress provides a mock function with given fields: _a0
func (_m *OnRampReader) RouterAddress(_a0 context.Context) (ccip.Address, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for RouterAddress")
	}

	var r0 ccip.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (ccip.Address, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) ccip.Address); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(ccip.Address)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OnRampReader_RouterAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RouterAddress'
type OnRampReader_RouterAddress_Call struct {
	*mock.Call
}

// RouterAddress is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *OnRampReader_Expecter) RouterAddress(_a0 interface{}) *OnRampReader_RouterAddress_Call {
	return &OnRampReader_RouterAddress_Call{Call: _e.mock.On("RouterAddress", _a0)}
}

func (_c *OnRampReader_RouterAddress_Call) Run(run func(_a0 context.Context)) *OnRampReader_RouterAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *OnRampReader_RouterAddress_Call) Return(_a0 ccip.Address, _a1 error) *OnRampReader_RouterAddress_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OnRampReader_RouterAddress_Call) RunAndReturn(run func(context.Context) (ccip.Address, error)) *OnRampReader_RouterAddress_Call {
	_c.Call.Return(run)
	return _c
}

// SourcePriceRegistryAddress provides a mock function with given fields: ctx
func (_m *OnRampReader) SourcePriceRegistryAddress(ctx context.Context) (ccip.Address, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for SourcePriceRegistryAddress")
	}

	var r0 ccip.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (ccip.Address, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) ccip.Address); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(ccip.Address)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OnRampReader_SourcePriceRegistryAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SourcePriceRegistryAddress'
type OnRampReader_SourcePriceRegistryAddress_Call struct {
	*mock.Call
}

// SourcePriceRegistryAddress is a helper method to define mock.On call
//   - ctx context.Context
func (_e *OnRampReader_Expecter) SourcePriceRegistryAddress(ctx interface{}) *OnRampReader_SourcePriceRegistryAddress_Call {
	return &OnRampReader_SourcePriceRegistryAddress_Call{Call: _e.mock.On("SourcePriceRegistryAddress", ctx)}
}

func (_c *OnRampReader_SourcePriceRegistryAddress_Call) Run(run func(ctx context.Context)) *OnRampReader_SourcePriceRegistryAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *OnRampReader_SourcePriceRegistryAddress_Call) Return(_a0 ccip.Address, _a1 error) *OnRampReader_SourcePriceRegistryAddress_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OnRampReader_SourcePriceRegistryAddress_Call) RunAndReturn(run func(context.Context) (ccip.Address, error)) *OnRampReader_SourcePriceRegistryAddress_Call {
	_c.Call.Return(run)
	return _c
}

// NewOnRampReader creates a new instance of OnRampReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOnRampReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *OnRampReader {
	mock := &OnRampReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
