/*
每个scheduler表示一个定时任务,可以包括多个task.
scheduler 不支持用户自定义，只提供内置的scheduler
*/
package scheduler

import (
	common "github.com/weridolin/simple-vedio-notifications/executor/common"
	"github.com/zeromicro/go-zero/core/logx"
)

func NewScheduler(period string, platform string, status int8, dbindex int) *Scheduler {
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
	logx.Info("start scheduler...task ->", s.Tasks)
	s.Status = 1
	var i ITask
	for _, t := range s.Tasks {
		i = t.(ITask)
		i.Run()
	}
}

func (s *Scheduler) Stop() {
	s.Status = 0
	var i ITask
	for _, t := range s.Tasks {
		i = t.(ITask)
		i.Stop()
	}
	logx.Info("stop scheduler finish", "DBIndex -> ", s.DBIndex)
}

func (s *Scheduler) AddTask(t interface{}) {
	for _, task := range s.Tasks {
		if task.(common.Meta).DBIndex == t.(common.Meta).DBIndex {
			logx.Info("task already exist in scheduler,period -> ", s.Period, " platform -> ",
				s.PlatForm, " taskID -> ", t.(common.Meta).DBIndex, "ScheduleID", s.DBIndex)
			return
		}
	}
	s.Tasks = append(s.Tasks, t)
}

func (s *Scheduler) RemoveTask(t common.Meta) {
	logx.Info("remove task ID -> ", t.DBIndex, " from scheduler ID -> ", s.DBIndex)
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
	logx.Info("delete scheduler...", s)
}
