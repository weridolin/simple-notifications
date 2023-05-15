/*
	message  ----> exchange  --根据exchange的类型/routing key-->  queue ----> consumer
*/

package clients

import (
	"github.com/streadway/amqp"
	config "github.com/weridolin/simple-vedio-notifications/configs"
)

var logger = config.GetLogger()

var appConfig = config.GetAppConfig()

//rabbitMQ结构体
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	// //队列名称
	// QueueName string
	// //交换机名称
	// Exchange string
	// //bind Key 名称
	// Key string
	// //连接信息
	Mqurl string
	// // 类型
	// Type_ string
	ID string
}

//创建结构体实例
func NewRabbitMQ(id string) *RabbitMQ {
	MQURL := config.GetAppConfig().RabbitMQUri
	instance := &RabbitMQ{Mqurl: MQURL, ID: id}
	var err error
	//获取connection
	instance.conn, err = amqp.Dial(instance.Mqurl)
	instance.failOnErr(err, "failed to connect rabbitmq!")
	//获取channel
	instance.channel, err = instance.conn.Channel()
	instance.channel.Qos(2, 0, false) // 限流，每个channel最多只能有200个未确认的消息,超过200个则不再推送给该channel上的 consumer，global参数表示是否应用到所有的channel上
	instance.failOnErr(err, "failed to open a channel")

	return instance
}

// func (r *RabbitMQ) CreateDl() {

// 	//声明死信交换器
// 	var dlxExchangeName =
// 	r.CreateExchange(dlxExchangeName, "direct")
// 	//声明队列
// 	_, err := r.channel.QueueDeclare("email.dlx.queue", true, false, false, false, nil)
// 	if err != nil {
// 		fmt.Println("set email dlx queue  err :", err)
// 		return
// 	}
// 	r.ExchangeBindQueue("email.dlx.queue", "email.dlx.queue", dlxExchangeName)

// }

func (r *RabbitMQ) CreateExchange(exchange, t string) *RabbitMQ {
	err := r.channel.ExchangeDeclare(
		exchange,
		//要改成topic
		t,
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an exchange")
	return r
}

func (r *RabbitMQ) CreateQueue(queue string, durable bool, argsParams map[string]interface{}) *RabbitMQ {
	// 消息持久化的3个条件
	// 1 投递消息的时候 durable 设置为 true，消息持久化；
	// 2 消息已经到达持久化交换器上；
	// 3 消息已经到达持久化的队列

	//创建队列
	_, err := r.channel.QueueDeclare(
		queue,
		durable, // 是否持久化
		false,   // 是否自动删除，当最后一个消费者断开连接之后队列是否自动被删除
		false,   // exclusive 是否排他，如果设置为true，那么只有创建这个队列的channel才能访问，其他channel访问会报错,同时channel关闭后队列会被删除
		false,   // no-wait 是否阻塞
		argsParams,
	)
	r.failOnErr(err, "Failed to declare a queue")
	return r
}

func (r *RabbitMQ) ExchangeBindQueue(queue, routeKey, exchange string) *RabbitMQ {
	// 绑定队列
	err := r.channel.QueueBind(
		queue,
		routeKey,
		exchange,
		false,
		nil)

	r.failOnErr(err, "Failed to bind Queue")
	return r
}

//断开channel 和 connection
func (r *RabbitMQ) Destory() {
	logger.Println("client destory", r.ID)
	r.channel.Close()
	r.conn.Close()
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		// log.Fatalf("%s:%s", message, err)
		logger.Panicln("rabbitmq error ->", err, message)
	}
}

//话题模式发送消息
func (r *RabbitMQ) Publish(exchange, key string, message []byte) {
	//2.发送消息
	err := r.channel.Publish(
		exchange,
		//要设置，队列的名称
		key,
		false, //如果为true 会根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返还给发送者
		false, //如果为true,当exchange发送消息到队列后发现队列上没有绑定消费者则会把消息返还给发送者
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         message,
			DeliveryMode: amqp.Persistent, //需要做持久化保留
		})
	r.failOnErr(err, "Failed to public message")

}

//话题模式接受消息
//要注意key,规则
//其中“*”用于匹配一个单词，“#”用于匹配多个单词（可以是零个）
//匹配 kuteng.* 表示匹配 kuteng.hello, kuteng.hello.one需要用kuteng.#才能匹配到
func (r *RabbitMQ) ReceiveTopic(queue string, callback func(msg []byte) error) {

	//消费消息
	messages, err := r.channel.Consume(
		queue,            //队列名称
		"email consumer", //消费者名称，区分多个消费者？
		false,            //是否自动应答 true为自动应答，只要消费了broker自定删除  false为手动应答,消费了之后需要手动告诉broker
		false,
		false,
		false, //是否阻塞
		nil,
	)
	r.failOnErr(err, "Failed to Consume")

	forever := make(chan bool)

	go func() {
		for d := range messages {
			err := callback(d.Body)
			if err != nil {
				logger.Println("consumer email notify message error ", err)
				d.Reject(false) // 拒绝消息，requeue为true会重新放回队列，否则放回死信队列
			} else {
				d.Ack(false) // false 确认当前消息   true确认所有未确认的消息  未确认的消息状态未un ack，等到客户端重新连接后会变为ready
			}
		}
	}()

	<-forever
	// r.Destory()
}
