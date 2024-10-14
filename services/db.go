package services

import (
	"gorm.io/gorm"
)

type DbService struct {
	Db *gorm.DB
}
