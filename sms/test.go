package sms

import (
	"context"
	"encoding/json"
	"fmt"
	"hermes/channel"
	"hermes/channel/message"
)

var _ channel.Vendor = (*test)(nil)

type test struct {
	listeners channel.Listeners
}

func (t *test) Id() string {
	return "test"
}

func (t *test) Type() string {
	return "sms"
}

func (t *test) Name() string {
	return "test"
}

func (t *test) Description() string {
	return "test"
}

func (t *test) AddListener(listener channel.VendorListener) {
	t.listeners = append(t.listeners, listener)
}

func (t *test) Send(ctx context.Context, message *message.Message) error {
	marshal, _ := json.Marshal(message)
	jsonStr := string(marshal)
	t.listeners.OnRequest(ctx, message, t, jsonStr)
	result := fmt.Sprintf("send sms: %s\n", jsonStr)
	fmt.Println(result)
	t.listeners.OnResponse(ctx, message, t, result, nil)
	return nil
}

func Test() channel.Vendor {
	return &test{}
}
