package db

import (
	"app/src/pkg/config"
	"app/src/pkg/errorhandle"
	"app/src/pkg/lib"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormClient struct {
	DB *gorm.DB
}

func NewGormClient(cfg config.ConfigVariables) (*GormClient, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&tls=%s",
		cfg.TiDB.User, cfg.TiDB.Password, cfg.TiDB.Host, cfg.TiDB.Port, DBNamePhotoDeepArchive, "true")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		lib.LogError("infra.NewGormClient().gorm.Open", err)
		return nil, errorhandle.Wrap("infra.NewGormClient().gorm.Open", err)
	}

	return &GormClient{
		DB: db,
	}, nil
}
