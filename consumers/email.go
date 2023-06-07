package main

import (
	"encoding/json"

	"github.com/weridolin/simple-vedio-notifications/clients"
	"github.com/weridolin/simple-vedio-notifications/tools"
	"github.com/zeromicro/go-zero/core/logx"
)

type EmailConsumer struct {
	MQClient       *clients.RabbitMQ
	DefaultSender  string
	DefaultPWD     string
	Interval       int //两次发送的时间间隔
	ConsumerConfig ConsumerConfig
}

func NewEmailConsumer(id string, config ConsumerConfig) *EmailConsumer {
	return &EmailConsumer{
		MQClient:       clients.NewRabbitMQ(id, config.RabbitMQConfig.RabbitMQUri),
		DefaultSender:  config.EmailConsumerConfig.Sender,
		DefaultPWD:     config.EmailConsumerConfig.PWD,
		Interval:       config.EmailConsumerConfig.Interval,
		ConsumerConfig: config,
	}
}

func (c *EmailConsumer) Start() {
	c.MQClient.ReceiveTopic(c.ConsumerConfig.RabbitMQConfig.EmailMessageQueueName, c.OnMessage)
}

func (c *EmailConsumer) OnMessage(message []byte) error {
	logx.Info("email consumer-> ", c.MQClient.ID, " get message from rabbitmq ->")
	var err error
	EmailNotifyMessage := EmailNotifyMessage{Subject: "bilibili up 订阅结果通知"}
	err = json.Unmarshal(message, &EmailNotifyMessage)
	if err != nil {
		logx.Info("反序列化失败", err)
	}
	logx.Info("email consumer get message from rabbitmq ->", EmailNotifyMessage)
	err = tools.SendEmail(EmailNotifyMessage.Receiver, EmailNotifyMessage.Subject,
		EmailNotifyMessage.Content, EmailNotifyMessage.Sender, EmailNotifyMessage.PWD)
	return err
}
