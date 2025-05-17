package server

import (
	"log"
	"net/http"
)

func Start() {
	log.SetPrefix("server: ")
	log.SetFlags(0)

	log.Println("Starting server on :8080...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
