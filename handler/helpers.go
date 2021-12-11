package handler

import (
	"net/http"

	"golang-memory/logger"
)

func RecoverPanic(r *http.Request, response *ApiResponse) {
	if rec := recover(); rec != nil {
		logger.Error.Printf("[RECOVERED PANIC] %+v", rec)
	}
	logger.Info.Printf("Received %s %s [Rsp: {%+v}]", r.Method, r.URL.Path, response)
}

func StoreData(key, value string) {
	InMemDB[key] = value
}
