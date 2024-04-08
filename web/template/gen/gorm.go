package main

import (
	"gorm.io/gen"
	"hermes/web/domain"
)

func main() {
	g := gen.NewGenerator(
		gen.Config{
			OutPath:      "web/template/gen/query",
			WithUnitTest: true,
			Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		},
	)
	g.ApplyBasic(domain.Template{})
	g.Execute()
}
