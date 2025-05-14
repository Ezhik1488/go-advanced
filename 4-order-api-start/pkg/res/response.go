package res

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, data any, status int) *http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return nil
	}
	return &w
}
