package main

import (
	"http-job-que-system/logger"
	"http-job-que-system/server"
	"log"
)

func main() {
	// Initialize logger
	err := logger.InitLogger()
	if err != nil {
		log.Fatal("Logger initialization failed: ", err)
	}
	var log = logger.Log
	log.Println("App started")

	// Start server
	server.Start()
}
