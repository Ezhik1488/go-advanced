package order

import (
	"errors"
	"order-api/internal/core/models"
	"order-api/internal/product"
	"order-api/internal/user"
)

type OrderServiceInt interface {
	CreateOrder(userID uint, body *CreateOrderRequest) (*models.Order, error)
	FindByID(orderID, userID uint) (*models.Order, error)
	FindByUserID(userID uint) ([]models.Order, error)
}

type OrderService struct {
	OrderRepo   OrderRepositoryInt
	UserRepo    user.UserRepositoryInt
	ProductRepo product.ProductRepositoryInt
}

func NewOrderService(orderRepo *OrderRepository, userRepo user.UserRepositoryInt, productRepo product.ProductRepositoryInt) *OrderService {
	return &OrderService{
		OrderRepo:   orderRepo,
		UserRepo:    userRepo,
		ProductRepo: productRepo,
	}
}

func (s *OrderService) CreateOrder(userID uint, body *CreateOrderRequest) (*models.Order, error) {
	var totalCost float64
	var productCount int
	order := models.Order{
		UserID: userID,
	}
	// Найти и добавить товар в заказ по артикулу
	for _, prodArticle := range body.Products {
		goods, err := s.ProductRepo.FindByArticle(prodArticle)
		if err != nil {
			continue
		}
		order.Products = append(order.Products, *goods)

		totalCost += goods.Price
		productCount++
	}
	if len(order.Products) == 0 {
		return nil, errors.New("empty cart")
	}
	order.TotalCost = totalCost
	order.ProductsCount = productCount

	return &order, s.OrderRepo.Create(order)
}

func (s *OrderService) FindByID(orderID, userID uint) (*models.Order, error) {
	result, err := s.OrderRepo.GetByID(orderID)
	if err != nil {
		return nil, err
	}

	if result.UserID != userID {
		return nil, ErrForbidden
	}
	return result, nil

}

func (s *OrderService) FindByUserID(userID uint) ([]models.Order, error) {
	return s.OrderRepo.GetByUserID(userID)
}
