package services

import (
	"gorm.io/gorm"
)

var DbInstance = DbService{}

type DbService struct {
	Db *gorm.DB
}
