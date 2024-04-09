package channel

import (
	"context"
	"hermes/channel/message"
)

// Interceptor 拦截器，可以在消息实际发送前对消息进行处理或拦截，可以修改消息内容，也可以根据消息内容决定是否发送
type Interceptor interface {
	Intercept(ctx context.Context, message *message.Message, vendor NamedChannel) error
}

type InterceptorChain []Interceptor

func (c InterceptorChain) Intercept(ctx context.Context, message *message.Message, vendor NamedChannel) error {
	for _, interceptor := range c {
		if err := interceptor.Intercept(ctx, message, vendor); err != nil {
			return err
		}
	}
	return nil
}
