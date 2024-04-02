package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"hermes/channel"
	"hermes/channel/message"
)

var _ channel.Vendor = (*b)(nil)

type b struct {
	listeners channel.Listeners
}

func (b *b) Id() string {
	return "b"
}

func (b *b) Type() string {
	return "sms"
}

func (b *b) Name() string {
	return "b channel"
}

func (b *b) Description() string {
	return "b channel"
}

func (b *b) AddListener(listener channel.VendorListener) {
}

func (b *b) Send(ctx context.Context, message *message.Message) error {
	marshal, _ := json.Marshal(message)
	jsonStr := string(marshal)
	fmt.Println(jsonStr)
	b.listeners.OnResponse(ctx, message, b, jsonStr, nil)
	return nil
}

func testB() channel.Vendor {
	return &b{}
}
