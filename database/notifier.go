package database

import "gorm.io/gorm"

type EmailNotifier struct {
	gorm.Model
	UserID   uint
	User     User    `gorm:"foreignKey:UserID;OnDelete:CASCADE"`
	Task     []*Task `gorm:"many2many:email_notifier_tasks;"`
	Sender   string  `gorm:"comment:邮箱;not null"`
	PWD      string  `gorm:"comment:邮箱授权码;not null"`
	Receiver string  `gorm:"comment:接收者"`
	Content  string  `gorm:"comment:内容"`
}

func (EmailNotifier) TableName() string {
	return "email_notifier"
}

type EmailNotifierTask struct {
	gorm.Model
	EmailNotifierID uint
	EmailNotifier   EmailNotifier `gorm:"foreignKey:EmailNotifierID;OnDelete:CASCADE"`
	TaskID          uint          `gorm:"comment:任务ID"`
	Task            Task          `gorm:"foreignKey:TaskID;OnDelete:CASCADE"`
}

func (EmailNotifierTask) TableName() string {
	return "email_notifier_tasks"
}
