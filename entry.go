package main

import (
	"fmt"

	// schedulers "github.com/weridolin/simple-vedio-notifications/schedulers"
	"time"

	"github.com/robfig/cron/v3"
	config "github.com/weridolin/simple-vedio-notifications/configs"
	"github.com/weridolin/simple-vedio-notifications/consumers"
	"github.com/weridolin/simple-vedio-notifications/database"
	"github.com/weridolin/simple-vedio-notifications/schedulers"
	"github.com/weridolin/simple-vedio-notifications/tools"
	// httpConsumers "github.com/weridolin/simple-vedio-notifications/servers/http/consumers"
)

// type Context struct {
// 	tp     *schedulers.TickerPool
// 	config *configs.Config
// }

// func NewContext(tp *schedulers.TickerPool, config *configs.Config) *Context {
// 	return &Context{
// 		tp:     tp,
// 		config: config,
// 	}
// }

func TestCron() {
	c := cron.New()
	i := 1

	c.Start()

	EntryID, err := c.AddFunc("*/1 * * * *", func() {
		fmt.Println(time.Now(), "每分钟执行一次", i)
		i++
	})
	fmt.Println(time.Now(), EntryID, err)

	c.Stop()
	time.Sleep(time.Minute * 1)
	time.Sleep(time.Minute * 2)
}

func Setup() {
	//  DB迁移
	DB := database.GetDB()
	// AutoMigrate : 1 添加字段会自定添加  2.删除字段原来的会保存，新的记录会默认删除的字段为空
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&database.User{},
		&database.Scheduler{},
		&database.Task{},
		&database.EmailNotifier{},
		&database.EmailNotifierTask{},
		&database.SchedulerTask{},
	)
}

func main() {
	Setup()

	sync := schedulers.NewSynchronizer()
	go sync.Start()
	config := config.GetAppConfig()
	// ctx := context.WithValue(context.Background(), "tp", schedulers.NewTickerPool(config.DefaultMaxTickerCount))
	// uuid := tools.GetUUID()
	// manager := schedulers.NewSchedulerManager(ctx, uuid)
	for i := 0; i < config.EmailConsumerCount; i++ {
		emailConsumer := consumers.NewEmailConsumer(tools.GetUUID())
		emailConsumer.MQClient.SetDLX()
		go emailConsumer.Start()
	}

	// rabbitMq := clients.NewRabbitMQ(tools.GetUUID())
	// rabbitMq.CreateExchange(common.EmailExchangeName, "topic").
	// 	CreateQueue(common.EmailMessageQueueName, true).
	// 	ExchangeBindQueue(common.EmailMessageQueueName, "*.email.*", common.EmailExchangeName)

	// http.Start()
	time.Sleep(time.Minute * 10)
}
