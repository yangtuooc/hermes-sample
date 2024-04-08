package tests

import (
	"context"
	"errors"
	"fmt"
	"hermes/channel"
	"hermes/channel/message"
)

var _ channel.Vendor = (*retry)(nil)

var tried = 0

type retry struct {
	listener channel.Listeners
}

func (r *retry) Id() string {
	return "retry"
}

func (r *retry) Type() string {
	return "sms"
}

func (r *retry) Name() string {
	return "retry"
}

func (r *retry) Description() string {
	return "retry"
}

func (r *retry) AddListener(listener channel.VendorListener) {
	r.listener = append(r.listener, listener)
}

func (r *retry) Send(ctx context.Context, message *message.Message) error {
	var err error
	if tried < 3 {
		info := fmt.Sprintf("send failed, retry the %d time\n", tried+1)
		fmt.Println(info)
		err = errors.New("retry")
		r.listener.OnResponse(ctx, message, r, nil, err)
	}
	tried++
	if tried == 3 {
		fmt.Println("send success")
		r.listener.OnResponse(ctx, message, r, nil, nil)
		return nil
	}
	return err
}

func newRetry() channel.Vendor {
	return &retry{
		listener: make(channel.Listeners, 0),
	}
}
