package tests

import (
	"context"
	"errors"
	"hermes/channel"
	"hermes/channel/message"
)

var _ channel.Vendor = (*a)(nil)

type a struct {
	listeners channel.Listeners
}

func (a *a) Id() string {
	return "a"
}

func (a *a) Type() string {
	return "sms"
}

func (a *a) Name() string {
	return "a channel"
}

func (a *a) AddListener(listener channel.VendorListener) {
	a.listeners = append(a.listeners, listener)
}

func (a *a) Send(ctx context.Context, message *message.Message) error {
	err := errors.New("a channel is not available")
	a.listeners.OnResponse(ctx, message, a, nil, err)
	return err
}

func testA() channel.Vendor {
	return &a{}
}
