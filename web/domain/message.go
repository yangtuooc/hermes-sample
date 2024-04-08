package domain

import (
	"context"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	TemplateId string            `gorm:"column:template_id"`
	To         []string          `gorm:"column:to;serializer:json"`
	Args       map[string]string `gorm:"column:args;serializer:json"`
	Timing     bool              `gorm:"column:timing"`
	Cron       string            `gorm:"column:cron"`
	Sent       bool              `gorm:"column:sent"`
}

func (Message) TableName() string {
	return "messages"
}

type SimpleMessage struct {
	TemplateId string            `json:"templateId"`
	To         []string          `json:"to"`
	Args       map[string]string `json:"args"`
}

type TimingMessage struct {
	SimpleMessage
	Cron string `json:"cron"`
}

type MessageRepository interface {
	Save(ctx context.Context, message *Message) error
}

type MessageService interface {
	SendSimpleMessage(ctx context.Context, message *SimpleMessage) error
	SendTimingMessage(ctx context.Context, message *TimingMessage) error
}
