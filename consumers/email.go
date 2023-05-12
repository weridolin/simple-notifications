package consumers

import (
	"github.com/weridolin/simple-vedio-notifications/clients"
	"github.com/weridolin/simple-vedio-notifications/common"
	config "github.com/weridolin/simple-vedio-notifications/configs"
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
	// c.MQClient.QueueName = "consumer.email.test"
	// c.MQClient.CreateExchangeAndBindQueue()
	// c.MQClient.CreateExchange(common.EmailExchangeName, "topic").
	// 	CreateQueue(common.EmailMessageQueueName, true)
	c.MQClient.ReceiveTopic(common.EmailMessageQueueName, c.OnMessage)
}

func (c *EmailConsumer) OnMessage(message []byte) error {
	var err error
	logger.Println("email consumer get message from rabbitmq ->", string(message))
	return err
}
