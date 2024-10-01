package logs

import (
	"log"
	"os"
	"path/filepath"
)

var logger *log.Logger

func InitLogger() {
	logDir := "./logs"
	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}

	logFilePath := filepath.Join(logDir, "app.log")
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	logger = log.New(file, "APP ", log.Ldate|log.Ltime|log.Lshortfile)
}

func GetLogger() *log.Logger {
	if logger == nil {
		InitLogger()
	}
	return logger
}
