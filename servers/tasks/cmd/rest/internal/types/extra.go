package types

import (
	"github.com/weridolin/simple-vedio-notifications/servers/tasks/models"
)

func (t Task) FromTaskModel(m *models.Task) *Task {
	return &Task{
		ID:          int(m.ID),
		PlatForm:    m.Platform,
		Ups:         m.Ups,
		Name:        m.Name,
		Description: m.Description,
		User: UserInfo{
			Username:     m.User.Username,
			Email:        m.User.Email,
			Phone:        m.User.Phone,
			Avatar:       m.User.Avatar,
			Role:         m.User.Role,
			IsSuperAdmin: m.User.IsSuperAdmin,
			Age:          m.User.Age,
			Gender:       m.User.Gender,
		},
		Schedulers: Scheduler{}.FromSchedulerModels(m.Schedulers),
	}
}

func (t Task) FromTaskModels(m []*models.Task) []*Task {
	res := make([]*Task, len(m))
	for i, v := range m {
		res[i] = t.FromTaskModel(v)
	}
	return res
}

func (s Scheduler) FromSchedulerModel(m *models.Scheduler) *Scheduler {
	return &Scheduler{
		ID:          int(m.ID),
		Period:      m.Period,
		Active:      m.Active,
		Name:        m.Name,
		Description: m.Description,
		User: UserInfo{
			Username:     m.User.Username,
			Email:        m.User.Email,
			Phone:        m.User.Phone,
			Avatar:       m.User.Avatar,
			Role:         m.User.Role,
			IsSuperAdmin: m.User.IsSuperAdmin,
			Age:          m.User.Age,
			Gender:       m.User.Gender,
		},
		Tasks: Task{}.FromTaskModels(m.Tasks),
	}
}

func (s Scheduler) FromSchedulerModels(m []*models.Scheduler) []*Scheduler {
	res := make([]*Scheduler, len(m))
	for i, v := range m {
		res[i] = s.FromSchedulerModel(v)
	}
	return res
}
