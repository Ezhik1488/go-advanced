package order

type CreateOrderRequest struct {
	Products []int `json:"products" validate:"required,gt=0,dive,gt=10000000000"`
}
