package consumers

import (
	"encoding/json"

	"github.com/weridolin/simple-vedio-notifications/clients"
	"github.com/weridolin/simple-vedio-notifications/common"
	config "github.com/weridolin/simple-vedio-notifications/configs"
	"github.com/weridolin/simple-vedio-notifications/tools"
)

var logger = config.GetLogger()

type EmailConsumer struct {
	MQClient      *clients.RabbitMQ
	DefaultSender string
	DefaultPWD    string
	Interval      uint //两次发送的时间间隔
}

func NewEmailConsumer() *EmailConsumer {
	var AppConfig = config.GetAppConfig()
	return &EmailConsumer{
		MQClient:      clients.NewRabbitMQ(),
		DefaultSender: AppConfig.DefaultSender,
		DefaultPWD:    AppConfig.DefaultPWD,
		Interval:      1,
	}
}

func (c *EmailConsumer) Start() {
	c.MQClient.ReceiveTopic(common.EmailMessageQueueName, c.OnMessage)
}

func (c *EmailConsumer) OnMessage(message []byte) error {
	var err error
	// logger.Println("email consumer get message from rabbitmq ->", string(message))
	EmailNotifyMessage := EmailNotifyMessage{Subject: "bilibili up 订阅结果通知"}
	err = json.Unmarshal(message, &EmailNotifyMessage)
	if err != nil {
		logger.Println("反序列化失败", err)
	}
	logger.Println("email consumer get message from rabbitmq ->", EmailNotifyMessage)
	tools.SendEmail(EmailNotifyMessage.Receiver, EmailNotifyMessage.Subject, EmailNotifyMessage.Content, EmailNotifyMessage.Sender, EmailNotifyMessage.PWD)
	return err
}
