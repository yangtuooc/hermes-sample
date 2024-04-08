package template

import (
	"context"
	"hermes/web/domain"
)

var _ domain.TemplateService = (*service)(nil)

type service struct {
	repo domain.TemplateRepository
}

func (s *service) CreateGenericTemplate(ctx context.Context, generic *domain.GenericTemplate) error {
	template := domain.NewTemplateWithGeneric(generic)
	return s.repo.Save(ctx, template)
}

func (s *service) CreateClientTemplate(ctx context.Context, client *domain.ClientTemplate) error {
	template := domain.NewTemplateWithClient(client)
	return s.repo.Save(ctx, template)
}

func (s *service) GetTemplate(ctx context.Context, templateId string) (*domain.Template, error) {
	return s.repo.FindByTemplateId(ctx, templateId)
}
