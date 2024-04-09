package interceptor

import (
	"context"
	"encoding/json"
	"fmt"
	"hermes/channel"
	"hermes/channel/message"
	"hermes/rest/domain"
	"hermes/sms"
)

var _ channel.Interceptor = (*aliyunInterceptor)(nil)

type aliyunInterceptor struct {
	vendorId string
	repo     domain.AliyunSmsTemplateRepository
}

func (a *aliyunInterceptor) Intercept(ctx context.Context, message *message.Message, vendor channel.NamedChannel) error {
	if a.vendorId != vendor.Id() {
		return nil
	}
	payload := message.GetHeader("args")
	if payload == nil {
		return fmt.Errorf("aliyun interceptor: payload is nil")
	}
	payloadJson, _ := json.Marshal(payload)
	message.Payload = payloadJson

	extra := message.GetHeader("extra")
	if extra == nil {
		return fmt.Errorf("aliyun interceptor: the send sms template is not specified")
	}
	m := extra.(map[string]string)
	aliyunSmsTemplateCode := m["templateCode"]
	template, err := a.repo.GetTemplateByCode(ctx, aliyunSmsTemplateCode)
	if err != nil {
		return fmt.Errorf("aliyun interceptor: get template by code failed, code: %s, err: %w", aliyunSmsTemplateCode, err)
	}
	message.SetHeader(sms.TemplateCodeKey, template.Code)
	message.SetHeader(sms.SignNameKey, template.SignName)
	return nil
}
