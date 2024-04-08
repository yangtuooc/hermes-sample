package message

import (
	"context"
	"hermes/channel"
	"hermes/web/domain"
)

var _ domain.MessageService = (*service)(nil)

type service struct {
	dispatcher      channel.MessageChannel
	templateService domain.TemplateService
}

func (s *service) SendSimpleMessage(ctx context.Context, message *domain.SimpleMessage) error {
	tpl, err := s.templateService.GetTemplate(ctx, message.TemplateId)
	if err != nil {
		return err
	}
	msg, err := buildSimpleMessage(tpl, message)
	if err != nil {
		return err
	}
	return s.dispatcher.Send(ctx, msg)
}

func (s *service) SendTimingMessage(ctx context.Context, message *domain.TimingMessage) error {
	//TODO implement me
	panic("implement me")
}
