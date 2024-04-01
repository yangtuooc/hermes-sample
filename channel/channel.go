package channel

import (
	"context"
	"encoding/json"
	"hermes/message"
)

type MessageChannel interface {
	Send(ctx context.Context, message *message.Message) error
}

var Console = func(ctx context.Context, msg *message.Message) error {
	marshal, _ := json.Marshal(msg)
	println(string(marshal))
	return nil
}

func SendWith(ctx context.Context, msg *message.Message, channel MessageChannel) error {
	return channel.Send(ctx, msg)
}
