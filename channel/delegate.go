package channel

import (
	"context"
	"hermes/channel/message"
)

var _ AbstractChannel = (*delegateMessageChannel)(nil)

type delegateMessageChannel struct {
	vendors  Vendors
	chain    InterceptorChain
	selector VendorSelector
}

func (a *delegateMessageChannel) Register(vendor Vendor) {
	a.vendors = append(a.vendors, vendor)
}

func (a *delegateMessageChannel) SetSelector(selector VendorSelector) {
	a.selector = selector
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

func NewDelegateChannel() AbstractChannel {
	return &delegateMessageChannel{
		vendors: make(Vendors, 0),
		chain:   make(InterceptorChain, 0),
	}
}
