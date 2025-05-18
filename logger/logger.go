package logger

import (
	"log"
	"os"
	"path/filepath"
)

const logFileName = "log.txt"

var Log *log.Logger

func InitLogger() error {
	// Use current working directory for log file path
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	logPath := filepath.Join(cwd, logFileName)

	// Open file in append mode, create if not exists
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	// Create the logger
	Log = log.New(file, "APP: ", log.Ldate|log.Ltime|log.Lshortfile)
	Log.Println("Logger initialized")
	return nil
}
