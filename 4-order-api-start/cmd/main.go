package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"order-api/config"
	"order-api/dbl"
	"order-api/internal/product"
	"order-api/pkg/middleware"
	"os"
)

func main() {
	// Setup logger
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	// Init
	cfg := config.LoadConfig()
	db := dbl.NewDB(cfg)
	router := http.NewServeMux()

	// Repositories
	productRepo := product.NewProductRepository(db)

	// Handlers
	product.NewProductHandler(router, &product.ProductHandlerDeps{
		ProductRepo: productRepo,
	})

	// Middleware
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	// Configuration server
	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	// Start server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
