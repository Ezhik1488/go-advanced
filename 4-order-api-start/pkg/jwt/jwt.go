package jwt

import (
	jwt2 "github.com/golang-jwt/jwt/v5"
	"order-api/config"
)

type JWTData struct {
	UserPhone string
}

type JWT struct {
	secret string
}

func NewJWT(cfg *config.Config) *JWT {
	return &JWT{secret: cfg.Auth.Secret}
}

func (j *JWT) GenerateToken(userID uint, userPhone, sessionID string) (string, error) {
	token := jwt2.NewWithClaims(jwt2.SigningMethodHS256, jwt2.MapClaims{
		"user_id":    userID,
		"user_phone": userPhone,
		"session_id": sessionID,
	})
	signToken, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}
	return signToken, nil
}

func (j *JWT) VerifyToken(tokenString string) (bool, *JWTData) {
	parse, err := jwt2.Parse(tokenString, func(token *jwt2.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		return false, nil
	}
	
	claims, ok := parse.Claims.(jwt2.MapClaims)
	if !ok {
		return false, nil
	}
	rawUserPhone, existed := claims["user_phone"]
	if !existed {
		return false, nil
	}
	userPhone, ok := rawUserPhone.(string)
	if !ok {
		return false, nil
	}

	return parse.Valid, &JWTData{UserPhone: userPhone}
}
