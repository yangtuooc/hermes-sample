package tests

import (
	"context"
	"hermes/channel"
	"hermes/channel/message"
)

var _ channel.Vendor = (*c)(nil)

type c struct {
	listeners channel.Listeners
}

func (c *c) Id() string {
	return "c"
}

func (c *c) Type() string {
	return "sms"
}

func (c *c) Name() string {
	return "c"
}

func (c *c) Description() string {
	return "c"
}

func (c *c) AddListener(listener channel.VendorListener) {
	c.listeners = append(c.listeners, listener)
}

func (c *c) Send(ctx context.Context, message *message.Message) error {
	c.listeners.OnRequest(ctx, message, c, nil)
	return nil
}

func testC() channel.Vendor {
	return &c{}
}
