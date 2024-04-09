package template

import (
	"gorm.io/gorm"
	"hermes/rest/domain"
	"hermes/rest/template/gen/query"
	"sync"
)

var (
	repo     domain.TemplateRepository
	repoOnce sync.Once

	svc     domain.TemplateService
	svcOnce sync.Once

	hdl     *Controller
	hdlOnce sync.Once
)

func ProvideRepository(db *gorm.DB) domain.TemplateRepository {
	repoOnce.Do(func() {
		repo = &repository{
			db:   db,
			repo: query.Use(db),
		}
	})
	return repo
}

func ProvideService(repo domain.TemplateRepository) domain.TemplateService {
	svcOnce.Do(func() {
		svc = &service{
			repo: repo,
		}
	})
	return svc
}

func ProvideController(svc domain.TemplateService) *Controller {
	hdlOnce.Do(func() {
		hdl = &Controller{
			svc: svc,
		}
	})
	return hdl
}
