package server

import (
	"fmt"
	"http-job-que-system/logger"
	"net/http"
)

func Start() {
	var log = logger.Log
	startupMsg := "Server listening on port 8080"

	log.Println(startupMsg)
	fmt.Println(startupMsg)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
