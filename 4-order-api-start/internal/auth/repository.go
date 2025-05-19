package auth

import "order-api/dbl"

type AuthRepository struct {
	DB *dbl.DB
}

func NewAuthRepository(db *dbl.DB) *AuthRepository {}
