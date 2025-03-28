// Code generated by mockery v2.53.0. DO NOT EDIT.

package mocks

import (
	context "context"

	functions "github.com/smartcontractkit/chainlink/v2/core/services/functions"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

type ORM_Expecter struct {
	mock *mock.Mock
}

func (_m *ORM) EXPECT() *ORM_Expecter {
	return &ORM_Expecter{mock: &_m.Mock}
}

// CreateRequest provides a mock function with given fields: ctx, request
func (_m *ORM) CreateRequest(ctx context.Context, request *functions.Request) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for CreateRequest")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *functions.Request) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ORM_CreateRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateRequest'
type ORM_CreateRequest_Call struct {
	*mock.Call
}

// CreateRequest is a helper method to define mock.On call
//   - ctx context.Context
//   - request *functions.Request
func (_e *ORM_Expecter) CreateRequest(ctx interface{}, request interface{}) *ORM_CreateRequest_Call {
	return &ORM_CreateRequest_Call{Call: _e.mock.On("CreateRequest", ctx, request)}
}

func (_c *ORM_CreateRequest_Call) Run(run func(ctx context.Context, request *functions.Request)) *ORM_CreateRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*functions.Request))
	})
	return _c
}

func (_c *ORM_CreateRequest_Call) Return(_a0 error) *ORM_CreateRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ORM_CreateRequest_Call) RunAndReturn(run func(context.Context, *functions.Request) error) *ORM_CreateRequest_Call {
	_c.Call.Return(run)
	return _c
}

// FindById provides a mock function with given fields: ctx, requestID
func (_m *ORM) FindById(ctx context.Context, requestID functions.RequestID) (*functions.Request, error) {
	ret := _m.Called(ctx, requestID)

	if len(ret) == 0 {
		panic("no return value specified for FindById")
	}

	var r0 *functions.Request
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, functions.RequestID) (*functions.Request, error)); ok {
		return rf(ctx, requestID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, functions.RequestID) *functions.Request); ok {
		r0 = rf(ctx, requestID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*functions.Request)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, functions.RequestID) error); ok {
		r1 = rf(ctx, requestID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ORM_FindById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindById'
type ORM_FindById_Call struct {
	*mock.Call
}

// FindById is a helper method to define mock.On call
//   - ctx context.Context
//   - requestID functions.RequestID
func (_e *ORM_Expecter) FindById(ctx interface{}, requestID interface{}) *ORM_FindById_Call {
	return &ORM_FindById_Call{Call: _e.mock.On("FindById", ctx, requestID)}
}

func (_c *ORM_FindById_Call) Run(run func(ctx context.Context, requestID functions.RequestID)) *ORM_FindById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(functions.RequestID))
	})
	return _c
}

func (_c *ORM_FindById_Call) Return(_a0 *functions.Request, _a1 error) *ORM_FindById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ORM_FindById_Call) RunAndReturn(run func(context.Context, functions.RequestID) (*functions.Request, error)) *ORM_FindById_Call {
	_c.Call.Return(run)
	return _c
}

// FindOldestEntriesByState provides a mock function with given fields: ctx, state, limit
func (_m *ORM) FindOldestEntriesByState(ctx context.Context, state functions.RequestState, limit uint32) ([]functions.Request, error) {
	ret := _m.Called(ctx, state, limit)

	if len(ret) == 0 {
		panic("no return value specified for FindOldestEntriesByState")
	}

	var r0 []functions.Request
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, functions.RequestState, uint32) ([]functions.Request, error)); ok {
		return rf(ctx, state, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, functions.RequestState, uint32) []functions.Request); ok {
		r0 = rf(ctx, state, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]functions.Request)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, functions.RequestState, uint32) error); ok {
		r1 = rf(ctx, state, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ORM_FindOldestEntriesByState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindOldestEntriesByState'
type ORM_FindOldestEntriesByState_Call struct {
	*mock.Call
}

// FindOldestEntriesByState is a helper method to define mock.On call
//   - ctx context.Context
//   - state functions.RequestState
//   - limit uint32
func (_e *ORM_Expecter) FindOldestEntriesByState(ctx interface{}, state interface{}, limit interface{}) *ORM_FindOldestEntriesByState_Call {
	return &ORM_FindOldestEntriesByState_Call{Call: _e.mock.On("FindOldestEntriesByState", ctx, state, limit)}
}

func (_c *ORM_FindOldestEntriesByState_Call) Run(run func(ctx context.Context, state functions.RequestState, limit uint32)) *ORM_FindOldestEntriesByState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(functions.RequestState), args[2].(uint32))
	})
	return _c
}

