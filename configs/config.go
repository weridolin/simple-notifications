package config

import (
	"os"
)

var ConfigInstance *Config

type Config struct {
	PlatFormWhiteList []string
	DBUri             string
	// NotifierBrokerUri              string
	DefaultTickerMaxSchedulerCount int
	DefaultMaxTickerCount          int
	RabbitMQUri                    string
	DefaultSender                  string
	DefaultPWD                     string
}

func GetAppConfig() *Config {
	if ConfigInstance == nil {
		ConfigInstance = &Config{
			PlatFormWhiteList: make([]string, 0),
			DBUri:             os.Getenv("DBUri"),
			// NotifierBrokerUri:              os.Getenv("NotifierBrokerUri"),
			RabbitMQUri:                    os.Getenv("RabbitMQUri"),
			DefaultTickerMaxSchedulerCount: 100,
			DefaultMaxTickerCount:          2,
			DefaultSender:                  os.Getenv("DefaultSender"),
			DefaultPWD:                     os.Getenv("DefaultPWD"),
		}
	}
	return ConfigInstance
}
