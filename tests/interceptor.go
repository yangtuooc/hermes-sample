package tests

import (
	"context"
	"errors"
	"hermes/channel"
	"hermes/channel/message"
)

var _ channel.Interceptor = (*deduplicationInterceptor)(nil)

type deduplicationInterceptor struct {
	sent map[string]bool
}

func (i *deduplicationInterceptor) Intercept(ctx context.Context, message *message.Message, vendor channel.MessageChannel) error {
	if i.sent[message.GetRequestId()] {
		return errors.New("message already sent")
	}
	i.sent[message.GetRequestId()] = true
	return nil
}

func newDeduplicationInterceptor() channel.Interceptor {
	return &deduplicationInterceptor{
		sent: make(map[string]bool),
	}
}
