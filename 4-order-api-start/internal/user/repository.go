package user

import (
	"order-api/dbl"
	"order-api/internal/core/models"
)

type UserRepositoryInt interface {
	FindByPhone(phone string) (*models.User, error)
	FindBySessionID(sessionID string) (*models.User, error)
	Create(user *models.User) error
	UpdateSessionID(user *models.User) error
	Delete(id uint) (int, error)
}

type UserRepository struct {
	DB *dbl.DB
}

func NewUserRepository(db *dbl.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (repo *UserRepository) FindByPhone(phone string) (*models.User, error) {
	var user models.User
	result := repo.DB.DB.First(&user, "phone_number = ?", phone)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepository) FindBySessionID(sessionID string) (*models.User, error) {
	var user models.User
	result := repo.DB.DB.First(&user, "session_id = ?", sessionID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepository) Create(user *models.User) error {
	result := repo.DB.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *UserRepository) UpdateSessionID(user *models.User) error {
	user.GenerateSessionID()
	result := repo.DB.DB.Model(user).Update("session_id", user.SessionID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *UserRepository) Delete(id uint) (int, error) {
	result := repo.DB.DB.Delete(id)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}
