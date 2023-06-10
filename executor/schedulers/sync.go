/*
从数据库中同步任务.触发方式:

	1.达到最大时间间隔
	2.有新的任务插入,信号触发
*/
package scheduler

import (
	"context"
	"time"

	"github.com/weridolin/simple-vedio-notifications/clients"
	"github.com/weridolin/simple-vedio-notifications/executor/platforms/bilibili"
	"github.com/weridolin/simple-vedio-notifications/tools"
	"github.com/zeromicro/go-zero/core/logx"
)

// var config = config.Getconfig()

type Synchronizer struct {
	MaxInterval      int
	MsgChannel       chan string
	StopSleepChannel chan bool
	SchedulerManager *SchedulerManager
	Ctx              context.Context
	AppConfig        SchedulerConfig
}

func NewSynchronizer(config SchedulerConfig) *Synchronizer {
	ctx := context.WithValue(context.Background(), "tp", NewTickerPool(config))
	uuid := tools.GetUUID()
	manager := NewSchedulerManager(ctx, uuid)
	return &Synchronizer{
		MaxInterval:      3000,
		SchedulerManager: manager,
		Ctx:              ctx,
		AppConfig:        config,
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
	s.Setup(s.AppConfig)
	s.MsgChannel = make(chan string)
	s.StopSleepChannel = make(chan bool)
	for {
		select { // 从通道中读取数据
		case msg := <-s.MsgChannel:
			if msg == "stop" {
				logx.Info("close synchronizer", msg)
				close(s.StopSleepChannel)
				close(s.MsgChannel)
				return
			}
		default:
			logx.Info("LOAD FROM DB")
			// 从数据库中加载数据,并同步到内存中
			records := s.LoadTaskFromDB()
			for _, record := range records {
				logx.Info("sync scheduler", "platform -> ", record.(Scheduler).PlatForm, "task ->", record.(Scheduler).Tasks, "Period ->", record.(Scheduler).Period)
				scheduler := NewScheduler(record.(Scheduler).Period, record.(Scheduler).PlatForm, 0, int(record.(Scheduler).DBIndex))
				for _, model := range record.(Scheduler).Tasks {
					var task interface{}
					switch model.(Task).PlatForm {
					case "bilibili":
						task = bilibili.NewBiliBiliTask(tools.Period{Cron: record.(Scheduler).Period}, model.(Task).Ups, model.(Task).DBIndex, model.(Task).Name, model.(Task).Description,
							s.Ctx.Value("tp").(*TickerPool).Storage)
					}
					if task != nil {
						scheduler.Tasks = append(scheduler.Tasks, task)
					}
				}
				_, err := s.SchedulerManager.AddScheduler(scheduler, true)
				if err != nil {
					logx.Info("add scheduler error", err)
				}
			}
			s.Sleep(s.MaxInterval)
		}

	}
}

func (s *Synchronizer) LoadTaskFromDB() []interface{} {
	// var schedulers  []Scheduler
	// 1.从数据库中加载所有的scheduler
	// TODO 通过RPC
	// schedulers, err := database.FullQuerySchedulers(&database.Scheduler{Active: true}, 1, 100)
	// if err != nil {
	// 	logger.Panicln("load scheduler from db error: ", err)
	// }
	// logger.Println("load scheduler finish", schedulers)
	// return schedulers
	return make([]interface{}, 1)
}

func (s *Synchronizer) Stop() {
	s.StopSleepChannel <- true
	s.MsgChannel <- "stop"
}

func (s *Synchronizer) SyncAtOnce() {
	s.StopSleepChannel <- true
}

func (s *Synchronizer) Setup(config SchedulerConfig) {
	rabbitMq := clients.NewRabbitMQ(tools.GetUUID(), config.RabbitMQUri)
	//创建1个email notify 死信交换机和队列config
	rabbitMq.CreateExchange(config.EmailMessageDlxExchangeName, "direct").
		CreateQueue(config.EmailMessageDlxQueueName, true, nil).
		ExchangeBindQueue(config.EmailMessageDlxQueueName, config.EmailMessageDlxQueueName, config.EmailMessageDlxExchangeName)

	// 配置队列参数
	var dlxExchangeName = config.EmailMessageDlxExchangeName
	argsQue := make(map[string]interface{})
	//添加死信队列交换器属性
	argsQue["x-dead-letter-exchange"] = dlxExchangeName
	//指定死信队列的路由key，不指定使用队列路由键
	argsQue["x-dead-letter-routing-key"] = config.EmailMessageDlxQueueName
	//添加过期时间
	argsQue["x-message-ttl"] = config.EmailMessageAckTimeOut //单位毫秒

	// rabbitmq创建一个EmailNotify相关的exchange和queue
	rabbitMq.CreateExchange(config.EmailMessageExchangeName, "topic").
		CreateQueue(config.EmailMessageQueueName, true, argsQue).
		ExchangeBindQueue(config.EmailMessageQueueName, "*.email.*", config.EmailMessageExchangeName)

}
