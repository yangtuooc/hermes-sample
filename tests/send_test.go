package tests

import (
	"context"
	"hermes/channel"
	"hermes/channel/message"
	"hermes/console"
	"hermes/mail"
	"hermes/sms"
	"testing"

	"github.com/google/uuid"
)

// case 1: 使用通道发送一条消息，并将消息内容输出到控制台
func TestSendCase1(t *testing.T) {
	ctx := context.Background()
	msg := message.New("hello world")
	err := channel.SendWith(ctx, msg, console.New())
	if err != nil {
		t.Error(err)
	}
}

// case 2: 使用短信通道发送一条消息，并将消息内容输出到控制台
func TestSendCase2(t *testing.T) {
	ctx := context.Background()
	msg := message.New("hello world")
	test := sms.Test()
	err := channel.SendWith(ctx, msg, test)
	if err != nil {
		t.Error(err)
	}
}

// case 3: 使用邮件通道发送一条消息，并将消息内容输出到控制台
func TestSendCase3(t *testing.T) {
	ctx := context.Background()
	msg := message.New("hello world")
	err := channel.SendWith(ctx, msg, mail.Test())
	if err != nil {
		t.Error(err)
	}
}

// case 4: 拦截用户重复发送的消息
func TestSendCase4(t *testing.T) {

	ctx := context.Background()
	msg := message.New("hello world")
	msg.SetRequestId(uuid.NewString())

	sc := sms.NewChannel()
	err := sc.Register(sms.Test())
	if err != nil {
		t.Error(err)
	}
	sc.SetSelector(channel.NewRoundRobinSelector())
	sc.AddInterceptor(newDeduplicationInterceptor())

	err = channel.SendWith(ctx, msg, sc)
	if err != nil {
		t.Error(err)
	}

	err = channel.SendWith(ctx, msg, sc)
	if err == nil {
		t.Error("expected error")
	}
}

// case 5: 使用通道供应商优先选择A供应商，A供应商发送失败，选择B供应商
func TestSendCase5(t *testing.T) {
	ctx := context.Background()
	msg := message.New("hello world")
	sc := sms.NewChannel()
	a := testA()
	b := testB()
	a.AddListener(newListener(func(vendor channel.Vendor, err error) {
		if err == nil {
			t.Error("expected error")
		}
	}))
	b.AddListener(newListener(func(vendor channel.Vendor, err error) {
		if err != nil {
			t.Error(err)
		}
	}))

	err := sc.Register(a)
	if err != nil {
		t.Error(err)
	}
	err = sc.Register(b)
	if err != nil {
		t.Error(err)
	}
	sc.SetSelector(newPrioritySelector([]string{a.Id(), b.Id()}))

	err = channel.SendWith(ctx, msg, sc)
	if err != nil {
		t.Error(err)
	}
}

// case 6: 通道发送失败，重试3次
func TestSendCase6(t *testing.T) {
	ctx := context.Background()
	msg := message.New("hello world")
	sc := sms.NewChannel()

	vendor := newRetry()
	vendor.AddListener(newListener(func(vendor channel.Vendor, err error) {
		if tried < 3 {
			if err == nil {
				t.Error("expected error")
			}
		} else {
			if err != nil {
				t.Error(err)
			}
		}
	}))
	err := sc.Register(vendor)
	if err != nil {
		t.Error(err)
	}
	sc.SetSelector(NewRetrySelector(3))

	err = channel.SendWith(ctx, msg, sc)
	if err != nil {
		t.Error(err)
	}
}

// case 7: 通过短信通道发送一条消息，指定发送渠道C
func TestSendCase7(t *testing.T) {
	ctx := context.Background()
	msg := message.New("hello world")
	sc := sms.NewChannel()

	a := testA()
	b := testB()
	c := testC()

	err := sc.Register(a)
	if err != nil {
		t.Error(err)
	}
	err = sc.Register(b)
	if err != nil {
		t.Error(err)
	}
	err = sc.Register(c)
	if err != nil {
		t.Error(err)
	}

	sc.SetSelector(newSpecificSelector())

	// not specified channel yet
	err = channel.SendWith(ctx, msg, sc)
	if err == nil {
		t.Error("expected error")
	}

	// specify channel c
	msg.SetHeader("vendorId", c.Id())
	err = channel.SendWith(ctx, msg, sc)
	if err != nil {
		t.Error(err)
	}
}
