package product

import (
	"gorm.io/gorm/clause"
	"order-api/dbl"
)

type ProductRepository struct {
	DB *dbl.DB
}

func NewProductRepository(db *dbl.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (repo *ProductRepository) FindALL() ([]Product, error) {
	var products []Product
	result := repo.DB.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (repo *ProductRepository) FindByID(id uint) (*Product, error) {
	var product *Product
	result := repo.DB.DB.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *ProductRepository) Create(product *Product) error {
	result := repo.DB.DB.Create(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *ProductRepository) Update(product *Product) (*Product, error) {
	result := repo.DB.DB.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *ProductRepository) Delete(id uint) (int, error) {
	result := repo.DB.DB.Delete(&Product{}, id)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}
