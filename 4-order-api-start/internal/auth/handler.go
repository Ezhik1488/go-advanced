package auth

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"order-api/pkg/req"
	"order-api/pkg/res"
)

type Response map[string]interface{}

type AuthHandler struct {
	AuthService *AuthService
}

func NewAuthHandler(router *http.ServeMux, authService *AuthService) *AuthHandler {
	handler := &AuthHandler{AuthService: authService}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/verify", handler.VerifyCode())
	return handler
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](w, r)
		if err != nil {
			return
		}
		sessionID, err := h.AuthService.Login(body.Number)
		if err != nil {
			res.JSON(w, Response{
				"result": err.Error(),
				"status": http.StatusInternalServerError,
			},
				http.StatusInternalServerError)
			return
		}
		data := LoginResponse{
			SessionId: sessionID,
		}
		res.JSON(w, data, http.StatusOK)
	}
}

func (h *AuthHandler) VerifyCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[VerifyRequest](w, r)
		if err != nil {
			return
		}
		token, err := h.AuthService.VerifyCode(body.SessionId, body.Code)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				res.JSON(w, Response{
					"result": "Invalid session ID",
					"status": http.StatusUnauthorized},
					http.StatusUnauthorized)
				return
			}
			res.JSON(w, Response{
				"result": err.Error(),
				"status": http.StatusUnauthorized},
				http.StatusUnauthorized)
			return
		}
		data := VerifyResponse{
			Token: token,
		}
		res.JSON(w, data, http.StatusOK)
	}
}
