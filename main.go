package main

import (
	"fmt"
	"net/http"

	"golang-memory/config"
	"golang-memory/handler"
	"golang-memory/job"
	"golang-memory/logger"
)

func main() {
	job.Startjob()
	job.CheckExistingData()
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", handler.Home)
	mux.HandleFunc("/set", handler.Set)
	mux.HandleFunc("/get", handler.Get)
	mux.HandleFunc("/flush", handler.Flush)
	logger.Info.Println("Server starting...")
	logger.Info.Printf("Server started at %s", config.API_PORT)
	logger.Fatal.Println(http.ListenAndServe(fmt.Sprintf(":%s", config.API_PORT), mux))
}
