package commands

import (
	"log"
	"os"
	"todo/models"
	"todo/services"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DbService = services.DbService{}

func init() {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}
	// TODO : faire en sorte que le chemin de la db soit configurable
	dbName := cacheDir + "/todo.db"

	log.Printf("Using database %s", dbName)

	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DbService.Db = db

	// Migrate the schema
	DbService.Db.AutoMigrate(&models.Task{})
}
