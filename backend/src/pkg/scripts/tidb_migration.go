package main

import (
	"app/src/pkg/config"
	"app/src/pkg/domain/entity"
	"app/src/pkg/infra/db"
)

func main() {
	cfg, err := config.NewConfigVariables()
	if err != nil {
		panic(err)
	}
	dbClient, err := db.NewGormClient(*cfg)
	if err != nil {
		panic(err)
	}

	// users
	err = dbClient.DB.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err)
	}

}
