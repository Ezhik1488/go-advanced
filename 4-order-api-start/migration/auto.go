package main

import (
	"order-api/config"
	"order-api/dbl"
	"order-api/internal/core/models"
)

func main() {
	cfg := config.LoadConfig()
	db := dbl.NewDB(cfg)
	err := db.AutoMigrate(&models.Product{}, &models.User{}, &models.Order{})
	if err != nil {
		panic(err)
	}
}