func (_c *ORM_FindOldestEntriesByState_Call) Return(_a0 []functions.Request, _a1 error) *ORM_FindOldestEntriesByState_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ORM_FindOldestEntriesByState_Call) RunAndReturn(run func(context.Context, functions.RequestState, uint32) ([]functions.Request, error)) *ORM_FindOldestEntriesByState_Call {
	_c.Call.Return(run)
	return _c
}

// PruneOldestRequests provides a mock function with given fields: ctx, maxRequestsInDB, batchSize
func (_m *ORM) PruneOldestRequests(ctx context.Context, maxRequestsInDB uint32, batchSize uint32) (uint32, uint32, error) {
	ret := _m.Called(ctx, maxRequestsInDB, batchSize)

	if len(ret) == 0 {
		panic("no return value specified for PruneOldestRequests")
	}

	var r0 uint32
	var r1 uint32
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint32) (uint32, uint32, error)); ok {
		return rf(ctx, maxRequestsInDB, batchSize)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint32) uint32); ok {
		r0 = rf(ctx, maxRequestsInDB, batchSize)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32, uint32) uint32); ok {
		r1 = rf(ctx, maxRequestsInDB, batchSize)
	} else {
		r1 = ret.Get(1).(uint32)
	}

	if rf, ok := ret.Get(2).(func(context.Context, uint32, uint32) error); ok {
		r2 = rf(ctx, maxRequestsInDB, batchSize)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ORM_PruneOldestRequests_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PruneOldestRequests'
type ORM_PruneOldestRequests_Call struct {
	*mock.Call
}

// PruneOldestRequests is a helper method to define mock.On call
//   - ctx context.Context
//   - maxRequestsInDB uint32
//   - batchSize uint32
func (_e *ORM_Expecter) PruneOldestRequests(ctx interface{}, maxRequestsInDB interface{}, batchSize interface{}) *ORM_PruneOldestRequests_Call {
	return &ORM_PruneOldestRequests_Call{Call: _e.mock.On("PruneOldestRequests", ctx, maxRequestsInDB, batchSize)}
}

func (_c *ORM_PruneOldestRequests_Call) Run(run func(ctx context.Context, maxRequestsInDB uint32, batchSize uint32)) *ORM_PruneOldestRequests_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(uint32))
	})
	return _c
}

func (_c *ORM_PruneOldestRequests_Call) Return(total uint32, pruned uint32, err error) *ORM_PruneOldestRequests_Call {
	_c.Call.Return(total, pruned, err)
	return _c
}

func (_c *ORM_PruneOldestRequests_Call) RunAndReturn(run func(context.Context, uint32, uint32) (uint32, uint32, error)) *ORM_PruneOldestRequests_Call {
	_c.Call.Return(run)
	return _c
}

// SetConfirmed provides a mock function with given fields: ctx, requestID
func (_m *ORM) SetConfirmed(ctx context.Context, requestID functions.RequestID) error {
	ret := _m.Called(ctx, requestID)

	if len(ret) == 0 {
		panic("no return value specified for SetConfirmed")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, functions.RequestID) error); ok {
		r0 = rf(ctx, requestID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ORM_SetConfirmed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetConfirmed'
type ORM_SetConfirmed_Call struct {
	*mock.Call
}

// SetConfirmed is a helper method to define mock.On call
//   - ctx context.Context
//   - requestID functions.RequestID
func (_e *ORM_Expecter) SetConfirmed(ctx interface{}, requestID interface{}) *ORM_SetConfirmed_Call {
	return &ORM_SetConfirmed_Call{Call: _e.mock.On("SetConfirmed", ctx, requestID)}
}

func (_c *ORM_SetConfirmed_Call) Run(run func(ctx context.Context, requestID functions.RequestID)) *ORM_SetConfirmed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(functions.RequestID))
	})
	return _c
}

