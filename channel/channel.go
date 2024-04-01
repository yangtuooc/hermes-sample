package channel

import (
	"context"
	"hermes/channel/message"
)

type MessageChannels []MessageChannel

type MessageChannel interface {
	Send(ctx context.Context, message *message.Message) error
}

type InterceptableChannel interface {
	MessageChannel
	AddInterceptor(interceptor Interceptor)
}

func SendWith(ctx context.Context, msg *message.Message, channel MessageChannel) error {
	return channel.Send(ctx, msg)
}
