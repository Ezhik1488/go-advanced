package main

import (
	"net/http"
	"order-api/config"
	"order-api/dbl"
	"order-api/internal/product"
)

func main() {
	cfg := config.LoadConfig()
	db := dbl.NewDB(cfg)

	router := http.NewServeMux()

	productRepo := product.NewProductRepository(db)

	product.NewProductHandler(router, &product.ProductHandlerDeps{
		ProductRepo: productRepo,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
