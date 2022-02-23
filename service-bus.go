package servicebus

import (
	"context"
	"errors"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/elvenworks/service-bus-conector/internal/azure"
)

type ServiceBusConfig struct {
	ConnectionString string
	Queue            string
	Topic            string
	Subscription     string
}

type serviceBus struct {
	config   ServiceBusConfig
	client   azure.ServicBusClient
	receiver azure.ServiceBusReceiver
}

func NewServiceBus(config ServiceBusConfig) ServiceBusConnector {
	return &serviceBus{
		config: config,
	}
}

func (serviceBus *serviceBus) Connect() error {

	if serviceBus.config.ConnectionString == "" {
		return errors.New("connection string is required")
	}

	var err error
	serviceBus.client, err = azservicebus.NewClientFromConnectionString(serviceBus.config.ConnectionString, nil)
	if err != nil {
		return err
	}
	return nil
}

func (serviceBus *serviceBus) Close(timeout context.Context) error {
	return serviceBus.client.Close(timeout)
}

func (serviceBus *serviceBus) PeakFromQueue(timeout context.Context) error {

	var err error
	serviceBus.receiver, err = serviceBus.client.NewReceiverForQueue(
		serviceBus.config.Queue,
		nil,
	)

	if err != nil {
		return err
	}

	return serviceBus.peak(timeout)
}

func (serviceBus *serviceBus) PeakFromTopic(timeout context.Context) error {

	var err error
	serviceBus.receiver, err = serviceBus.client.NewReceiverForSubscription(
		serviceBus.config.Topic,
		serviceBus.config.Subscription,
		nil,
	)

	if err != nil {
		return err
	}

	return serviceBus.peak(timeout)
}

func (serviceBus *serviceBus) peak(timeout context.Context) error {

	_, err := serviceBus.receiver.PeekMessages(timeout, 1, nil)
	if err != nil {
		return err
	}

	return nil
}
