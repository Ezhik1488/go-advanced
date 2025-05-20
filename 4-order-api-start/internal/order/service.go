package order

import "order-api/pkg/di"

type OrderService struct {
	OrderRepo *OrderRepository
	UserRepo  di.UserRepositoryInt
}

func NewOrderService(orderRepo *OrderRepository, userRepo di.UserRepositoryInt) *OrderService {
	return &OrderService{
		OrderRepo: orderRepo,
		UserRepo:  userRepo,
	}
}

func (service *OrderService) CreateOrder(order *Order) error {
	return service.OrderRepo.Create(order)
}
