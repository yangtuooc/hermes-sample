package domain

import (
	"context"
	"gorm.io/gorm"
)

type AliyunSmsTemplate struct {
	gorm.Model
	Code     string `gorm:"column:code"`
	SignName string `gorm:"column:sign_name"`
}

func (a *AliyunSmsTemplate) TableName() string {
	return "aliyun_sms_templates"
}

type AliyunSmsTemplateRepository interface {
	GetTemplateByCode(ctx context.Context, code string) (*AliyunSmsTemplate, error)
}
