package user

import "order-api/dbl"

type UserRepository struct {
	DB *dbl.DB
}

func NewUserRepository(db *dbl.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (repo *UserRepository) FindByPhone(phone string) (*User, error) {
	var user User
	result := repo.DB.DB.First(&user, "phone_number = ?", phone)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepository) FindBySessionID(sessionID string) (*User, error) {
	var user User
	result := repo.DB.DB.First(&user, "session_id = ?", sessionID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepository) Create(user *User) error {
	result := repo.DB.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *UserRepository) UpdateSessionID(user *User) error {
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
