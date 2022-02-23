// Code generated by mockery v1.0.0. DO NOT EDIT.

package azuremocks

import azservicebus "github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
import context "context"
import mock "github.com/stretchr/testify/mock"

// ServiceBusReceiver is an autogenerated mock type for the ServiceBusReceiver type
type ServiceBusReceiver struct {
	mock.Mock
}

// AbandonMessage provides a mock function with given fields: ctx, message, options
func (_m *ServiceBusReceiver) AbandonMessage(ctx context.Context, message *azservicebus.ReceivedMessage, options *azservicebus.AbandonMessageOptions) error {
	ret := _m.Called(ctx, message, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *azservicebus.ReceivedMessage, *azservicebus.AbandonMessageOptions) error); ok {
		r0 = rf(ctx, message, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields: ctx
func (_m *ServiceBusReceiver) Close(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CompleteMessage provides a mock function with given fields: ctx, message
func (_m *ServiceBusReceiver) CompleteMessage(ctx context.Context, message *azservicebus.ReceivedMessage) error {
	ret := _m.Called(ctx, message)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *azservicebus.ReceivedMessage) error); ok {
		r0 = rf(ctx, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeadLetterMessage provides a mock function with given fields: ctx, message, options
func (_m *ServiceBusReceiver) DeadLetterMessage(ctx context.Context, message *azservicebus.ReceivedMessage, options *azservicebus.DeadLetterOptions) error {
	ret := _m.Called(ctx, message, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *azservicebus.ReceivedMessage, *azservicebus.DeadLetterOptions) error); ok {
		r0 = rf(ctx, message, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeferMessage provides a mock function with given fields: ctx, message, options
func (_m *ServiceBusReceiver) DeferMessage(ctx context.Context, message *azservicebus.ReceivedMessage, options *azservicebus.DeferMessageOptions) error {
	ret := _m.Called(ctx, message, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *azservicebus.ReceivedMessage, *azservicebus.DeferMessageOptions) error); ok {
		r0 = rf(ctx, message, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PeekMessages provides a mock function with given fields: ctx, maxMessageCount, options
func (_m *ServiceBusReceiver) PeekMessages(ctx context.Context, maxMessageCount int, options *azservicebus.PeekMessagesOptions) ([]*azservicebus.ReceivedMessage, error) {
	ret := _m.Called(ctx, maxMessageCount, options)

	var r0 []*azservicebus.ReceivedMessage
	if rf, ok := ret.Get(0).(func(context.Context, int, *azservicebus.PeekMessagesOptions) []*azservicebus.ReceivedMessage); ok {
		r0 = rf(ctx, maxMessageCount, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*azservicebus.ReceivedMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, *azservicebus.PeekMessagesOptions) error); ok {
		r1 = rf(ctx, maxMessageCount, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReceiveDeferredMessages provides a mock function with given fields: ctx, sequenceNumbers
func (_m *ServiceBusReceiver) ReceiveDeferredMessages(ctx context.Context, sequenceNumbers []int64) ([]*azservicebus.ReceivedMessage, error) {
	ret := _m.Called(ctx, sequenceNumbers)

	var r0 []*azservicebus.ReceivedMessage
	if rf, ok := ret.Get(0).(func(context.Context, []int64) []*azservicebus.ReceivedMessage); ok {
		r0 = rf(ctx, sequenceNumbers)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*azservicebus.ReceivedMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []int64) error); ok {
		r1 = rf(ctx, sequenceNumbers)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReceiveMessages provides a mock function with given fields: ctx, maxMessages, options
func (_m *ServiceBusReceiver) ReceiveMessages(ctx context.Context, maxMessages int, options *azservicebus.ReceiveMessagesOptions) ([]*azservicebus.ReceivedMessage, error) {
	ret := _m.Called(ctx, maxMessages, options)

	var r0 []*azservicebus.ReceivedMessage
	if rf, ok := ret.Get(0).(func(context.Context, int, *azservicebus.ReceiveMessagesOptions) []*azservicebus.ReceivedMessage); ok {
		r0 = rf(ctx, maxMessages, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*azservicebus.ReceivedMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, *azservicebus.ReceiveMessagesOptions) error); ok {
		r1 = rf(ctx, maxMessages, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RenewMessageLock provides a mock function with given fields: ctx, msg
func (_m *ServiceBusReceiver) RenewMessageLock(ctx context.Context, msg *azservicebus.ReceivedMessage) error {
	ret := _m.Called(ctx, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *azservicebus.ReceivedMessage) error); ok {
		r0 = rf(ctx, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
