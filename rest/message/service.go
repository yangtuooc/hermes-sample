package message

import (
	"context"
	"hermes/channel"
	"hermes/rest/domain"
)

var _ domain.MessageService = (*service)(nil)

type service struct {
	dispatcher      channel.MessageChannel
	templateService domain.TemplateService
}

func (s *service) SendSimpleMessage(ctx context.Context, simpleMessage *domain.SimpleMessage) error {
	tpl, err := s.templateService.GetTemplate(ctx, simpleMessage.TemplateId)
	if err != nil {
		return err
	}
	message := buildSimpleMessage(tpl, simpleMessage)
	return s.dispatcher.Send(ctx, message)
}

func (s *service) SendTimingMessage(ctx context.Context, message *domain.TimingMessage) error {
	//TODO implement me
	panic("implement me")
}
