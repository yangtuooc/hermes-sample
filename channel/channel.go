package channel

import (
	"context"
	"hermes/channel/message"
)

type MessageChannels []MessageChannel

// MessageChannel 消息通道，提供基础的消息发送能力
type MessageChannel interface {
	Send(ctx context.Context, message *message.Message) error // 发送消息
}

// InterceptableChannel 可拦截的消息通道，可以添加拦截器对消息进行处理
type InterceptableChannel interface {
	MessageChannel
	AddInterceptor(interceptor Interceptor) // 添加消息通道的拦截器
}

func SendWith(ctx context.Context, msg *message.Message, channel MessageChannel) error {
	return channel.Send(ctx, msg)
}
