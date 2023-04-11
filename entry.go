package main

import (
	"context"
	"fmt"

	schedulers "github.com/weridolin/simple-vedio-notifications/schedulers"
)

func main() {
	ctx := context.Background()
	manager := schedulers.NewSchedulerManager(ctx)
	scheduler := schedulers.NewScheduler(schedulers.Period{Cron: "@hourly"}, "bilibili", []string{"敬汉卿"}, 0, 1)
	manager.AddScheduler(scheduler)
	manager.StartAll()
	fmt.Println(manager.Schedulers)

}
