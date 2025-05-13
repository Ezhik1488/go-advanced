package main

import (
	"log"
	"net/http"
	"validation-api/config"
	"validation-api/internal/verify"
	"validation-api/pkg/email_client"
	"validation-api/pkg/storage"
)

func main() {
	cfg := config.LoadConfig()
	router := http.NewServeMux()
	emailClient := email_client.NewEmailClient(cfg)
	localStorage := storage.NewLocalStorage("data.json")
	verify.NewVerifyHandler(router, emailClient, localStorage)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return
	}
}
