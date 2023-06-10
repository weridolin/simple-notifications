package types

import "github.com/weridolin/simple-vedio-notifications/servers/consumers/models"

func (s EmailNotifier) FromEmailNotifierModel(m models.EmailNotifier) *EmailNotifier {
	return &EmailNotifier{
		Sender:   m.Sender,
		PWD:      m.PWD,
		Receiver: m.Receiver,
		Content:  m.Content,
	}

}

func (s EmailNotifier) FromEmailNotifierModels(m []*models.EmailNotifier) []*EmailNotifier {
	var res []*EmailNotifier
	for _, v := range m {
		res = append(res, s.FromEmailNotifierModel(*v))
	}
	return res
}
