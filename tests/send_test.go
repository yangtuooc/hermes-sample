package tests

import (
	"context"
	"github.com/google/uuid"
	"hermes/channel"
	"hermes/console"
	"hermes/mail"
	"hermes/message"
	"hermes/sms"
	"testing"
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
	sc.Register(sms.Test())
	sc.AddSelector(channel.NewRoundRobinSelector())
	sc.AddInterceptor(newDeduplicationInterceptor())

	err := channel.SendWith(ctx, msg, sc)
	if err != nil {
		t.Error(err)
	}

	err = channel.SendWith(ctx, msg, sc)
	if err == nil {
		t.Error("expected error")
	}
}

// case 5: 短信通道供应商阿里云不可用时，使用腾讯云发送短信
func TestSendCase5(t *testing.T) {
	//var aliyunVendor = func(ctx context.Context, msg *message.Message) error {
	//	return errors.New("aliyun vendor not available")
	//}
	//
	//var tencentVendor = func(ctx context.Context, msg *message.Message) error {
	//	jsonBytes, _ := json.Marshal(msg)
	//	fmt.Println(string(jsonBytes))
	//	return nil
	//}

}
