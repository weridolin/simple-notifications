package schedulers

import (
	"github.com/weridolin/simple-vedio-notifications/database"
	"github.com/weridolin/simple-vedio-notifications/servers/http/users"
)

type SchedulerSerializer struct {
	ID     uint                         `json:"id"`
	Period string                       `json:"period"`
	Active bool                         `json:"active"`
	User   users.UserResponseSerializer `json:"user"`
	Tasks  []TaskSerializer             `json:"tasks"`
}

func (u SchedulerSerializer) FromSchedulerModel(m []database.Scheduler, user database.User) []SchedulerSerializer {
	var res []SchedulerSerializer
	for _, v := range m {
		s := SchedulerSerializer{
			ID:     v.ID,
			Period: v.Period,
			Active: v.Active,
			User:   users.UserResponseSerializer{}.FromUserModel(&user),
			// Roles:    m.Roles,

		}
		res = append(res, s)
	}
	return res
}

type TaskSerializer struct {
}
