package channel

import (
	"context"
	"hermes/channel/message"
)

type Vendors []Vendor

type Vendor interface {
	AddListener(listener VendorListener)
	NamedChannel
}

type StrategyChannel interface {
	MessageChannel
	SetSelector(selector VendorSelector)
}

type NamedChannel interface {
	Id() string
	Name() string
	Type() string
	Description() string
	MessageChannel
}

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

type VendorRegistry interface {
	Register(vendor Vendor) error
}

type VendorListener interface {
	OnRequest(ctx context.Context, message *message.Message, vendor Vendor, request any)
	OnResponse(ctx context.Context, message *message.Message, vendor Vendor, response any, err error)
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
