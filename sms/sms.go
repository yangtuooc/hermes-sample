package sms

import (
	"context"
	"hermes/channel"
	"hermes/message"
)

var _ channel.SwitchableChannel = (*smsChannel)(nil)

type smsChannel struct {
	delegate channel.SwitchableChannel
}

func (s *smsChannel) Register(vendor channel.Vendor) {
	s.delegate.Register(vendor)
}

func (s *smsChannel) AddSelector(selector channel.Selector) {
	s.delegate.AddSelector(selector)
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

func (s *smsChannel) AddListener(listener channel.Listener) {
	s.delegate.AddListener(listener)
}

func (s *smsChannel) Send(ctx context.Context, message *message.Message) error {
	return s.delegate.Send(ctx, message)
}

func NewChannel() channel.SwitchableChannel {
	return &smsChannel{
		delegate: channel.NewDelegateChannel(),
	}
}
