package utils

import (
	"golang-memory/log"
	"net/http"
)

func MapPanic(r *http.Request, response *ApiResponse) {
	if recover := recover(); recover != nil {
		log.Error.Printf("[Recievedd PANIC] %+v", recover)
	}
	log.Info.Printf("Received %s %s [Rsp: {%+v}]", r.Method, r.URL.Path, response)
}

func StoreData(key, value string) {
	InMemStore[key] = value
}