func (_c *ORM_SetConfirmed_Call) Return(_a0 error) *ORM_SetConfirmed_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ORM_SetConfirmed_Call) RunAndReturn(run func(context.Context, functions.RequestID) error) *ORM_SetConfirmed_Call {
	_c.Call.Return(run)
	return _c
}

// SetError provides a mock function with given fields: ctx, requestID, errorType, computationError, readyAt, readyForProcessing
func (_m *ORM) SetError(ctx context.Context, requestID functions.RequestID, errorType functions.ErrType, computationError []byte, readyAt time.Time, readyForProcessing bool) error {
	ret := _m.Called(ctx, requestID, errorType, computationError, readyAt, readyForProcessing)

	if len(ret) == 0 {
		panic("no return value specified for SetError")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, functions.RequestID, functions.ErrType, []byte, time.Time, bool) error); ok {
		r0 = rf(ctx, requestID, errorType, computationError, readyAt, readyForProcessing)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ORM_SetError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetError'
type ORM_SetError_Call struct {
	*mock.Call
}

// SetError is a helper method to define mock.On call
//   - ctx context.Context
//   - requestID functions.RequestID
//   - errorType functions.ErrType
//   - computationError []byte
//   - readyAt time.Time
//   - readyForProcessing bool
func (_e *ORM_Expecter) SetError(ctx interface{}, requestID interface{}, errorType interface{}, computationError interface{}, readyAt interface{}, readyForProcessing interface{}) *ORM_SetError_Call {
	return &ORM_SetError_Call{Call: _e.mock.On("SetError", ctx, requestID, errorType, computationError, readyAt, readyForProcessing)}
}

func (_c *ORM_SetError_Call) Run(run func(ctx context.Context, requestID functions.RequestID, errorType functions.ErrType, computationError []byte, readyAt time.Time, readyForProcessing bool)) *ORM_SetError_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(functions.RequestID), args[2].(functions.ErrType), args[3].([]byte), args[4].(time.Time), args[5].(bool))
	})
	return _c
}

func (_c *ORM_SetError_Call) Return(_a0 error) *ORM_SetError_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ORM_SetError_Call) RunAndReturn(run func(context.Context, functions.RequestID, functions.ErrType, []byte, time.Time, bool) error) *ORM_SetError_Call {
	_c.Call.Return(run)
	return _c
}

// SetFinalized provides a mock function with given fields: ctx, requestID, reportedResult, reportedError
func (_m *ORM) SetFinalized(ctx context.Context, requestID functions.RequestID, reportedResult []byte, reportedError []byte) error {
	ret := _m.Called(ctx, requestID, reportedResult, reportedError)

	if len(ret) == 0 {
		panic("no return value specified for SetFinalized")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, functions.RequestID, []byte, []byte) error); ok {
		r0 = rf(ctx, requestID, reportedResult, reportedError)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ORM_SetFinalized_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetFinalized'
type ORM_SetFinalized_Call struct {
	*mock.Call
}

// SetFinalized is a helper method to define mock.On call
//   - ctx context.Context
//   - requestID functions.RequestID
//   - reportedResult []byte
//   - reportedError []byte
func (_e *ORM_Expecter) SetFinalized(ctx interface{}, requestID interface{}, reportedResult interface{}, reportedError interface{}) *ORM_SetFinalized_Call {
	return &ORM_SetFinalized_Call{Call: _e.mock.On("SetFinalized", ctx, requestID, reportedResult, reportedError)}
}

func (_c *ORM_SetFinalized_Call) Run(run func(ctx context.Context, requestID functions.RequestID, reportedResult []byte, reportedError []byte)) *ORM_SetFinalized_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(functions.RequestID), args[2].([]byte), args[3].([]byte))
	})
	return _c
}

