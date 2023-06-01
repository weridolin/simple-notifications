package models

import "gorm.io/gorm"

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
