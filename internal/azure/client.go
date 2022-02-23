package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

type ServicBusClient interface {
	NewReceiverForQueue(queueName string, options *azservicebus.ReceiverOptions) (*azservicebus.Receiver, error)
	NewReceiverForSubscription(topicName string, subscriptionName string, options *azservicebus.ReceiverOptions) (*azservicebus.Receiver, error)
	NewSender(queueOrTopic string, options *azservicebus.NewSenderOptions) (*azservicebus.Sender, error)
	AcceptSessionForQueue(ctx context.Context, queueName string, sessionID string, options *azservicebus.SessionReceiverOptions) (*azservicebus.SessionReceiver, error)
	AcceptSessionForSubscription(ctx context.Context, topicName string, subscriptionName string, sessionID string, options *azservicebus.SessionReceiverOptions) (*azservicebus.SessionReceiver, error)
	AcceptNextSessionForQueue(ctx context.Context, queueName string, options *azservicebus.SessionReceiverOptions) (*azservicebus.SessionReceiver, error)
	AcceptNextSessionForSubscription(ctx context.Context, topicName string, subscriptionName string, options *azservicebus.SessionReceiverOptions) (*azservicebus.SessionReceiver, error)
	Close(ctx context.Context) error
}
