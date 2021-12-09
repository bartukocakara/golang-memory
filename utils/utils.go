package utils

import (
	"encoding/json"
	"net/http"
	"os"
)

func GetEnv(key string, value string) string {
	if req, fine := os.LookupEnv(key); fine {
		return req
	}
	return value
}

func CreateResponse(w http.ResponseWriter, response interface{}) {
	if err := json.NewEncoder(w).Encode(response); 
	err != nil {
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		w.WriteHeader(http.StatusInternalServerError)
	}
}

