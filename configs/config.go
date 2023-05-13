package config

import (
	"os"
	"strconv"
)

var ConfigInstance *Config

type Config struct {
	PlatFormWhiteList []string
	DBUri             string
	// NotifierBrokerUri              string
	DefaultTickerMaxSchedulerCount int
	DefaultMaxTickerCount          int
	RabbitMQUri                    string
	DefaultSender                  string // EmailConsumer默认发送账号，没有指定的话
	DefaultPWD                     string // EmailConsumer默认发送账号密码，没有指定的话
	EmailConsumerCount             int    // EmailConsumer的数量
}

func StrToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func GetAppConfig() *Config {
	if ConfigInstance == nil {
		ConfigInstance = &Config{
			PlatFormWhiteList: make([]string, 0),
			DBUri:             os.Getenv("DBUri"),
			// NotifierBrokerUri:              os.Getenv("NotifierBrokerUri"),
			RabbitMQUri:                    os.Getenv("RabbitMQUri"),
			DefaultTickerMaxSchedulerCount: StrToInt(os.Getenv("DefaultTickerMaxSchedulerCount")),
			DefaultMaxTickerCount:          StrToInt(os.Getenv("DefaultMaxTickerCount")),
			DefaultSender:                  os.Getenv("DefaultSender"),
			DefaultPWD:                     os.Getenv("DefaultPWD"),
			EmailConsumerCount:             StrToInt(os.Getenv("EmailConsumerCount")),
		}
	}
	return ConfigInstance
}
