package product

import (
	"gorm.io/gorm/clause"
	"order-api/dbl"
	"order-api/internal/core/models"
)

type ProductRepositoryInt interface {
	FindALL() ([]models.Product, error)
	FindByID(id uint) (*models.Product, error)
	FindByArticle(article int) (*models.Product, error)
	Create(product *models.Product) error
	Update(product *models.Product) (*models.Product, error)
	Delete(id uint) (int, error)
}

type ProductRepository struct {
	DB *dbl.DB
}

func NewProductRepository(db *dbl.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (repo *ProductRepository) FindALL() ([]models.Product, error) {
	var products []models.Product
	result := repo.DB.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (repo *ProductRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	result := repo.DB.DB.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (repo *ProductRepository) FindByArticle(article int) (*models.Product, error) {
	var product models.Product
	result := repo.DB.DB.Where("article = ?", article).First(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (repo *ProductRepository) Create(product *models.Product) error {
	result := repo.DB.DB.Create(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *ProductRepository) Update(product *models.Product) (*models.Product, error) {
	result := repo.DB.DB.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *ProductRepository) Delete(id uint) (int, error) {
	result := repo.DB.DB.Delete(&models.Product{}, id)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}
