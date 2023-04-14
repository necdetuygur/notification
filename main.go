package main

import (
	"net/http"
	handler "notification/api"
	"os"
)

func main() {
	http.HandleFunc("/", handler.Handler)
	port := "3333"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	http.ListenAndServe(":"+port, nil)
}
