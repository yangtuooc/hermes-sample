package mail

import (
	"context"
	"hermes/channel"
	"hermes/channel/message"
)

var _ channel.AbstractChannel = (*mailChannel)(nil)

type mailChannel struct {
	delegate channel.AbstractChannel
}

func (m *mailChannel) Send(ctx context.Context, message *message.Message) error {
	return m.delegate.Send(ctx, message)
}

func (m *mailChannel) AddInterceptor(interceptor channel.Interceptor) {
	m.delegate.AddInterceptor(interceptor)
}

func (m *mailChannel) SetSelector(selector channel.VendorSelector) {
	m.delegate.SetSelector(selector)
}

func (m *mailChannel) Register(vendor channel.Vendor) error {
	return m.delegate.Register(vendor)
}

func (m *mailChannel) Id() string {
	return "mail"
}

func (m *mailChannel) Name() string {
	return "邮件通道"
}

func (m *mailChannel) Type() string {
	return "mail"
}

func (m *mailChannel) Description() string {
	return "邮件通道"
}
