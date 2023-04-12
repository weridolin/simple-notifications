package schedulers

import (
	"fmt"
	// tools "github.com/weridolin/simple-vedio-notifications/tools"
	bilibili "github.com/weridolin/simple-vedio-notifications/platforms/bilibili"
)

type Period struct {
	Cron string
}

type Scheduler struct {
	Period   Period
	PlatForm string
	Ups      []string
	Status   int8 // 0 停止 1 启动  2 暂停
	DBIndex  int  //唯一索引
}

func NewScheduler(period Period, platform string, ups []string, status int8, dbindex int) *Scheduler {
	return &Scheduler{period, platform, ups, status, dbindex}
}

func (s *Scheduler) Start() {
	s.Status = 1
	var t Task
	switch s.PlatForm {
	case "bilibili":
		t = &bilibili.BiliBiliTask{"央视网", "2020-04-01", 9824766}
	case "youtube":
		fmt.Println("start youtube scheduler...", s)
	}
	t.Run()
	fmt.Println("start scheduler...", s)
}

func (s *Scheduler) Stop() {
	s.Status = 0
	fmt.Println("stop scheduler...", s)
}

func (s *Scheduler) AddUp() {
	fmt.Println("add up...", s)
}

func (s *Scheduler) RemoveUp() {
	fmt.Println("remove up...", s)
}

func (s *Scheduler) Delete() {
	s.Status = 0
	fmt.Println("delete scheduler...", s)
}
