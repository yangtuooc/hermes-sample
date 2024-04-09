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
	RequestId  string            `json:"requestId" binding:"required"`
	TemplateId string            `json:"templateId" binding:"required"`
	To         []string          `json:"to" binding:"required"`
	Args       map[string]string `json:"args" binding:"required"`
	Extra      map[string]string `json:"extra"`
}

type TimingMessage struct {
	SimpleMessage
	Cron string `json:"cron" binding:"required"`
}

type MessageRepository interface {
	Save(ctx context.Context, message *Message) error
}

type MessageService interface {
	SendSimpleMessage(ctx context.Context, message *SimpleMessage) error
	SendTimingMessage(ctx context.Context, message *TimingMessage) error
}
