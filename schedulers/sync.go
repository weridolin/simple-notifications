/*
	从数据库中同步任务.触发方式:
		1.达到最大时间间隔
		2.有新的任务插入,信号触发

*/
package schedulers

import (
	"context"
	"time"

	"github.com/weridolin/simple-vedio-notifications/clients"
	config "github.com/weridolin/simple-vedio-notifications/configs"
	"github.com/weridolin/simple-vedio-notifications/database"
	"github.com/weridolin/simple-vedio-notifications/platforms/bilibili"
	"github.com/weridolin/simple-vedio-notifications/tools"
)

var appConfig = config.GetAppConfig()

type Synchronizer struct {
	MaxInterval      int
	MsgChannel       chan string
	StopSleepChannel chan bool
	SchedulerManager *SchedulerManager
}

func NewSynchronizer() *Synchronizer {
	config := config.GetAppConfig()
	ctx := context.WithValue(context.Background(), "tp", NewTickerPool(config.DefaultMaxTickerCount))
	uuid := tools.GetUUID()
	manager := NewSchedulerManager(ctx, uuid)
	return &Synchronizer{
		MaxInterval:      3000,
		SchedulerManager: manager,
	}
}

func (s *Synchronizer) Sleep(interval int) {
	// 把延时拆成 0.1 秒的间隔
	for i := 0; i < interval*10; i++ {
		// 判断数据库是否有更新
		select {
		case signal := <-s.StopSleepChannel:
			if signal {
				return
			}
		default:
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func (s *Synchronizer) Start() {
	s.Setup()
	s.MsgChannel = make(chan string)
	s.StopSleepChannel = make(chan bool)
	for {
		select { // 从通道中读取数据
		case msg := <-s.MsgChannel:
			if msg == "stop" {
				logger.Println("close synchronizer", msg)
				close(s.StopSleepChannel)
				close(s.MsgChannel)
				return
			}
		default:
			logger.Println("LOAD FROM DB")
			// 从数据库中加载数据,并同步到内存中
			records := s.LoadTaskFromDB()
			for _, record := range records {
				logger.Println("sync scheduler", "platform -> ", record.Platform, "task ->", record.Tasks, "Period ->", record.Period)
				scheduler := NewScheduler(tools.Period{Cron: record.Period}, record.Platform, 0, int(record.ID))
				for _, model := range record.Tasks {
					var task interface{}
					switch model.Platform {
					case "bilibili":
						task = bilibili.NewBiliBiliTask(tools.Period{Cron: record.Period}, model.Ups, model.ID, model.Name, model.Description, model.EmailNotifiers)
					}
					if task != nil {
						scheduler.Tasks = append(scheduler.Tasks, task)
					}
				}
				_, err := s.SchedulerManager.AddScheduler(scheduler, true)
				if err != nil {
					logger.Println("add scheduler error", err)
				}
			}
			s.Sleep(s.MaxInterval)
		}

	}
}

func (s *Synchronizer) LoadTaskFromDB() []*database.Scheduler {
	// var schedulers  []Scheduler
	// 1.从数据库中加载所有的scheduler
	schedulers, err := database.FullQuerySchedulers(&database.Scheduler{Active: true}, 1, 100)
	if err != nil {
		logger.Panicln("load scheduler from db error: ", err)
	}
	logger.Println("load scheduler finish", schedulers)
	return schedulers
}

func (s *Synchronizer) Stop() {
	s.StopSleepChannel <- true
	s.MsgChannel <- "stop"
}

func (s *Synchronizer) SyncAtOnce() {
	s.StopSleepChannel <- true
}

func (s *Synchronizer) Setup() {
	rabbitMq := clients.NewRabbitMQ(tools.GetUUID())
	//创建1个email notify 死信交换机和队列
	rabbitMq.CreateExchange(appConfig.EmailMessageDlxExchangeName, "direct").
		CreateQueue(appConfig.EmailMessageDlxQueueName, true, nil).
		ExchangeBindQueue(appConfig.EmailMessageDlxQueueName, appConfig.EmailMessageDlxQueueName, appConfig.EmailMessageDlxExchangeName)

	// 配置队列参数
	var dlxExchangeName = appConfig.EmailMessageDlxExchangeName
	argsQue := make(map[string]interface{})
	//添加死信队列交换器属性
	argsQue["x-dead-letter-exchange"] = dlxExchangeName
	//指定死信队列的路由key，不指定使用队列路由键
	argsQue["x-dead-letter-routing-key"] = appConfig.EmailMessageDlxQueueName
	//添加过期时间
	argsQue["x-message-ttl"] = appConfig.EmailMessageAckTimeOut //单位毫秒

	// rabbitmq创建一个EmailNotify相关的exchange和queue
	rabbitMq.CreateExchange(appConfig.EmailMessageExchangeName, "topic").
		CreateQueue(appConfig.EmailMessageQueueName, true, argsQue).
		ExchangeBindQueue(appConfig.EmailMessageQueueName, "*.email.*", appConfig.EmailMessageExchangeName)

}
