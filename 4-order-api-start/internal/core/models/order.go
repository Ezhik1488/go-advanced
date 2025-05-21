package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ProductsCount int       `json:"products_count"`
	TotalCost     float64   `json:"total_cost"`
	UserID        uint      `json:"user_id"`
	Products      []Product `gorm:"many2many:order_products;" json:"products"`
}
