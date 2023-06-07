package main

import (
	"os"

	"github.com/zeromicro/go-zero/core/logx"
	"gopkg.in/yaml.v3"
)

type EmailNotifyMessage struct {
	Sender   string   `json:"sender"`
	PWD      string   `json:"pwd"`
	Content  string   `json:"content"`
	Receiver []string `json:"receiver"`
	Subject  string   `json:"subject"`
}

type EmailConfigConsumer struct {
	Sender   string `json:"sender" yaml:"sender"`
	PWD      string `json:"pwd" yaml:"pwd"`
	Interval int    `json:"interval" yaml:"interval"` //两次消费之间的间隔
}

type RabbitMQConfig struct {
	RabbitMQUri           string `json:"rabbitmqUri" yaml:"rabbitmqUri"`
	EmailMessageQueueName string `json:"emailMessageQueueName" yaml:"emailMessageQueueName"`
}

type ConsumerConfig struct {
	EmailConsumerConfig EmailConfigConsumer `json:"emailConsumerConfig" yaml:"emailConsumerConfig"`
	RabbitMQConfig      RabbitMQConfig      `json:"rabbitmqConfig" yaml:"rabbitmqConfig"`
}

func (config *ConsumerConfig) FromYamlFile(configFile string) *ConsumerConfig {
	dataBytes, err := os.ReadFile(configFile)
	if err != nil {
		logx.Error("读取文件失败：", err)
		return nil
	}
	var ConfigInstance = &ConsumerConfig{}
	err = yaml.Unmarshal(dataBytes, ConfigInstance)
	if err != nil {
		logx.Error("yaml反序列化失败:", err)
		return nil
	}
	return ConfigInstance
}
