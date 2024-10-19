package commands

import (
	"os"
	"todo/models"
	"todo/services"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DbService = services.DbService{}

func init() {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}
	// TODO : faire en sorte que le chemin de la db soit configurable
	dbName := cacheDir + "/todo.db"
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic("failed to connect database")
	}

	DbService.Db = db

	// Migrate the schema
	DbService.Db.AutoMigrate(&models.Task{})
}
