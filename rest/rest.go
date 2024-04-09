package rest

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hermes/rest/template"
)

func Register(engine *gin.Engine, db *gorm.DB) {

	templateController := template.NewController(db)
	engine.POST("/template/generic", templateController.CreateGenericTemplate)
	engine.POST("/template/client", templateController.CreateClientTemplate)
}
