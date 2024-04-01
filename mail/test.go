package mail

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
	return "test-mail"
}

func (t *test) Type() string {
	return "mail"
}

func (t *test) Name() string {
	return "test-mail"
}

func (t *test) AddListener(listener channel.VendorListener) {
	t.listeners = append(t.listeners, listener)
}

func (t *test) Send(ctx context.Context, message *message.Message) error {
	marshal, _ := json.Marshal(message)
	jsonStr := string(marshal)
	t.listeners.OnRequest(ctx, message, t, jsonStr)
	result := fmt.Sprintln("send mail: ", jsonStr)
	fmt.Println(result)
	t.listeners.OnResponse(ctx, message, t, result, nil)
	return nil
}

func Test() channel.Vendor {
	return &test{}
}
