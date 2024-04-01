package console

import (
	"context"
	"encoding/json"
	"fmt"
	"hermes/channel"
	"hermes/channel/message"
)

var _ channel.MessageChannel = (*console)(nil)

type console struct{}

func (c *console) Send(ctx context.Context, message *message.Message) error {
	marshal, _ := json.Marshal(message)
	fmt.Println(string(marshal))
	return nil
}

func New() channel.MessageChannel {
	return &console{}
}
