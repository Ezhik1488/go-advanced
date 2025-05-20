package order

import (
	"gorm.io/gorm"
	"order-api/internal/product"
)

type Order struct {
	gorm.Model
	NumberProducts int               `json:"number_products"`
	TotalCost      float64           `json:"total_cost"`
	UserID         uint              `json:"user_id"`
	Products       []product.Product `gorm:"many2many:order_products;" json:"products"`
}
