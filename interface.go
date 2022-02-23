package servicebus

import "context"

type ServiceBusConnector interface {
	Connect() error
	Close(timeout context.Context) error
	PeakFromQueue(timeout context.Context) error
	PeakFromTopic(timeout context.Context) error
}
