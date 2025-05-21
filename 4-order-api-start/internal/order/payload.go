package order

import "order-api/internal/core/models"

type CreateOrderRequest struct {
	Products []int `json:"products" validate:"required,gt=0,dive,gt=10000000000"`
}

type GetOrderByIDResponse struct {
	Order  models.Order `json:"products"`
	Status int          `json:"status"`
}

type GetOrderByUserIDResponse struct {
	Orders []models.Order `json:"orders"`
	Status int            `json:"status"`
}
