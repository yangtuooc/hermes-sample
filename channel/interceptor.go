package channel

import (
	"context"
	"hermes/message"
)

type Interceptor interface {
	Intercept(ctx context.Context, message *message.Message, vendor Vendor) error
}

type InterceptorChain []Interceptor

func (c InterceptorChain) Intercept(ctx context.Context, message *message.Message, vendor Vendor) error {
	for _, interceptor := range c {
		if err := interceptor.Intercept(ctx, message, vendor); err != nil {
			return err
		}
	}
	return nil
}
