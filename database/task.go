package database

import "gorm.io/gorm"

type Ups map[string]interface{}

type Task struct {
	SchedulerID uint
	Scheduler   []*Scheduler `gorm:"many2many:scheduler_task;"` // scheduler和task是多对多的关系
	gorm.Model
	PlatForm      string           `gorm:"comment:平台"`
	Ups           Ups              `gorm:"comment:订阅的该平台的up主;type:json"` // 用json存储map
	EmailNotifier []*EmailNotifier `gorm:"many2many:email_notifier_tasks;"`
}

func (Task) TableName() string {
	return "task"
}

type SchedulerTask struct {
	gorm.Model
	SchedulerID uint
	Scheduler   Scheduler `gorm:"foreignKey:SchedulerID;OnDelete:CASCADE;AssociationForeignKey:ID"` // scheduler和task是多对多的关系
	TaskID      uint
	Task        Task `gorm:"foreignKey:TaskID;OnDelete:CASCADE;AssociationForeignKey:ID"` // scheduler和task是多对多的关系
}

func (SchedulerTask) TableName() string {
	return "scheduler_task"
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
