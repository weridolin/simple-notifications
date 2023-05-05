package schedulers

import "errors"

type Task struct {
	DBIndex int
}

func (t *Task) Run() error {
	return errors.New("not implement")
}

func (t *Task) Stop() error {
	return errors.New("not implement")
}

type ITask interface {
	Run()
	GetUpInfo()
	Stop()
}
