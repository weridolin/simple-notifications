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
	//email notify 相关配置
	DefaultSender               string // EmailConsumer默认发送账号，没有指定的话
	DefaultPWD                  string // EmailConsumer默认发送账号密码，没有指定的话
	EmailConsumerCount          int    // EmailConsumer的数量
	EmailMessageAckTimeOut      int    // EmailConsumer消息ack超市时间
	EmailMessageDlxExchangeName string // EmailConsumer消息死信交换机名称
	EmailMessageDlxQueueName    string // EmailConsumer消息死信队列名称
	EmailMessageExchangeName    string // EmailConsumer消息交换机名称
	EmailMessageQueueName       string // EmailConsumer消息队列名称
	MongoDbUri                  string // MongoDB连接地址
	MongoDbName                 string // SimpleNotification连接地址
	StorageType                 string // 存储类型
	StorageFileRelativePath     string // 结果存储文件路径

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
			EmailMessageAckTimeOut:         StrToInt(os.Getenv("EmailMessageAckTimeOut")), // EmailConsumer消息ack超市时间
			EmailMessageDlxExchangeName:    os.Getenv("EmailMessageDlxExchangeName"),      // EmailConsumer消息死信交换机名称
			EmailMessageDlxQueueName:       os.Getenv("EmailMessageDlxQueueName"),         // EmailConsumer消息死信队列名称
			EmailMessageExchangeName:       os.Getenv("EmailMessageExchangeName"),         // EmailConsumer消息交换机名称
			EmailMessageQueueName:          os.Getenv("EmailMessageQueueName"),            // EmailConsumer消息队列名称
			MongoDbUri:                     os.Getenv("MongoDbUri"),                       // MongoDB连接地址
			MongoDbName:                    os.Getenv("MongoDbName"),                      // SimpleNotification连接地址
			StorageType:                    os.Getenv("StorageType"),                      // 存储类型
			StorageFileRelativePath:        os.Getenv("StorageFileRelativePath"),          // 结果存储文件路径
		}
	}
	return ConfigInstance
}
