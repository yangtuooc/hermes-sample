package domain

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"strings"
)

var (
	ErrTemplateDisabled = errors.New("template is disabled")
)

type Template struct {
	gorm.Model
	TemplateId  string `gorm:"column:template_id"`  // 模板ID
	Name        string `gorm:"column:name"`         // 模板名称
	Channel     string `gorm:"column:channel"`      // 消息发送渠道
	Content     string `gorm:"column:content"`      // 模板内容，可以使用${}占位符
	SendAccount string `gorm:"column:send_account"` // 发送账号，相同渠道下可能存在多个账户
	ClientId    string `gorm:"column:client_id"`    // 客户端id，指与系统对接的第三方系统，如果为空则表示通用模板
	Enabled     bool   `gorm:"column:enabled"`      // 是否启用
	Comment     string `gorm:"column:comment"`      // 备注
}

func (t *Template) TableName() string {
	return "templates"
}

func (t *Template) Render(args map[string]string) string {
	if t.Content == "" {
		return ""
	}
	var content = t.Content
	for k, v := range args {
		content = strings.ReplaceAll(t.Content, "${"+k+"}", v)
	}
	return content
}

func (t *Template) ShallowCopy() Template {
	return *t
}

type GenericTemplate struct {
	Name        string `json:"name"`
	Channel     string `json:"channel"`
	Content     string `json:"content"`
	SendAccount string `json:"sendAccount"`
}

func NewTemplateWithGeneric(generic *GenericTemplate) *Template {
	return &Template{
		Name:        generic.Name,
		Channel:     generic.Channel,
		Content:     generic.Content,
		SendAccount: generic.SendAccount,
		Enabled:     true,
	}
}

type ClientTemplate struct {
	ClientId string `json:"clientId"`
	GenericTemplate
}

func NewTemplateWithClient(client *ClientTemplate) *Template {
	return &Template{
		Name:        client.Name,
		Channel:     client.Channel,
		Content:     client.Content,
		SendAccount: client.SendAccount,
		ClientId:    client.ClientId,
		Enabled:     true,
	}
}

type TemplateRepository interface {
	Save(ctx context.Context, template *Template) error
	FindByTemplateId(ctx context.Context, templateId string) (*Template, error)
}

type TemplateService interface {
	CreateGenericTemplate(ctx context.Context, generic *GenericTemplate) error
	CreateClientTemplate(ctx context.Context, client *ClientTemplate) error
	GetTemplate(ctx context.Context, templateId string) (*Template, error)
}
