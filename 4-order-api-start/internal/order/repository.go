package order

import (
	"order-api/dbl"
)

type OrderRepository struct {
	DB *dbl.DB
}

func NewOrderRepository(db *dbl.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (repo *OrderRepository) GetByID(id uint) (*Order, error) {
	var order Order
	result := repo.DB.DB.First(&order, id)
	return &order, result.Error
}

func (repo *OrderRepository) GetByUserPhone(userID uint) ([]Order, error) {
	var orders []Order
	result := repo.DB.DB.Where("user_id = ?", userID).Find(&orders)
	return orders, result.Error
}

func (repo *OrderRepository) Create(order *Order) error {
	result := repo.DB.DB.Create(&order)
	return result.Error
}
