package main

import (
	"context"
	"fmt"

	schedulers "github.com/weridolin/simple-vedio-notifications/schedulers"
	tools "github.com/weridolin/simple-vedio-notifications/tools"
)

func main() {
	ctx := context.Background()
	uuid := tools.GetUUID()
	manager := schedulers.NewSchedulerManager(ctx, uuid)
	scheduler := schedulers.NewScheduler(schedulers.Period{Cron: "@hourly"}, "bilibili", []string{"敬汉卿"}, 0, 1)
	manager.AddScheduler(scheduler)
	manager.StartAll()
	fmt.Println(manager.Schedulers, manager.PlatFormSchedulerCache)

}
