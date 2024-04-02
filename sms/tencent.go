package sms

import (
	"context"
	"encoding/json"
	"fmt"
	"hermes/channel"
	"hermes/channel/message"
)

type TencentConfig struct {
}

var _ channel.Vendor = (*tencentVendor)(nil)

type tencentVendor struct {
	config    *TencentConfig
	listeners channel.Listeners
}

func (t *tencentVendor) Id() string {
	return "tencent"
}

func (t *tencentVendor) Type() string {
	return "sms"
}

func (t *tencentVendor) Name() string {
	return "腾讯云"
}

func (t *tencentVendor) Description() string {
	return "腾讯云短信"
}

func (t *tencentVendor) AddListener(listener channel.VendorListener) {
	t.listeners = append(t.listeners, listener)
}

func (t *tencentVendor) Send(ctx context.Context, message *message.Message) error {
	t.listeners.OnRequest(ctx, message, t, nil)
	marshal, _ := json.Marshal(message)
	fmt.Println(string(marshal))
	t.listeners.OnResponse(ctx, message, t, nil, nil)
	return nil
}

func NewTencentVendor(config *TencentConfig) channel.Vendor {
	return &tencentVendor{config: config}
}
