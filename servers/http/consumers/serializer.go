package consumers

import (
	"github.com/weridolin/simple-vedio-notifications/database"
	"github.com/weridolin/simple-vedio-notifications/servers/common"
	"github.com/weridolin/simple-vedio-notifications/servers/http/schedulers"
)

type EmailNotifierSerializer struct {
	common.BaseSerializer
	ID       uint                        `json:"id"`
	Tasks    []schedulers.TaskSerializer `JSON:"tasks"`
	Sender   string                      `json:"sender"`
	PWD      string                      `json:"pwd"`
	Receiver []string                    `json	:"receiver"`
	Content  string                      `json "content"`
}

func (u EmailNotifierSerializer) FromSchedulerModel(m *database.EmailNotifier, user database.User) *EmailNotifierSerializer {
	return &EmailNotifierSerializer{
		ID:       m.ID,
		Sender:   m.Sender,
		PWD:      m.PWD,
		Receiver: m.Receiver,
		Content:  m.Content,
		// Roles:    m.Roles,
		Tasks: schedulers.TaskSerializer{}.FromTaskModels(m.Tasks, user),
	}

}

func (u EmailNotifierSerializer) FromSchedulerModels(m []*database.EmailNotifier, user database.User) []*EmailNotifierSerializer {
	var res []*EmailNotifierSerializer
	for _, v := range m {
		s := u.FromSchedulerModel(v, user)
		res = append(res, s)
	}
	return res
}
