package db

import (
	"gorm.io/gorm"
)

type GormClient struct {
	DB *gorm.DB
}
