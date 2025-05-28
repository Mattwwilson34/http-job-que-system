// Package logger provides a custom logger that writes application logs
// to a file in the current working directory. It exposes a globally accessible
// log.Logger instance that can be used throughout the application.
package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const logFileName = "log.txt"

// Log is the globally accessible logger instance. It is initialized by calling InitLogger.
var Log *log.Logger

// Log structure for HTTP based log entries
type HttpLogMsg struct {
	FuncName   string
	Method     string
	UrlPath    string
	RemoteAddr string
	Status     int
	Message    string
}

// InitLogger initializes the global logger to write log output to a file named "log.txt"
// in the application's current working directory. If initialization succeeds, it returns nil;
// otherwise, it returns an error. This function should be called at the very start of the application.
//
// The logger is configured to include the date, time, and short file name in each log entry.
func InitLogger() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	logPath := filepath.Join(cwd, logFileName)

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	Log = log.New(file, "APP: ", log.Ldate|log.Ltime|log.Lshortfile)
	Log.Println("Logger initialized")
	return nil
}

// Format HttpLogMsg string output
func (h HttpLogMsg) String() string {
	return fmt.Sprintf("%s: %s %s from %s -> %d (%s)",
		h.FuncName,
		h.Method,
		h.UrlPath,
		h.RemoteAddr,
		h.Status,
		h.Message)
}
