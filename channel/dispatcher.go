package channel

import (
	"context"
	"hermes/channel/message"
)

var _ MessageChannel = (*dispatcher)(nil)
var _ InterceptableChannel = (*dispatcher)(nil)

type dispatcher struct {
	factory MessageChannelFactory
	chain   InterceptorChain // global interceptor chain, all messages will go through this chain
}

func (d *dispatcher) AddInterceptor(interceptor Interceptor) {
	d.chain = append(d.chain, interceptor)
}

func (d *dispatcher) Send(ctx context.Context, message *message.Message) error {
	channel := d.factory.GetChannel(message.GetChannelId())
	if channel == nil {
		return ErrChannelNotFound(message.GetChannelId())
	}
	if err := d.chain.Intercept(ctx, message, channel); err != nil {
		return err
	}
	return channel.Send(ctx, message)
}

func NewDispatcher(factory MessageChannelFactory) MessageChannel {
	return &dispatcher{factory: factory}
}
