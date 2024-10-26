package commands

import (
	"os"
	"todo/models"
	"todo/services"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

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

	services.DbInstance.Db = db

	// Migrate the schema
	services.DbInstance.Db.AutoMigrate(
		&models.Task{},
		&models.Project{},
	)
}
