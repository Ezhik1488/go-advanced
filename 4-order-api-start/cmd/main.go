package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"order-api/config"
	"order-api/dbl"
	"order-api/internal/auth"
	"order-api/internal/order"
	"order-api/internal/product"
	"order-api/internal/user"
	"order-api/pkg/jwt"
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
	jwtCust := jwt.NewJWT(cfg)

	// Repositories
	productRepo := product.NewProductRepository(db)
	userRepo := user.NewUserRepository(db)
	orderRepo := order.NewOrderRepository(db)

	// Services
	authService := auth.NewAuthService(userRepo, jwtCust, cfg)
	orderService := order.NewOrderService(orderRepo, userRepo)

	// Handlers
	product.NewProductHandler(router, &product.ProductHandlerDeps{
		ProductRepo: productRepo,
		Config:      cfg,
	})

	auth.NewAuthHandler(router, authService)
	order.NewOrderHandler(router, orderService, cfg)

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
