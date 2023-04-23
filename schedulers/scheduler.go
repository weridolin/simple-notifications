package schedulers

import (
	"fmt"

	bilibili "github.com/weridolin/simple-vedio-notifications/platforms/bilibili"
	tools "github.com/weridolin/simple-vedio-notifications/tools"
)

type Scheduler struct {
	Period   tools.Period
	PlatForm string
	Ups      map[string]interface{}
	Status   int8 // 0 停止 1 启动  2 暂停
	DBIndex  int  //唯一索引
}

func NewScheduler(period tools.Period, platform string, ups map[string]interface{}, status int8, dbindex int) *Scheduler {
	return &Scheduler{period, platform, ups, status, dbindex}
}

func (s *Scheduler) Start() {
	s.Status = 1
	var t Task
	switch s.PlatForm {
	case "bilibili":
		t = &bilibili.BiliBiliTask{Period: s.Period, Ups: s.Ups}
	case "youtube":
		fmt.Println("start youtube scheduler...", s)
	}
	t.Run()
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
