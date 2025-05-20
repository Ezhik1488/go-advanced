package di

import (
	"order-api/internal/order"
	"order-api/internal/user"
)

type OrderRepositoryInt interface {
	GetByID(id uint) (*order.Order, error)
	GetByUserPhone(userID uint) ([]order.Order, error)
	Create(order *order.Order) error
}

type OrderServiceInt interface {
	CreateOrder(order *order.Order) error
}

type UserRepositoryInt interface {
	FindByPhone(phone string) (*user.User, error)
	FindBySessionID(sessionID string) (*user.User, error)
	Create(user *user.User) error
	UpdateSessionID(user *user.User) error
	Delete(id uint) (int, error)
}
