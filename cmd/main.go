package main

import (
	"log"
	"net/http"

	"github.com/posolwar/softweather-test/internal/controller/http/handlers"
	"github.com/posolwar/softweather-test/internal/controller/http/middlewares"
	"github.com/posolwar/softweather-test/internal/helpers"
)

func main() {
	serverPort := helpers.DefaultPort

	log.Print("Server starting on " + serverPort)

	http.HandleFunc("/api/v1/math", middlewares.MethodCheck(http.MethodPost, middlewares.AccessCheck(handlers.Calculate)))
	if err := http.ListenAndServe(serverPort, nil); err != nil {
		log.Fatal(err.Error())
	}
}
