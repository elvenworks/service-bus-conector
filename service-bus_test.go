package servicebus

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/elvenworks/service-bus-conector/internal/azure"
	azuremocks "github.com/elvenworks/service-bus-conector/internal/azure/mocks"
)

func TestNewServiceBus(t *testing.T) {
	type args struct {
		config ServiceBusConfig
	}
	tests := []struct {
		name string
		args args
		want ServiceBusConnector
	}{
		{
			name: "Success",
			args: args{
				config: ServiceBusConfig{},
			},
			want: &serviceBus{
				config: ServiceBusConfig{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServiceBus(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServiceBus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceBus_Connect(t *testing.T) {
	type fields struct {
		config   ServiceBusConfig
		client   *azuremocks.ServicBusClient
		receiver *azuremocks.ServiceBusReceiver
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Error creating client",
			fields: fields{
				config: ServiceBusConfig{ConnectionString: ""},
			},
			wantErr: true,
		},
		{
			name: "Internal Error creating client",
			fields: fields{
				config: ServiceBusConfig{ConnectionString: "qqr"},
				client: &azuremocks.ServicBusClient{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceBus := &serviceBus{
				config:   tt.fields.config,
				client:   tt.fields.client,
				receiver: tt.fields.receiver,
			}
			if err := serviceBus.Connect(); (err != nil) != tt.wantErr {
				t.Errorf("serviceBus.Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_serviceBus_Close(t *testing.T) {
	type fields struct {
		config   ServiceBusConfig
		client   *azuremocks.ServicBusClient
		receiver *azuremocks.ServiceBusReceiver
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Error closing",
			fields: fields{
				config: ServiceBusConfig{},
				client: &azuremocks.ServicBusClient{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
			defer cancel()

			tt.fields.client.On("Close", ctx).Return(errors.New("err"))

			serviceBus := &serviceBus{
				config:   tt.fields.config,
				client:   tt.fields.client,
				receiver: tt.fields.receiver,
			}
			if err := serviceBus.Close(ctx); (err != nil) != tt.wantErr {
				t.Errorf("serviceBus.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_serviceBus_PeakFromQueue(t *testing.T) {
	type fields struct {
		config   ServiceBusConfig
		client   *azuremocks.ServicBusClient
		receiver *azuremocks.ServiceBusReceiver
	}
	type args struct {
		timeout context.Context
	}
	tests := []struct {
		name                     string
		fields                   fields
		args                     args
		wantErr                  bool
		errorNewReceiverForQueue error
	}{

		{
			name: "error creating receiver",
			fields: fields{
				config:   ServiceBusConfig{},
				client:   &azuremocks.ServicBusClient{},
				receiver: &azuremocks.ServiceBusReceiver{},
			},
			args:                     args{},
			errorNewReceiverForQueue: errors.New("err"),
			wantErr:                  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var receiveroptions *azservicebus.ReceiverOptions
			tt.fields.client.On("NewReceiverForQueue", tt.fields.config.Queue, receiveroptions).Return(&azservicebus.Receiver{}, tt.errorNewReceiverForQueue)

			serviceBus := &serviceBus{
				config:   tt.fields.config,
				client:   tt.fields.client,
				receiver: tt.fields.receiver,
			}
			if err := serviceBus.PeakFromQueue(tt.args.timeout); (err != nil) != tt.wantErr {
				t.Errorf("serviceBus.PeakFromQueue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_serviceBus_PeakFromTopic(t *testing.T) {
	type fields struct {
		config   ServiceBusConfig
		client   *azuremocks.ServicBusClient
		receiver azure.ServiceBusReceiver
	}
	type args struct {
		timeout context.Context
	}
	tests := []struct {
		name                            string
		fields                          fields
		args                            args
		wantErr                         bool
		errorNewReceiverForSubscription error
	}{

		{
			name: "error creating receiver",
			fields: fields{
				config: ServiceBusConfig{},
				client: &azuremocks.ServicBusClient{},
			},
			args:                            args{},
			errorNewReceiverForSubscription: errors.New("err"),
			wantErr:                         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var receiveroptions *azservicebus.ReceiverOptions
			tt.fields.client.On("NewReceiverForSubscription", tt.fields.config.Topic, tt.fields.config.Subscription, receiveroptions).Return(tt.fields.receiver, tt.errorNewReceiverForSubscription)

			serviceBus := &serviceBus{
				config:   tt.fields.config,
				client:   tt.fields.client,
				receiver: tt.fields.receiver,
			}
			if err := serviceBus.PeakFromTopic(tt.args.timeout); (err != nil) != tt.wantErr {
				t.Errorf("serviceBus.PeakFromQueue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_serviceBus_peak(t *testing.T) {
	type fields struct {
		config   ServiceBusConfig
		client   *azuremocks.ServicBusClient
		receiver *azuremocks.ServiceBusReceiver
	}
	type args struct {
		timeout context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Error peaking message",
			fields: fields{
				receiver: &azuremocks.ServiceBusReceiver{},
			},
			args: args{
				timeout: context.TODO(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var peakMessageoptions *azservicebus.PeekMessagesOptions
			tt.fields.receiver.On("PeekMessages", tt.args.timeout, 1, peakMessageoptions).Return([]*azservicebus.ReceivedMessage{}, errors.New("err"))

			serviceBus := &serviceBus{
				config:   tt.fields.config,
				client:   tt.fields.client,
				receiver: tt.fields.receiver,
			}
			if err := serviceBus.peak(tt.args.timeout); (err != nil) != tt.wantErr {
				t.Errorf("serviceBus.peak() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
