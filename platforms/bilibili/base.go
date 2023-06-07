package bilibili

import (
	"errors"

	"github.com/weridolin/simple-vedio-notifications/storage"
)

type Meta struct {
	DBIndex     int
	Name        string
	Description string
	CallBacks   []func()                 //每次运行回调函数列表
	Storage     storage.StorageInterface //结果储存
}

func (t *Meta) Run() error {
	return errors.New("not implement")
}

func (t *Meta) Stop() error {
	return errors.New("not implement")
}
