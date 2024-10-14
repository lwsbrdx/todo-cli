package commands

import (
	"todo/models"
	"todo/services"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DbService = services.DbService{}

func init() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{}) // TODO: make it configurable + chemin de la db dans les dossiers temporaire

	if err != nil {
		panic("failed to connect database")
	}

	DbService.Db = db

	// Migrate the schema
	DbService.Db.AutoMigrate(&models.Task{})

	DbService.Db.Create(&models.Task{
		Name:   "Hello, World!",
		Status: models.Todo,
	})
}
