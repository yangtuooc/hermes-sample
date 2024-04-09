package channel

import (
	"context"
	"hermes/channel/message"
)

type Vendors []Vendor

// Vendor 消息通道的提供者，实际发送消息的处理者，由各个供应商实现，如阿里云、腾讯云等。
type Vendor interface {
	AddListener(listener VendorListener) // 添加消息通道的监听器
	NamedChannel
}

// StrategyChannel 策略消息通道，用于选择消息通道的实际发送者
type StrategyChannel interface {
	MessageChannel
	SetSelector(selector VendorSelector) // 设置消息通道的选择策略，用于选择消息通道的实际发送者
}

type NamedChannels []NamedChannel

// NamedChannel 带命名的消息通道，可以标识消息通道的唯一性
type NamedChannel interface {
	Id() string          // 消息通道的唯一标识
	Name() string        // 消息通道的名称
	Type() string        // 消息通道的类型
	Description() string // 消息通道的描述
	MessageChannel
}

// AbstractChannel 抽象消息通道，包含了消息通道的基本功能
type AbstractChannel interface {
	InterceptableChannel
	StrategyChannel
	VendorRegistry
	NamedChannel
}

func (vs Vendors) MapById() map[string]Vendor {
	m := make(map[string]Vendor)
	for _, v := range vs {
		m[v.Id()] = v
	}
	return m
}

func (vs Vendors) First() Vendor {
	if len(vs) > 0 {
		return vs[0]
	}
	return nil
}

func (vs Vendors) Find(id string) Vendor {
	for _, v := range vs {
		if v.Id() == id {
			return v
		}
	}
	return nil
}

// VendorRegistry 消息通道的注册中心，用于注册消息通道
type VendorRegistry interface {
	Register(vendor Vendor) error // 注册消息通道
}

// VendorListener 消息通道的监听器，用于监听消息通道的请求和响应
type VendorListener interface {
	OnRequest(ctx context.Context, message *message.Message, vendor Vendor, request any)              // 消息通道的请求监听器
	OnResponse(ctx context.Context, message *message.Message, vendor Vendor, response any, err error) // 消息通道的响应监听器
}

type Listeners []VendorListener

func (ls Listeners) OnRequest(ctx context.Context, message *message.Message, vendor Vendor, request any) {
	for _, listener := range ls {
		listener.OnRequest(ctx, message, vendor, request)
	}
}

func (ls Listeners) OnResponse(ctx context.Context, message *message.Message, vendor Vendor, response any, err error) {
	for _, listener := range ls {
		listener.OnResponse(ctx, message, vendor, response, err)
	}
}
