package order

import (
	"order-api/dbl"
	"order-api/internal/core/models"
)

type OrderRepositoryInt interface {
	GetByID(id uint) (*models.Order, error)
	GetByUserPhone(userID uint) ([]models.Order, error)
	Create(order *models.Order) error
}

type OrderRepository struct {
	DB *dbl.DB
}

func NewOrderRepository(db *dbl.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (repo *OrderRepository) GetByID(id uint) (*models.Order, error) {
	var order models.Order
	result := repo.DB.DB.First(&order, id)
	return &order, result.Error
}

func (repo *OrderRepository) GetByUserPhone(userID uint) ([]models.Order, error) {
	var orders []models.Order
	result := repo.DB.DB.Where("user_id = ?", userID).Find(&orders)
	return orders, result.Error
}

func (repo *OrderRepository) Create(order *models.Order) error {
	result := repo.DB.DB.Create(&order)
	return result.Error
}
