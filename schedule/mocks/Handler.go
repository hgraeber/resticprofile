// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	calendar "github.com/creativeprojects/resticprofile/calendar"
	mock "github.com/stretchr/testify/mock"

	schedule "github.com/creativeprojects/resticprofile/schedule"
)

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

type Handler_Expecter struct {
	mock *mock.Mock
}

func (_m *Handler) EXPECT() *Handler_Expecter {
	return &Handler_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *Handler) Close() {
	_m.Called()
}

// Handler_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type Handler_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *Handler_Expecter) Close() *Handler_Close_Call {
	return &Handler_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *Handler_Close_Call) Run(run func()) *Handler_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Handler_Close_Call) Return() *Handler_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *Handler_Close_Call) RunAndReturn(run func()) *Handler_Close_Call {
	_c.Call.Return(run)
	return _c
}

// CreateJob provides a mock function with given fields: job, schedules, permission
func (_m *Handler) CreateJob(job *schedule.Config, schedules []*calendar.Event, permission string) error {
	ret := _m.Called(job, schedules, permission)

	if len(ret) == 0 {
		panic("no return value specified for CreateJob")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*schedule.Config, []*calendar.Event, string) error); ok {
		r0 = rf(job, schedules, permission)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Handler_CreateJob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateJob'
type Handler_CreateJob_Call struct {
	*mock.Call
}

// CreateJob is a helper method to define mock.On call
//   - job *schedule.Config
//   - schedules []*calendar.Event
//   - permission string
func (_e *Handler_Expecter) CreateJob(job interface{}, schedules interface{}, permission interface{}) *Handler_CreateJob_Call {
	return &Handler_CreateJob_Call{Call: _e.mock.On("CreateJob", job, schedules, permission)}
}

func (_c *Handler_CreateJob_Call) Run(run func(job *schedule.Config, schedules []*calendar.Event, permission string)) *Handler_CreateJob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*schedule.Config), args[1].([]*calendar.Event), args[2].(string))
	})
	return _c
}

func (_c *Handler_CreateJob_Call) Return(_a0 error) *Handler_CreateJob_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Handler_CreateJob_Call) RunAndReturn(run func(*schedule.Config, []*calendar.Event, string) error) *Handler_CreateJob_Call {
	_c.Call.Return(run)
	return _c
}

// DisplayJobStatus provides a mock function with given fields: job
func (_m *Handler) DisplayJobStatus(job *schedule.Config) error {
	ret := _m.Called(job)

	if len(ret) == 0 {
		panic("no return value specified for DisplayJobStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*schedule.Config) error); ok {
		r0 = rf(job)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Handler_DisplayJobStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DisplayJobStatus'
type Handler_DisplayJobStatus_Call struct {
	*mock.Call
}

// DisplayJobStatus is a helper method to define mock.On call
//   - job *schedule.Config
func (_e *Handler_Expecter) DisplayJobStatus(job interface{}) *Handler_DisplayJobStatus_Call {
	return &Handler_DisplayJobStatus_Call{Call: _e.mock.On("DisplayJobStatus", job)}
}

func (_c *Handler_DisplayJobStatus_Call) Run(run func(job *schedule.Config)) *Handler_DisplayJobStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*schedule.Config))
	})
	return _c
}

func (_c *Handler_DisplayJobStatus_Call) Return(_a0 error) *Handler_DisplayJobStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Handler_DisplayJobStatus_Call) RunAndReturn(run func(*schedule.Config) error) *Handler_DisplayJobStatus_Call {
	_c.Call.Return(run)
	return _c
}

// DisplayParsedSchedules provides a mock function with given fields: command, events
func (_m *Handler) DisplayParsedSchedules(command string, events []*calendar.Event) {
	_m.Called(command, events)
}

// Handler_DisplayParsedSchedules_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DisplayParsedSchedules'
type Handler_DisplayParsedSchedules_Call struct {
	*mock.Call
}

// DisplayParsedSchedules is a helper method to define mock.On call
//   - command string
//   - events []*calendar.Event
func (_e *Handler_Expecter) DisplayParsedSchedules(command interface{}, events interface{}) *Handler_DisplayParsedSchedules_Call {
	return &Handler_DisplayParsedSchedules_Call{Call: _e.mock.On("DisplayParsedSchedules", command, events)}
}

func (_c *Handler_DisplayParsedSchedules_Call) Run(run func(command string, events []*calendar.Event)) *Handler_DisplayParsedSchedules_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]*calendar.Event))
	})
	return _c
}

func (_c *Handler_DisplayParsedSchedules_Call) Return() *Handler_DisplayParsedSchedules_Call {
	_c.Call.Return()
	return _c
}

func (_c *Handler_DisplayParsedSchedules_Call) RunAndReturn(run func(string, []*calendar.Event)) *Handler_DisplayParsedSchedules_Call {
	_c.Call.Return(run)
	return _c
}

