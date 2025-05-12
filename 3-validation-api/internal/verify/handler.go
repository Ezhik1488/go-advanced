package verify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"validation-api/pkg/email_client"
	"validation-api/pkg/storage"
)

type VerifyHandler struct {
	ec *email_client.EmailClient
	ls *storage.LocalStorage
}

func NewVerifyHandler(router *http.ServeMux, email *email_client.EmailClient, storage *storage.LocalStorage) *VerifyHandler {
	handler := &VerifyHandler{ec: email, ls: storage}
	router.HandleFunc("POST /send", handler.SendVerifyEmail())
	router.HandleFunc("GET /verify/{hash}", handler.VerifyEmail())
	return handler
}

func (h *VerifyHandler) SendVerifyEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		recipient := h.ec.Cfg.Email.Recipient
		url := "http://localhost:8081/verify/" + recipient
		// Сохранить в JSON hash email: hash
		err := h.ls.Write(map[string]string{recipient: url})
		if err != nil {
			return
		}
		// Отправить Email
		err = h.ec.SendEmailWithTLS(recipient, url)
		if err != nil {
			return
		}
		// Отдать ответ клиенту
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("Verification email sent"))
		if err != nil {
			return
		}
	}
}

func (h *VerifyHandler) VerifyEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		data, err := h.ls.Read()
		if err != nil {
			fmt.Println(err)
		}
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			return
		}
	}
}
