package pkg

import (
	jwt2 "github.com/golang-jwt/jwt/v5"
	"order-api/config"
)

type JWT struct {
	secret string
}

func NewJWT(cfg *config.Config) *JWT {
	return &JWT{secret: cfg.Auth.Secret}
}

func (j *JWT) GenerateToken(userID int, userPhone string) (string, error) {
	token := jwt2.NewWithClaims(jwt2.SigningMethodHS256, jwt2.MapClaims{
		"user_id":    userID,
		"user_phone": userPhone,
	})
	signToken, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}
	return signToken, nil
}
