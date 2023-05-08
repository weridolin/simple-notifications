package config

import (
	"log"
	"os"
)

var logger *log.Logger = nil

func GetLogger() *log.Logger {
	if logger == nil {
		logger = log.New(os.Stdout, "<simple-notification> ", log.Lshortfile|log.Ldate|log.Lmicroseconds)
	}
	return logger
}
