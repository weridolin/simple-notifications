package common

import "errors"

type Meta struct {
	DBIndex     uint
	Name        string
	Description string
	CallBacks   []func() //每次运行回调函数列表
}

func (t *Meta) Run() error {
	return errors.New("not implement")
}

func (t *Meta) Stop() error {
	return errors.New("not implement")
}

type ITask interface {
	Run()
	GetUpInfo()
	Stop()
}