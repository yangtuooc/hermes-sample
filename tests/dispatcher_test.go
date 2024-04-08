package tests

import (
	"context"
	"hermes/channel"
	"hermes/channel/message"
	"hermes/sms"
	"testing"
)

// case 1: create a new dispatcher without any channels
func TestDispatcherCase1(t *testing.T) {
	factory := channel.NewFactory()
	dispatcher := channel.NewDispatcher(factory)
	if dispatcher == nil {
		t.Error("expected dispatcher to be created")
	}
}

// case 2: create a new dispatcher without any channels and try to send a message
func TestDispatcherCase2(t *testing.T) {
	t.Run("not specified channel", func(t *testing.T) {
		factory := channel.NewFactory()
		dispatcher := channel.NewDispatcher(factory)

		ctx := context.Background()
		msg := &message.Message{Payload: "test"}
		err := dispatcher.Send(ctx, msg)
		if err == nil {
			t.Error("expected error to be returned")
		}
		t.Log(err)
	})

	t.Run("specified channel", func(t *testing.T) {
		factory := channel.NewFactory()
		dispatcher := channel.NewDispatcher(factory)

		ctx := context.Background()
		msg := message.New("test")
		msg.SetChannel("test")

		err := dispatcher.Send(ctx, msg)
		if err == nil {
			t.Error("expected error to be returned")
		}
		t.Log(err)
	})
}

// case 3: create a new dispatcher with a channel and try to send a message
func TestDispatcherCase3(t *testing.T) {
	smsChannel := sms.NewChannel()
	vendor := sms.Test()
	vendor.AddListener(channel.LoggerVendorListener)

	if err := smsChannel.Register(vendor); err != nil {
		t.Error(err)
	}
	factory := channel.NewFactory(smsChannel)
	dispatcher := channel.NewDispatcher(factory)

	ctx := context.Background()
	msg := message.New("test")
	msg.SetChannel(smsChannel.Id())

	err := dispatcher.Send(ctx, msg)
	if err != nil {
		t.Error(err)
	}
}
