package order

import (
	"net/http"
	"order-api/config"
	"order-api/pkg/middleware"
	"order-api/pkg/req"
	"order-api/pkg/res"
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

	}
}

func (h *OrderHandler) GetByUserID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
