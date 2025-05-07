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

	err := server.ListenAndServe()
	if err != nil {
		return
	}

}

func random(w http.ResponseWriter, req *http.Request) {
	randomNum := rand.Intn(6)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(fmt.Sprintf("%d", randomNum)))
	if err != nil {
		return
	}
	return
}
