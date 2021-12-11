package utils

import (
	"encoding/json"
	"net/http"
	"os"
)

func GetEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func CreateResponse(w http.ResponseWriter, response interface{}) {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		w.WriteHeader(http.StatusInternalServerError)
	}
}
