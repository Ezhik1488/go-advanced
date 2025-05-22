package e2e

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

func App() http.Handler {
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
	orderService := order.NewOrderService(orderRepo, userRepo, productRepo)
	productService := product.NewProductService()

	// Handlers
	product.NewProductHandler(router, &product.ProductHandlerDeps{
		ProductRepo:    productRepo,
		ProductService: productService,
		Config:         cfg,
	})

	auth.NewAuthHandler(router, authService, cfg)
	order.NewOrderHandler(router, orderService, cfg)

	// Middleware
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	return stack(router)
}
