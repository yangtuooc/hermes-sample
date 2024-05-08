package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"hermes/app"
)

func main() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"))
	if err != nil {
		panic("failed to connect database")
	}
	engine := gin.Default()
	app.Register(engine, db)

	if err := engine.Run(":8080"); err != nil {
		panic("failed to start server")
	}
}
