package product

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Article     int            `json:"article" gorm:"not null,uniqueIndex" validate:"required,min=10"`
	Name        string         `json:"name" gorm:"not null" validate:"required,min=1,max=255"`
	Price       float64        `json:"price" gorm:"not null" validate:"required,min=0"`
	Description string         `json:"description" gorm:"null"`
	Image       pq.StringArray `json:"image" gorm:"type:text[]"`
}

func NewProduct(product *ProductCreateRequest) *Product {
	return &Product{
		Article:     product.Article,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Image:       product.Image,
	}
}