func (_c *ORM_SetFinalized_Call) Return(_a0 error) *ORM_SetFinalized_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ORM_SetFinalized_Call) RunAndReturn(run func(context.Context, functions.RequestID, []byte, []byte) error) *ORM_SetFinalized_Call {
	_c.Call.Return(run)
	return _c
}

// SetResult provides a mock function with given fields: ctx, requestID, computationResult, readyAt
func (_m *ORM) SetResult(ctx context.Context, requestID functions.RequestID, computationResult []byte, readyAt time.Time) error {
	ret := _m.Called(ctx, requestID, computationResult, readyAt)

	if len(ret) == 0 {
		panic("no return value specified for SetResult")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, functions.RequestID, []byte, time.Time) error); ok {
		r0 = rf(ctx, requestID, computationResult, readyAt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ORM_SetResult_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetResult'
type ORM_SetResult_Call struct {
	*mock.Call
}

// SetResult is a helper method to define mock.On call
//   - ctx context.Context
//   - requestID functions.RequestID
//   - computationResult []byte
//   - readyAt time.Time
func (_e *ORM_Expecter) SetResult(ctx interface{}, requestID interface{}, computationResult interface{}, readyAt interface{}) *ORM_SetResult_Call {
	return &ORM_SetResult_Call{Call: _e.mock.On("SetResult", ctx, requestID, computationResult, readyAt)}
}

func (_c *ORM_SetResult_Call) Run(run func(ctx context.Context, requestID functions.RequestID, computationResult []byte, readyAt time.Time)) *ORM_SetResult_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(functions.RequestID), args[2].([]byte), args[3].(time.Time))
	})
	return _c
}

func (_c *ORM_SetResult_Call) Return(_a0 error) *ORM_SetResult_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ORM_SetResult_Call) RunAndReturn(run func(context.Context, functions.RequestID, []byte, time.Time) error) *ORM_SetResult_Call {
	_c.Call.Return(run)
	return _c
}

// TimeoutExpiredResults provides a mock function with given fields: ctx, cutoff, limit
func (_m *ORM) TimeoutExpiredResults(ctx context.Context, cutoff time.Time, limit uint32) ([]functions.RequestID, error) {
	ret := _m.Called(ctx, cutoff, limit)

	if len(ret) == 0 {
		panic("no return value specified for TimeoutExpiredResults")
	}

	var r0 []functions.RequestID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, uint32) ([]functions.RequestID, error)); ok {
		return rf(ctx, cutoff, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, uint32) []functions.RequestID); ok {
		r0 = rf(ctx, cutoff, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]functions.RequestID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, time.Time, uint32) error); ok {
		r1 = rf(ctx, cutoff, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ORM_TimeoutExpiredResults_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TimeoutExpiredResults'
type ORM_TimeoutExpiredResults_Call struct {
	*mock.Call
}

// TimeoutExpiredResults is a helper method to define mock.On call
//   - ctx context.Context
//   - cutoff time.Time
//   - limit uint32
func (_e *ORM_Expecter) TimeoutExpiredResults(ctx interface{}, cutoff interface{}, limit interface{}) *ORM_TimeoutExpiredResults_Call {
	return &ORM_TimeoutExpiredResults_Call{Call: _e.mock.On("TimeoutExpiredResults", ctx, cutoff, limit)}
}

func (_c *ORM_TimeoutExpiredResults_Call) Run(run func(ctx context.Context, cutoff time.Time, limit uint32)) *ORM_TimeoutExpiredResults_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(time.Time), args[2].(uint32))
	})
	return _c
}

func (_c *ORM_TimeoutExpiredResults_Call) Return(_a0 []functions.RequestID, _a1 error) *ORM_TimeoutExpiredResults_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ORM_TimeoutExpiredResults_Call) RunAndReturn(run func(context.Context, time.Time, uint32) ([]functions.RequestID, error)) *ORM_TimeoutExpiredResults_Call {
	_c.Call.Return(run)
	return _c
}

// NewORM creates a new instance of ORM. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewORM(t interface {
	mock.TestingT
	Cleanup(func())
}) *ORM {
	mock := &ORM{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
