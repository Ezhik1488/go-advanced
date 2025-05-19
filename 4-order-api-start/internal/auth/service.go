package auth

import (
	"errors"
	"order-api/config"
	"order-api/internal/user"
	"order-api/pkg/jwt"
)

type AuthService struct {
	UserRepo *user.UserRepository
	JWT      *jwt.JWT
	Config   *config.Config
}

func NewAuthService(userRepo *user.UserRepository, jwt *jwt.JWT, cfg *config.Config) *AuthService {
	return &AuthService{UserRepo: userRepo, JWT: jwt, Config: cfg}
}

func (s *AuthService) Login(number string) (string, error) {
	existedUser, _ := s.UserRepo.FindByPhone(number)
	if existedUser == nil {
		newUser := &user.User{
			PhoneNumber: number,
		}
		newUser.GenerateSessionID()
		err := s.UserRepo.Create(newUser)
		if err != nil {
			return "", err
		}
		return newUser.SessionID, nil
	}
	err := s.UserRepo.UpdateSessionID(existedUser)
	if err != nil {
		return "", err
	}
	return existedUser.SessionID, nil
}

func (s *AuthService) VerifyCode(sessionID string, code int) (string, error) {
	foundedUser, err := s.UserRepo.FindBySessionID(sessionID)
	if err != nil {
		return "", err
	}
	if code != s.Config.Auth.VerifyCode {
		return "", errors.New("invalid verification code")
	}
	token, err := s.JWT.GenerateToken(foundedUser.ID, foundedUser.PhoneNumber, sessionID)
	if err != nil {
		return "", err
	}
	return token, nil
}
