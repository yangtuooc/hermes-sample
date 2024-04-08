package sms

import (
	"context"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"hermes/channel"
	"hermes/channel/message"
	"strings"
)

const (
	SignNameKey     = "signName"
	TemplateCodeKey = "templateCode"
)

type AliyunConfig struct {
	AccessKeyId     string
	AccessKeySecret string
}

var _ channel.Vendor = (*aliyunVendor)(nil)

type aliyunVendor struct {
	config    *AliyunConfig
	listeners channel.Listeners
	client    *client.Client
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

	adapter := NewAliyunMessageAdapter(message)
	request := &client.SendSmsRequest{
		SignName:      adapter.SignName(),
		TemplateCode:  adapter.TemplateCode(),
		TemplateParam: adapter.TemplateParam(),
		PhoneNumbers:  adapter.PhoneNumbers(),
	}

	a.listeners.OnRequest(ctx, message, a, request)
	response, err := a.client.SendSms(request)
	a.listeners.OnResponse(ctx, message, a, response, err)

	if err != nil {
		return fmt.Errorf("aliyun vendor: send failed, %w", err)
	}

	if *response.Body.Code == "OK" {
		return nil
	}
	return fmt.Errorf("aliyun vendor: send failed, code: %s, message: %s", *response.Body.Code, *response.Body.Message)
}

func NewAliyunVendor(config *AliyunConfig) (channel.Vendor, error) {
	c, err := client.NewClient(&openapi.Config{
		AccessKeyId:     &config.AccessKeyId,
		AccessKeySecret: &config.AccessKeySecret,
	})
	if err != nil {
		return nil, err
	}

	return &aliyunVendor{
		config: config,
		client: c,
	}, nil
}

type AliyunMessageAdapter interface {
	SignName() *string
	TemplateCode() *string
	TemplateParam() *string
	PhoneNumbers() *string
}

var _ AliyunMessageAdapter = (*adapter)(nil)

type adapter struct {
	message *message.Message
}

func (a *adapter) SignName() *string {
	signName := a.message.GetHeader(SignNameKey)
	if s, ok := signName.(string); ok {
		return &s
	}
	return nil
}

func (a *adapter) TemplateCode() *string {
	templateCode := a.message.GetHeader(TemplateCodeKey)
	if s, ok := templateCode.(string); ok {
		return &s
	}
	return nil
}

func (a *adapter) TemplateParam() *string {
	if p, ok := a.message.Payload.(string); ok {
		return &p
	}
	return nil
}

func (a *adapter) PhoneNumbers() *string {
	to := a.message.To()
	if len(to) == 0 {
		return nil
	}
	joined := strings.Join(to, ",")
	return &joined
}

func NewAliyunMessageAdapter(message *message.Message) AliyunMessageAdapter {
	return &adapter{message: message}
}
