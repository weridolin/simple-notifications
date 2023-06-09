package models

import (
	"fmt"

	user "github.com/weridolin/simple-vedio-notifications/servers/users/models"
	"gorm.io/gorm"
)

type (
	TaskModel interface {
		Create(userId int, ups Ups, platform, name, description string, DB *gorm.DB) (*Task, error)
		Update(data interface{}, DB *gorm.DB) error
		Delete(id int, DB *gorm.DB) error
		Query(condition interface{}, DB *gorm.DB) (*Task, error)
		FullQueryTasks(condition interface{}, page, size int, DB *gorm.DB) ([]*Task, error)
	}

	Task struct {
		Schedulers []*Scheduler `gorm:"many2many:scheduler_task"` // scheduler和task是多对多的关系
		gorm.Model
		Platform string `gorm:"comment:平台"`
		Ups      Ups    `gorm:"comment:订阅的该平台的up主;type:json"` // 用json存储map
		// EmailNotifiers []*EmailNotifier `gorm:"many2many:email_notifier_tasks;"`
		User        user.User `gorm:"foreignKey:UserID;OnDelete:CASCADE;AssociationForeignKey:ID"` // 外键约束
		UserID      int       `gorm:"comment:用户ID"`
		Active      bool      `gorm:"default:true"`
		Name        string    `gorm:"comment:任务名称"`
		Description string    `gorm:"comment:任务描述"`
	}

	DefaultTaskModel struct {
		Table string `gorm:"-" json:"-" binding:"-"`
	}
)

func (Task) TableName() string {
	return "task"
}

func NewTaskModel(table string) TaskModel {
	return &DefaultTaskModel{
		Table: table,
	}
}

func (t *DefaultTaskModel) Delete(id int, DB *gorm.DB) error {
	fmt.Println("delete id ", id)
	err := DB.Table(t.Table).Delete(t.Table, map[string]interface{}{"id": id}).Error //硬删除，如果包括DeletedAt,软删除
	return err
}

func (t *DefaultTaskModel) Update(data interface{}, DB *gorm.DB) error {
	err := DB.Table(t.Table).Updates(t).Error
	return err
}

func (t *DefaultTaskModel) Create(userId int, ups Ups, platform, name, description string, DB *gorm.DB) (*Task, error) {
	new := Task{
		Platform:    platform,
		Ups:         ups,
		UserID:      userId,
		Name:        name,
		Description: description,
	}
	err := DB.Table(t.Table).Create(&new).Error
	return &new, err
}

func (t *DefaultTaskModel) Query(condition interface{}, DB *gorm.DB) (*Task, error) {
	var task *Task
	err := DB.Table(t.Table).Where(condition).First(&task).Error
	return task, err
}

func (t *DefaultTaskModel) GetTaskById(id int, DB *gorm.DB) (*Task, error) {
	// db := GetDB()
	var task *Task
	err := DB.Table(t.Table).Where("id = ?", id).First(&task).Error
	return task, err
}

func (t *DefaultTaskModel) FullQueryTasks(condition interface{}, page, size int, DB *gorm.DB) ([]*Task, error) {
	// db := GetDB()
	var tasks []*Task
	err := DB.Debug().Table(t.Table).Preload("Schedulers").Where(condition).Offset((page - 1) * size).Limit(size).Find(&tasks).Error
	return tasks, err
}
