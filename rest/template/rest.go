package template

import (
	"github.com/gin-gonic/gin"
	"hermes/rest/api"
	"hermes/rest/domain"
)

type Controller struct {
	svc domain.TemplateService
}

func (c *Controller) CreateGenericTemplate(ctx *gin.Context) {
	var req domain.GenericTemplate
	if err := ctx.ShouldBind(&req); err != nil {
		api.ResponseBadRequest(ctx, err.Error())
		return
	}
	err := c.svc.CreateGenericTemplate(ctx, &req)
	if err != nil {
		api.ResponseInternalServerError(ctx, err.Error())
		return
	}
	api.ResponseCreated(ctx, nil)
}

func (c *Controller) CreateClientTemplate(ctx *gin.Context) {
	var req domain.ClientTemplate
	if err := ctx.ShouldBind(&req); err != nil {
		api.ResponseBadRequest(ctx, err.Error())
		return
	}
	err := c.svc.CreateClientTemplate(ctx, &req)
	if err != nil {
		api.ResponseInternalServerError(ctx, err.Error())
		return
	}
	api.ResponseCreated(ctx, nil)
}
