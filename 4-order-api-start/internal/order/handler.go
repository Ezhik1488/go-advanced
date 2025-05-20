package order

import (
	"net/http"
	"order-api/config"
	"order-api/pkg/di"
	"order-api/pkg/middleware"
)

type OrderHandler struct {
	OrderService di.OrderServiceInt
	config       *config.Config
}

func NewOrderHandler(router *http.ServeMux, orderService di.OrderServiceInt, cfg *config.Config) *OrderHandler {
	handler := &OrderHandler{orderService, cfg}
	router.Handle("POST /order", middleware.Auth(handler.CreateOrder(), cfg))
	router.Handle("GET /order/{id}", middleware.Auth(handler.GetByID(), cfg))
	router.Handle("GET /my-orders", middleware.Auth(handler.GetByUserID(), cfg))
	return handler
}

func (h *OrderHandler) CreateOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *OrderHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *OrderHandler) GetByUserID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
