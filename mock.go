package servicebus

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ServiceBusConnectorMock is an autogenerated mock type for the ServiceBusConnectorMock type
type ServiceBusConnectorMock struct {
	mock.Mock
}

// Close provides a mock function with given fields: _a0
func (_m *ServiceBusConnectorMock) Close(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Connect provides a mock function with given fields:
func (_m *ServiceBusConnectorMock) Connect() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PeakFromQueue provides a mock function with given fields: _a0
func (_m *ServiceBusConnectorMock) PeakFromQueue(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PeakFromTopic provides a mock function with given fields: _a0
func (_m *ServiceBusConnectorMock) PeakFromTopic(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
