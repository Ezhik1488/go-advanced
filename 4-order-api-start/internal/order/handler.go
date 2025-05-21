package order

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"order-api/config"
	"order-api/pkg/middleware"
	"order-api/pkg/req"
	"order-api/pkg/res"
	"strconv"
)

type Response map[string]interface{}

type OrderHandler struct {
	OrderService OrderServiceInt
	config       *config.Config
}

func NewOrderHandler(router *http.ServeMux, orderService OrderServiceInt, cfg *config.Config) *OrderHandler {
	handler := &OrderHandler{orderService, cfg}
	router.Handle("POST /order", middleware.Auth(handler.CreateOrder(), cfg))
	router.Handle("GET /order/{id}", middleware.Auth(handler.GetByID(), cfg))
	router.Handle("GET /my-orders", middleware.Auth(handler.GetByUserID(), cfg))
	return handler
}

func (h *OrderHandler) CreateOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[CreateOrderRequest](w, r)
		if err != nil {
			res.JSON(w, Response{
				"result": "Invalid body",
				"status": http.StatusBadRequest,
			},
				http.StatusBadRequest)
			return
		}
		userID, ok := r.Context().Value(middleware.ContextUserID).(uint)
		if !ok {
			res.JSON(w, Response{
				"result": "Something went wrong",
				"status": http.StatusInternalServerError,
			}, http.StatusInternalServerError)
			return
		}

		createdOrder, err := h.OrderService.CreateOrder(userID, body)
		if err != nil {
			res.JSON(w, Response{
				"result": "Error when creating an order",
				"status": http.StatusInternalServerError,
			},
				http.StatusInternalServerError)
			return
		}
		res.JSON(w, Response{
			"result": Response{
				"detail":         "The order was successfully created",
				"total_cost":     createdOrder.TotalCost,
				"products_count": createdOrder.ProductsCount,
			},
			"status": http.StatusCreated,
		},
			http.StatusCreated)
	}
}

func (h *OrderHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productID, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
		if err != nil || productID < 1 {
			res.JSON(w, Response{
				"result": "Invalid path params - ID",
				"status": http.StatusBadRequest,
			},
				http.StatusBadRequest)
			return
		}
		userID, ok := r.Context().Value(middleware.ContextUserID).(uint)
		if !ok {
			res.JSON(w, Response{
				"result": "Something went wrong",
				"status": http.StatusInternalServerError,
			}, http.StatusInternalServerError)
			return
		}

		order, err := h.OrderService.FindByID(uint(productID), userID)
		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, ErrForbidden) {
			res.JSON(w, Response{
				"result": "Order not found",
				"status": http.StatusNotFound,
			},
				http.StatusNotFound)
			return
		}

		if err != nil {
			res.JSON(w, Response{
				"result": "Something went wrong",
				"status": http.StatusInternalServerError,
			},
				http.StatusInternalServerError)
			return
		}

		res.JSON(w, GetOrderByIDResponse{
			Order:  *order,
			Status: http.StatusOK,
		},
			http.StatusOK)
	}
}

func (h *OrderHandler) GetByUserID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(middleware.ContextUserID).(uint)
		if !ok {
			res.JSON(w, Response{
				"result": "Something went wrong",
				"status": http.StatusInternalServerError,
			}, http.StatusInternalServerError)
			return
		}
		result, err := h.OrderService.FindByUserID(userID)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res.JSON(w, Response{
				"result": "Order not found",
				"status": http.StatusNotFound,
			},
				http.StatusNotFound)
			return
		}
		res.JSON(w, GetOrderByUserIDResponse{
			Orders: result,
			Status: http.StatusOK,
		},
			http.StatusOK)
	}
}
