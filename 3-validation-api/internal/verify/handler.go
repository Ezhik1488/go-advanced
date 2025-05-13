package verify

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
	"validation-api/pkg/email_client"
	"validation-api/pkg/hash"
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

		var body VerifyRequest
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			fmt.Println(err.Error())
			return
		}

		validate := validator.New()
		err = validate.Struct(&body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			fmt.Println(err.Error())
			return
		}

		hashMail := hash.GetMD5Hash(body.Email)

		url := "http://localhost:8081/verify/" + hashMail
		// Сохранить в JSON hash email: hash
		err = h.ls.Write(map[string]string{hashMail: body.Email})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// Отправить Email
		err = h.ec.SendEmailWithTLS(body.Email, url)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// Отдать ответ клиенту
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Verification email sent"))
	}
}

func (h *VerifyHandler) VerifyEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		isValid := false
		fmt.Println(h.ls.Data)

		if _, ok := h.ls.Data[(strings.Split(r.URL.Path, "/"))[2]]; ok {
			isValid = true
			delete(h.ls.Data, (strings.Split(r.URL.Path, "/"))[2])
			fmt.Println(h.ls.Data)
			err := h.ls.Save()
			if err != nil {
				fmt.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
		}

		res := VerifyResponse{IsValid: isValid}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
