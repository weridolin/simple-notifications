/*
	每个scheduler表示一个定时任务,可以包括多个task.
	scheduler 不支持用户自定义，只提供内置的scheduler

*/
package schedulers

import (
	"github.com/weridolin/simple-vedio-notifications/common"
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
	logger.Println("start scheduler...task ->", s.Tasks)
	s.Status = 1
	var i common.ITask
	for _, t := range s.Tasks {
		i = t.(common.ITask)
		i.Run()
	}
}

func (s *Scheduler) Stop() {
	s.Status = 0
	var i common.ITask
	for _, t := range s.Tasks {
		i = t.(common.ITask)
		i.Stop()
	}
	logger.Println("stop scheduler finish", "DBIndex -> ", s.DBIndex)
}

func (s *Scheduler) AddTask(t interface{}) {
	for _, task := range s.Tasks {
		if task.(common.Meta).DBIndex == t.(common.Meta).DBIndex {
			logger.Println("task already exist in scheduler,period -> ", s.Period.Cron, " platform -> ",
				s.PlatForm, " taskID -> ", t.(common.Meta).DBIndex, "ScheduleID", s.DBIndex)
			return
		}
	}
	s.Tasks = append(s.Tasks, t)
}

func (s *Scheduler) RemoveTask(t common.Meta) {
	logger.Println("remove task ID -> ", t.DBIndex, " from scheduler ID -> ", s.DBIndex)
	for i, task := range s.Tasks {
		if task.(common.Meta).DBIndex == t.DBIndex {
			s.Tasks = append(s.Tasks[:i], s.Tasks[i+1:]...)
			// task.(Task).Stop() #TODO
		}
	}
}

func (s *Scheduler) Delete() {
	s.Status = 0
	s.ticker.Stop()
	logger.Println("delete scheduler...", s)
}
