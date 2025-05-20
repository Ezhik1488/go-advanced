package main

import (
	"order-api/config"
	"order-api/dbl"
	"order-api/internal/order"
	"order-api/internal/product"
	"order-api/internal/user"
)

func main() {
	cfg := config.LoadConfig()
	db := dbl.NewDB(cfg)
	err := db.AutoMigrate(&product.Product{}, &user.User{}, &order.Order{})
	if err != nil {
		panic(err)
	}
}
