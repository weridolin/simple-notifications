package main

import (
	"context"
	"fmt"

	// schedulers "github.com/weridolin/simple-vedio-notifications/schedulers"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/weridolin/simple-vedio-notifications/database"
	"github.com/weridolin/simple-vedio-notifications/platforms/bilibili"
	"github.com/weridolin/simple-vedio-notifications/schedulers"
	"github.com/weridolin/simple-vedio-notifications/tools"
)

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
	// Setup()
	// http.Start()
	// TestCron()
	ctx := context.WithValue(context.Background(), "tp", schedulers.NewTickerPool(1))
	uuid := tools.GetUUID()
	manager := schedulers.NewSchedulerManager(ctx, uuid)
	task := bilibili.NewBiliBiliTask(
		tools.Period{Cron: tools.Minutely},
		map[string]interface{}{"盗月社食遇记": 99157282},
		0,
	)
	scheduler := schedulers.NewScheduler(tools.Period{Cron: tools.Minutely}, "bilibili", 0, 1)
	scheduler.AddTask(task)
	manager.AddScheduler(scheduler)
	fmt.Println(manager.Schedulers, manager.PlatFormSchedulerCache)
	manager.StartAll()
	time.Sleep(time.Minute * 2)
}