// DisplaySchedules provides a mock function with given fields: command, schedules
func (_m *Handler) DisplaySchedules(command string, schedules []string) error {
	ret := _m.Called(command, schedules)

	if len(ret) == 0 {
		panic("no return value specified for DisplaySchedules")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(command, schedules)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Handler_DisplaySchedules_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DisplaySchedules'
type Handler_DisplaySchedules_Call struct {
	*mock.Call
}

// DisplaySchedules is a helper method to define mock.On call
//   - command string
//   - schedules []string
func (_e *Handler_Expecter) DisplaySchedules(command interface{}, schedules interface{}) *Handler_DisplaySchedules_Call {
	return &Handler_DisplaySchedules_Call{Call: _e.mock.On("DisplaySchedules", command, schedules)}
}

func (_c *Handler_DisplaySchedules_Call) Run(run func(command string, schedules []string)) *Handler_DisplaySchedules_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]string))
	})
	return _c
}

func (_c *Handler_DisplaySchedules_Call) Return(_a0 error) *Handler_DisplaySchedules_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Handler_DisplaySchedules_Call) RunAndReturn(run func(string, []string) error) *Handler_DisplaySchedules_Call {
	_c.Call.Return(run)
	return _c
}

// DisplayStatus provides a mock function with given fields: profileName
func (_m *Handler) DisplayStatus(profileName string) error {
	ret := _m.Called(profileName)

	if len(ret) == 0 {
		panic("no return value specified for DisplayStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(profileName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Handler_DisplayStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DisplayStatus'
type Handler_DisplayStatus_Call struct {
	*mock.Call
}

// DisplayStatus is a helper method to define mock.On call
//   - profileName string
func (_e *Handler_Expecter) DisplayStatus(profileName interface{}) *Handler_DisplayStatus_Call {
	return &Handler_DisplayStatus_Call{Call: _e.mock.On("DisplayStatus", profileName)}
}

func (_c *Handler_DisplayStatus_Call) Run(run func(profileName string)) *Handler_DisplayStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Handler_DisplayStatus_Call) Return(_a0 error) *Handler_DisplayStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Handler_DisplayStatus_Call) RunAndReturn(run func(string) error) *Handler_DisplayStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields:
func (_m *Handler) Init() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Init")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Handler_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type Handler_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
func (_e *Handler_Expecter) Init() *Handler_Init_Call {
	return &Handler_Init_Call{Call: _e.mock.On("Init")}
}

func (_c *Handler_Init_Call) Run(run func()) *Handler_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Handler_Init_Call) Return(_a0 error) *Handler_Init_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Handler_Init_Call) RunAndReturn(run func() error) *Handler_Init_Call {
	_c.Call.Return(run)
	return _c
}

// ParseSchedules provides a mock function with given fields: schedules
func (_m *Handler) ParseSchedules(schedules []string) ([]*calendar.Event, error) {
	ret := _m.Called(schedules)

	if len(ret) == 0 {
		panic("no return value specified for ParseSchedules")
	}

	var r0 []*calendar.Event
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) ([]*calendar.Event, error)); ok {
		return rf(schedules)
	}
	if rf, ok := ret.Get(0).(func([]string) []*calendar.Event); ok {
		r0 = rf(schedules)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*calendar.Event)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(schedules)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Handler_ParseSchedules_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ParseSchedules'
type Handler_ParseSchedules_Call struct {
	*mock.Call
}

// ParseSchedules is a helper method to define mock.On call
//   - schedules []string
func (_e *Handler_Expecter) ParseSchedules(schedules interface{}) *Handler_ParseSchedules_Call {
	return &Handler_ParseSchedules_Call{Call: _e.mock.On("ParseSchedules", schedules)}
}

func (_c *Handler_ParseSchedules_Call) Run(run func(schedules []string)) *Handler_ParseSchedules_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *Handler_ParseSchedules_Call) Return(_a0 []*calendar.Event, _a1 error) *Handler_ParseSchedules_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Handler_ParseSchedules_Call) RunAndReturn(run func([]string) ([]*calendar.Event, error)) *Handler_ParseSchedules_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveJob provides a mock function with given fields: job, permission
func (_m *Handler) RemoveJob(job *schedule.Config, permission string) error {
	ret := _m.Called(job, permission)

	if len(ret) == 0 {
		panic("no return value specified for RemoveJob")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*schedule.Config, string) error); ok {
		r0 = rf(job, permission)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Handler_RemoveJob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveJob'
type Handler_RemoveJob_Call struct {
	*mock.Call
}

// RemoveJob is a helper method to define mock.On call
//   - job *schedule.Config
//   - permission string
func (_e *Handler_Expecter) RemoveJob(job interface{}, permission interface{}) *Handler_RemoveJob_Call {
	return &Handler_RemoveJob_Call{Call: _e.mock.On("RemoveJob", job, permission)}
}

func (_c *Handler_RemoveJob_Call) Run(run func(job *schedule.Config, permission string)) *Handler_RemoveJob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*schedule.Config), args[1].(string))
	})
	return _c
}

func (_c *Handler_RemoveJob_Call) Return(_a0 error) *Handler_RemoveJob_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Handler_RemoveJob_Call) RunAndReturn(run func(*schedule.Config, string) error) *Handler_RemoveJob_Call {
	_c.Call.Return(run)
	return _c
}

// NewHandler creates a new instance of Handler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *Handler {
	mock := &Handler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
