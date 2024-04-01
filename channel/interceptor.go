package channel

import (
	"context"
	"hermes/channel/message"
)

type Interceptor interface {
	Intercept(ctx context.Context, message *message.Message, vendor MessageChannel) error
}

type InterceptorChain []Interceptor

func (c InterceptorChain) Intercept(ctx context.Context, message *message.Message, vendor MessageChannel) error {
	for _, interceptor := range c {
		if err := interceptor.Intercept(ctx, message, vendor); err != nil {
			return err
		}
	}
	return nil
}
