package channel

import (
	"context"
	"hermes/message"
)

type Listener interface {
	OnRequest(ctx context.Context, message *message.Message, request any)
	OnResponse(ctx context.Context, message *message.Message, response any, err error)
}

type Listeners []Listener

func (ls Listeners) OnRequest(ctx context.Context, message *message.Message, request any) {
	for _, listener := range ls {
		listener.OnRequest(ctx, message, request)
	}
}

func (ls Listeners) OnResponse(ctx context.Context, message *message.Message, response any, err error) {
	for _, listener := range ls {
		listener.OnResponse(ctx, message, response, err)
	}
}
