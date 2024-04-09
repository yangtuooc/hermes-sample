package main

import (
	"gorm.io/gen"
	"hermes/app/domain"
)

func main() {
	g := gen.NewGenerator(
		gen.Config{
			OutPath:      "api/template/gen/query",
			WithUnitTest: true,
			Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		},
	)
	g.ApplyBasic(domain.Template{})
	g.Execute()
}
