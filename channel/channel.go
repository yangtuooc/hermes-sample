package channel

import (
	"context"
	"hermes/channel/message"
)

type MessageChannels []MessageChannel

type MessageChannel interface {
	Send(ctx context.Context, message *message.Message) error // 发送消息
}

type InterceptableChannel interface {
	MessageChannel
	AddInterceptor(interceptor Interceptor) // 添加消息通道的拦截器
}

func SendWith(ctx context.Context, msg *message.Message, channel MessageChannel) error {
	return channel.Send(ctx, msg)
}
