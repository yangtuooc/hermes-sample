package template

import (
	"github.com/gin-gonic/gin"
	"hermes/web/domain"
	"hermes/web/rest"
)

type Controller struct {
	svc domain.TemplateService
}

func (c *Controller) CreateGenericTemplate(ctx *gin.Context) {
	var req domain.GenericTemplate
	if err := ctx.ShouldBind(&req); err != nil {
		rest.ResponseBadRequest(ctx, err.Error())
		return
	}
	err := c.svc.CreateGenericTemplate(ctx, &req)
	if err != nil {
		rest.ResponseInternalServerError(ctx, err.Error())
		return
	}
	rest.ResponseCreated(ctx, nil)
}

func (c *Controller) CreateClientTemplate(ctx *gin.Context) {
	var req domain.ClientTemplate
	if err := ctx.ShouldBind(&req); err != nil {
		rest.ResponseBadRequest(ctx, err.Error())
		return
	}
	err := c.svc.CreateClientTemplate(ctx, &req)
	if err != nil {
		rest.ResponseInternalServerError(ctx, err.Error())
		return
	}
	rest.ResponseCreated(ctx, nil)
}
