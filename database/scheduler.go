package database

import (
	"gorm.io/gorm"
)

type Scheduler struct {
	UserID uint
	User   User `gorm:"foreignKey:UserID;OnDelete:CASCADE;AssociationForeignKey:ID"` // 外键约束
	gorm.Model
	Period string `gorm:"comment:时间"`

	Task    []Task
	Deleted bool `gorm:"default:false"`
	Active  bool `gorm:"default:true"`
}

type Ups map[string]interface{}

type Task struct {
	SchedulerID uint
	Scheduler   Scheduler `gorm:"foreignKey:SchedulerID;OnDelete:CASCADE;AssociationForeignKey:ID"` // 外键约束
	gorm.Model
	PlatForm      string           `gorm:"comment:平台"`
	Ups           Ups              `gorm:"comment:订阅的该平台的up主;type:json"` // 用json存储map
	EmailNotifier []*EmailNotifier `gorm:"many2many:email_notifier_tasks;"`
}

func (t *Task) Create() error {
	DB.Create(&t)
	return nil
}

func (t *Task) Delete() error {
	DB.Delete(&t)
	return nil
}

func (t *Task) Update() error {
	DB.Model(&t).Updates(t)
	return nil
}

// func (t *Task) Query() (*Task, error) {
// 	var tasks []User
// 	DB.Where(t).Find(&tasks)
// 	if t.ID == 0 {
// 		return nil, errors.New("任务不存在")
// 	} else {
// 		return t, nil
// 	}
// }
