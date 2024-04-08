//go:build wireinject
// +build wireinject

package template

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(
	ProvideRepository,
	ProvideService,
	ProvideController,
)

func NewController(db *gorm.DB) *Controller {
	panic(wire.Build(ProviderSet))
}
