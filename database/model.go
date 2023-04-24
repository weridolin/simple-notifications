package database

import (
	"gorm.io/gorm"
)

// gorm.Model 的定义
// type Model struct {
// 	ID        uint           `gorm:"primaryKey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
//   }

type User struct {
	gorm.Model
	Username     string `gorm:"uniqueIndex;not null;comment:用户名;size:256"`
	Password     string `gorm:"not null;comment:密码"`
	Email        string `gorm:"comment:邮箱"`
	Phone        string `gorm:"comment:手机号"`
	Avatar       string `gorm:"comment:头像连接"`
	Role         int    `gorm:"comment:角色"`
	IsSuperAdmin bool   `gorm:"default:false"`
	Deleted      bool   `gorm:"default:false"`
	Age          int    `gorm:"comment:年龄"`
	Sex          int8   `gorm:"comment:性别"`
}

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

type EmailNotifierTask struct {
	gorm.Model
	EmailNotifierID uint
	EmailNotifier   EmailNotifier `gorm:"foreignKey:EmailNotifierID;OnDelete:CASCADE"`
	TaskID          uint          `gorm:"comment:任务ID"`
	Task            Task          `gorm:"foreignKey:TaskID;OnDelete:CASCADE"`
}
