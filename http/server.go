package http

import (
	"log"
	"net/http"
)

func Start(portNumber string) bool {
	log.Println("starting server...")
	port := ":" + portNumber
	log.Fatal(http.ListenAndServe(port, Router))

	return true
}
