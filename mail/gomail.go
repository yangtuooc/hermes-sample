package mail

import (
	"context"
	"gopkg.in/gomail.v2"
	"hermes/channel"
	"hermes/channel/message"
	"io"
)

type GomailConfig struct {
	Host     string
	Port     int
	From     string
	Password string
}

var _ channel.Vendor = (*gomailVendor)(nil)

type gomailVendor struct {
	listeners channel.Listeners
	client    *gomail.Dialer
	config    *GomailConfig
}

func (g *gomailVendor) AddListener(listener channel.VendorListener) {
	g.listeners = append(g.listeners, listener)
}

func (g *gomailVendor) Id() string {
	return "gomail"
}

func (g *gomailVendor) Name() string {
	return "gomail"
}

func (g *gomailVendor) Type() string {
	return "mail"
}

func (g *gomailVendor) Description() string {
	return "gomail vendor"
}

func (g *gomailVendor) Send(ctx context.Context, message *message.Message) error {
	adapter := NewGomailAdapter(message)

	mail := gomail.NewMessage()
	mail.SetHeader("From", g.config.From)
	mail.SetHeader("To", adapter.To()...)
	mail.SetHeader("Subject", adapter.Subject()...)
	mail.SetBody(adapter.ContentType(), adapter.Body())
	for filename, attach := range adapter.Attachment() {
		mail.Attach(filename, gomail.SetCopyFunc(func(writer io.Writer) error {
			_, err := io.Copy(writer, attach)
			return err
		}))
	}

	g.listeners.OnRequest(ctx, message, g, mail)
	if err := g.client.DialAndSend(mail); err != nil {
		g.listeners.OnResponse(ctx, message, g, nil, err)
		return err
	}
	g.listeners.OnResponse(ctx, message, g, nil, nil)
	return nil
}

type GomailAdapter interface {
	To() []string
	Subject() []string
	Body() string
	ContentType() string
	Attachment() map[string]io.Reader
}

var _ GomailAdapter = (*adapter)(nil)

type adapter struct {
	message *message.Message
}

func (a *adapter) Attachment() map[string]io.Reader {
	attachment := a.message.GetHeader("Attachment")
	if a, ok := attachment.(map[string]io.Reader); ok {
		return a
	}
	return map[string]io.Reader{}
}

func (a *adapter) ContentType() string {
	contentType := a.message.GetHeader("ContentType")
	if c, ok := contentType.(string); ok {
		return c
	}
	return ""
}

func (a *adapter) To() []string {
	return a.message.To()
}

func (a *adapter) Subject() []string {
	subject := a.message.GetHeader("Subject")
	if s, ok := subject.([]string); ok {
		return s
	}
	return []string{}
}

func (a *adapter) Body() string {
	body := a.message.GetHeader("Body")
	if b, ok := body.(string); ok {
		return b
	}
	return ""
}

func NewGomailAdapter(message *message.Message) GomailAdapter {
	return &adapter{
		message: message,
	}
}

func NewGomailVendor(config *GomailConfig) channel.Vendor {
	return &gomailVendor{
		config: config,
		client: gomail.NewDialer(config.Host, config.Port, config.From, config.Password),
	}
}
