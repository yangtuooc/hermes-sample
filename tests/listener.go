package tests

import (
	"context"
	"hermes/channel"
	"hermes/channel/message"
)

var _ channel.VendorListener = (*listener)(nil)

type callback func(vendor channel.Vendor, err error)

type listener struct {
	callback callback
}

func (l *listener) OnRequest(ctx context.Context, message *message.Message, vendor channel.Vendor, request any) {

}

func (l *listener) OnResponse(ctx context.Context, message *message.Message, vendor channel.Vendor, response any, err error) {
	l.callback(vendor, err)
}

func newListener(callback callback) channel.VendorListener {
	return &listener{callback: callback}
}
