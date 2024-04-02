package sms

import (
	"context"
	"errors"
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
	request, err := adapter.Adapt()
	if err != nil {
		return err
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
	Adapt() (*client.SendSmsRequest, error)
}

var _ AliyunMessageAdapter = (*adapter)(nil)

type adapter struct {
	message *message.Message
}

func (a *adapter) Adapt() (*client.SendSmsRequest, error) {
	if a.message == nil {
		return nil, errors.New("aliyun message adapter: message is nil")
	}

	header := a.message.GetHeader(SignNameKey)
	signName, ok := header.(string)
	if !ok {
		return nil, fmt.Errorf("aliyun message adapter: signName is required to be string, got %T", header)
	}

	header = a.message.GetHeader(TemplateCodeKey)
	templateCode, ok := header.(string)
	if !ok {
		return nil, fmt.Errorf("aliyun message adapter: templateCode is required to be string, got %T", header)
	}

	to := a.message.GetTo()
	if len(to) == 0 {
		return nil, errors.New("aliyun message adapter: to is required")
	}
	joinedTo := strings.Join(to, ",")

	payload, ok := a.message.Payload.(string)
	if !ok {
		return nil, fmt.Errorf("aliyun message adapter: payload is required to be string, got %T", a.message.Payload)
	}

	return &client.SendSmsRequest{
		SignName:      &signName,
		TemplateCode:  &templateCode,
		TemplateParam: &payload,
		PhoneNumbers:  &joinedTo,
	}, nil
}

func NewAliyunMessageAdapter(message *message.Message) AliyunMessageAdapter {
	return &adapter{message: message}
}
