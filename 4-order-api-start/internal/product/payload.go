package product

import "github.com/lib/pq"

type ProductCreateRequest struct {
	Article     int            `json:"article" validate:"required,min=10"`
	Name        string         `json:"name" validate:"required,min=1,max=255"`
	Price       float64        `json:"price" validate:"required,min=0"`
	Description string         `json:"description,omitempty"`
	Image       pq.StringArray `json:"image,omitempty"`
}

type ProductUpdateRequest struct {
	Name        string         `json:"name" validate:"min=1,max=255"`
	Price       float64        `json:"price" validate:"min=0"`
	Description string         `json:"description,omitempty"`
	Image       pq.StringArray `json:"image,omitempty"`
}
