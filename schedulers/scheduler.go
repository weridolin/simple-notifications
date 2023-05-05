/*
	每个scheduler表示一个定时任务,可以包括多个task
*/
package schedulers

import (
	"fmt"

	tools "github.com/weridolin/simple-vedio-notifications/tools"
)

type Scheduler struct {
	Period   tools.Period  //定时周期
	PlatForm string        //	scheduler对应的平台名称
	Tasks    []interface{} // scheduler对应的task列表
	Status   int8          // 0 停止 1 启动  2 暂停
	DBIndex  int           //唯一索引
	ticker   *Ticker       //绑定的ticker
}

func NewScheduler(period tools.Period, platform string, status int8, dbindex int) *Scheduler {
	return &Scheduler{
		period,
		platform,
		make([]interface{}, 0),
		status,
		dbindex,
		nil,
	}
}

func (s *Scheduler) Start() {
	s.Status = 1
	var i ITask
	for _, t := range s.Tasks {
		i = t.(ITask)
		i.Run()
	}
	// var t Task
	// switch s.PlatForm {
	// case "bilibili":
	// 	t = &bilibili.BiliBiliTask{Period: s.Period, Ups: s.Ups}
	// case "youtube":
	// 	fmt.Println("start youtube scheduler...", s)
	// }
	// t.Run()
}

func (s *Scheduler) Stop() {
	s.Status = 0
	var i ITask
	for _, t := range s.Tasks {
		i = t.(ITask)
		i.Stop()
	}
	fmt.Println("stop scheduler...", s)
}

func (s *Scheduler) AddTask(t interface{}) {
	for _, task := range s.Tasks {
		if task.(Task).DBIndex == t.(Task).DBIndex {
			return
		}
	}
	s.Tasks = append(s.Tasks, t)
	fmt.Println("add task...", s)
}

func (s *Scheduler) RemoveTask(t Task) {
	fmt.Println("remove up...", s)
	for i, task := range s.Tasks {
		if task.(Task).DBIndex == t.DBIndex {
			s.Tasks = append(s.Tasks[:i], s.Tasks[i+1:]...)
		}
	}
}

func (s *Scheduler) Delete() {
	s.Status = 0
	fmt.Println("delete scheduler...", s)
}
