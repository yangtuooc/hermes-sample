package channel

import (
	"context"
	"hermes/channel/message"
	"sync"
)

var (
	dispatcherInstance MessageChannel
	dispatcherOnce     sync.Once
)

var _ MessageChannel = (*dispatcher)(nil)
var _ InterceptableChannel = (*dispatcher)(nil)

// 这是一个消息通道的统一调度器，由于消息通道的实现可能会有不同的实现，所以需要一个调度器来统一调度
type dispatcher struct {
	factory MessageChannelFactory
	chain   InterceptorChain // global interceptor chain, all messages will go through this chain
}

func (d *dispatcher) AddInterceptor(interceptor Interceptor) {
	d.chain = append(d.chain, interceptor)
}

func (d *dispatcher) Send(ctx context.Context, message *message.Message) error {
	channel := d.factory.GetChannel(message.ChannelId())
	if channel == nil {
		return ErrChannelNotFound(message.ChannelId())
	}
	if err := d.chain.Intercept(ctx, message, channel); err != nil {
		return err
	}
	return channel.Send(ctx, message)
}

func NewDispatcher(factory MessageChannelFactory) MessageChannel {
	dispatcherOnce.Do(func() {
		dispatcherInstance = &dispatcher{
			factory: factory,
			chain:   make(InterceptorChain, 0),
		}
	})
	return dispatcherInstance
}
