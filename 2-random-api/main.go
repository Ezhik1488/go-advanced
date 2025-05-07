package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	router.HandleFunc("/random", random)

	server.ListenAndServe()
}

func random(w http.ResponseWriter, req *http.Request) {
	randomNum := rand.Intn(6) + 1
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%d", randomNum)))
}
