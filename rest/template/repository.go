package template

import (
	"context"
	"gorm.io/gorm"
	"hermes/rest/domain"
	"hermes/rest/template/gen/query"
)

var _ domain.TemplateRepository = (*repository)(nil)

type repository struct {
	db   *gorm.DB
	repo *query.Query
}

func (r *repository) Save(ctx context.Context, template *domain.Template) error {
	return r.repo.Template.WithContext(ctx).Save(template)
}

func (r *repository) FindByTemplateId(ctx context.Context, templateId string) (*domain.Template, error) {
	tpl := r.repo.Template
	return tpl.WithContext(ctx).Where(tpl.TemplateId.Eq(templateId)).First()
}
