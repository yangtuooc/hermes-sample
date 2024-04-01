package channel

import (
	"context"
	"hermes/message"
)

var _ SwitchableChannel = (*delegateMessageChannel)(nil)

type delegateMessageChannel struct {
	vendors   Vendors
	chain     InterceptorChain
	selector  Selector
	listeners Listeners
}

func (a *delegateMessageChannel) Register(vendor Vendor) {
	a.vendors = append(a.vendors, vendor)
}

func (a *delegateMessageChannel) AddSelector(selector Selector) {
	a.selector = selector
}

func (a *delegateMessageChannel) Id() string {
	return "abstract-channel"
}

func (a *delegateMessageChannel) Name() string {
	return "abstract-channel"
}

func (a *delegateMessageChannel) Type() string {
	return "abstract-channel"
}

func (a *delegateMessageChannel) AddListener(listener Listener) {
	a.listeners = append(a.listeners, listener)
}

func (a *delegateMessageChannel) Send(ctx context.Context, message *message.Message) error {
	return a.selector.Select(ctx, message, a.vendors, func(vendor Vendor) error {
		if err := a.chain.Intercept(ctx, message, vendor); err != nil {
			return err
		}
		return vendor.Send(ctx, message)
	})
}

func (a *delegateMessageChannel) AddInterceptor(interceptor Interceptor) {
	a.chain = append(a.chain, interceptor)
}

func NewDelegateChannel() SwitchableChannel {
	return &delegateMessageChannel{
		vendors:   make(Vendors, 0),
		chain:     make(InterceptorChain, 0),
		listeners: make(Listeners, 0),
	}
}
