package main

import (
	"order-api/config"
	"order-api/dbl"
	"order-api/internal/product"
)

func main() {
	cfg := config.LoadConfig()
	db := dbl.NewDB(cfg)
	err := db.AutoMigrate(&product.Product{})
	if err != nil {
		panic(err)
	}
}
