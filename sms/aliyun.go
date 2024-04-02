package sms

import (
	"context"
	"hermes/channel"
	"hermes/channel/message"
)

type AliyunConfig struct {
}

var _ channel.Vendor = (*aliyunVendor)(nil)

type aliyunVendor struct {
	config    *AliyunConfig
	listeners channel.Listeners
}

func (a *aliyunVendor) Id() string {
	return "aliyun"
}

func (a *aliyunVendor) Name() string {
	return "阿里云"
}

func (a *aliyunVendor) Type() string {
	return "sms"
}

func (a *aliyunVendor) Description() string {
	return "阿里云短信通道"
}

func (a *aliyunVendor) AddListener(listener channel.VendorListener) {
	a.listeners = append(a.listeners, listener)
}

func (a *aliyunVendor) Send(ctx context.Context, message *message.Message) error {
	return nil
}

func NewAliyunVendor(config *AliyunConfig) channel.Vendor {
	return &aliyunVendor{config: config}
}
