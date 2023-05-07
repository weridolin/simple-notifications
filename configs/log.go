package configs

import (
	"log"
	"os"
)

var logger *log.Logger = nil

func init() {
	if logger == nil {
		logger = log.New(os.Stdout, "<simple-notification> ", log.Lshortfile|log.Ldate|log.Lmicroseconds)
	}
}
