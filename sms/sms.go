package sms

import (
	"context"
	"hermes/channel"
	"hermes/channel/message"
)

var _ channel.AbstractChannel = (*smsChannel)(nil)

type smsChannel struct {
	delegate channel.AbstractChannel
}

func (s *smsChannel) Register(vendor channel.Vendor) {
	s.delegate.Register(vendor)
}

func (s *smsChannel) SetSelector(selector channel.VendorSelector) {
	s.delegate.SetSelector(selector)
}

func (s *smsChannel) AddInterceptor(interceptor channel.Interceptor) {
	s.delegate.AddInterceptor(interceptor)
}

func (s *smsChannel) Id() string {
	return "sms-channel"
}

func (s *smsChannel) Type() string {
	return "sms-channel"
}

func (s *smsChannel) Name() string {
	return "sms-channel"
}

func (s *smsChannel) Send(ctx context.Context, message *message.Message) error {
	return s.delegate.Send(ctx, message)
}

func NewChannel() channel.AbstractChannel {
	return &smsChannel{
		delegate: channel.NewDelegateChannel(),
	}
}