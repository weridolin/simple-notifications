package schedulers

import (
	"fmt"

	"github.com/weridolin/simple-vedio-notifications/database"
	"github.com/weridolin/simple-vedio-notifications/servers/common"
	"github.com/weridolin/simple-vedio-notifications/servers/http/users"
)

type SchedulerSerializer struct {
	common.BaseSerializer
	ID     uint                         `json:"id"`
	Period string                       `json:"period"`
	Active bool                         `json:"active"`
	User   users.UserResponseSerializer `json:"user"`
	Tasks  []TaskSerializer             `json:"tasks"`
}

func (u SchedulerSerializer) FromSchedulerModel(m *database.Scheduler, user database.User) *SchedulerSerializer {
	return &SchedulerSerializer{
		ID:     m.ID,
		Period: m.Period,
		Active: m.Active,
		User:   users.UserResponseSerializer{}.FromUserModel(&user),
		// Roles:    m.Roles,
		Tasks: TaskSerializer{}.FromTaskModels(m.Tasks, user),
	}

}

func (u SchedulerSerializer) FromSchedulerModels(m []*database.Scheduler, user database.User) []*SchedulerSerializer {
	var res []*SchedulerSerializer
	for _, v := range m {
		s := u.FromSchedulerModel(v, user)
		res = append(res, s)
	}
	return res
}

type TaskSerializer struct {
	ID   uint                         `json:"id"`
	User users.UserResponseSerializer `json:"user"`
	common.BaseSerializer
	PlatForm   string                 `json:"platform"`
	Ups        database.Ups           `json:"ups"`
	Schedulers []*SchedulerSerializer `json:"schedulers"`
	// EmailNotifier []*EmailNotifier `json:"email_notifiers"`
}

func (u TaskSerializer) FromTaskModel(m *database.Task, user database.User) TaskSerializer {
	fmt.Println(">>>", m.Schedulers)
	res := TaskSerializer{
		ID:       m.ID,
		PlatForm: m.PlatForm,
		Ups:      database.Ups(m.Ups),
		// EmailNotifier: v.EmailNotifier,
		User:       users.UserResponseSerializer{}.FromUserModel(&user),
		Schedulers: SchedulerSerializer{}.FromSchedulerModels(m.Schedulers, user),
	}
	return res
}

func (u TaskSerializer) FromTaskModels(m []*database.Task, user database.User) []TaskSerializer {
	var res []TaskSerializer
	for _, v := range m {
		fmt.Println(v.Schedulers)
		s := u.FromTaskModel(v, user)
		res = append(res, s)
	}
	return res
}
