package product

import "order-api/internal/core/models"

type ProductServiceInt interface {
	NewProduct(product *ProductCreateRequest) *models.Product
}

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) NewProduct(product *ProductCreateRequest) *models.Product {
	return &models.Product{
		Article:     product.Article,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Image:       product.Image,
	}
}
