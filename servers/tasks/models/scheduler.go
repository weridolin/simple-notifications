package models

import (
	user "github.com/weridolin/simple-vedio-notifications/servers/users/models"
	"github.com/weridolin/simple-vedio-notifications/tools"
	"gorm.io/gorm"
)

type (
	SchedulerModel interface {
		Create(user user.User, period string, platform, name, description string, DB *gorm.DB) error
		Update(data interface{}, DB *gorm.DB) error
		Delete(id int, DB *gorm.DB) error
		Query(condition interface{}, DB *gorm.DB) (Scheduler, error)
		BindTask(schedulerId uint, taskIds []uint, DB *gorm.DB) error
	}

	Scheduler struct {
		User user.User `gorm:"foreignKey:UserID;OnDelete:CASCADE;AssociationForeignKey:ID"` // 外键约束
		gorm.Model
		Period      string  `gorm:"comment:时间"`
		UserID      uint    `gorm:"comment:用户ID"`
		Tasks       []*Task `gorm:"many2many:scheduler_task"`   // scheduler和task是多对多的关系
		Type        string  `gorm:"comment:类型;default:builtin"` // builtin:内置, custom:自定义
		Active      bool    `gorm:"default:true"`
		Platform    string  `gorm:"comment:平台"`
		Name        string  `gorm:"comment:计划名称"`
		Description string  `gorm:"comment:计划描述"`
	}

	DefaultSchedulerModel struct {
		Table string `gorm:"-" json:"-" binding:"-"`
	}
)

func (Scheduler) TableName() string {
	return "scheduler"
}

func NewSchedulerModel(table string) SchedulerModel {
	return &DefaultSchedulerModel{
		Table: table,
	}
}

func (s *DefaultSchedulerModel) Create(user user.User, period string, platform, name, description string, DB *gorm.DB) error {
	scheduler := Scheduler{
		UserID:      user.ID,
		Period:      period,
		Type:        "custom",
		Platform:    platform,
		Name:        name,
		Description: description,
	}
	_, e := s.Query(scheduler, DB)
	if e != nil {
		err := DB.Table(s.Table).Create(&scheduler).Error
		return err
	}
	return tools.SchedulerIsExistError
}

func (s *DefaultSchedulerModel) Update(data interface{}, DB *gorm.DB) error {
	err := DB.Table(s.Table).Updates(data).Error
	return err
}

func (s *DefaultSchedulerModel) Delete(id int, DB *gorm.DB) error {
	err := DB.Table(s.Table).Delete(map[string]interface{}{"id": id}).Error
	return err
}

func (s *DefaultSchedulerModel) Query(condition interface{}, DB *gorm.DB) (Scheduler, error) {
	var scheduler Scheduler
	err := DB.Table(s.Table).Where(condition).First(&scheduler).Error
	return scheduler, err
}

func (s *DefaultSchedulerModel) FullQuery(condition interface{}, page, size int, DB *gorm.DB) ([]*Scheduler, error) {
	var schedulers []*Scheduler
	err := DB.Table(s.Table).Preload("Tasks").Preload("Tasks.EmailNotifiers").Where(condition).Offset((page - 1) * size).Limit(size).Find(&schedulers).Error
	return schedulers, err
}

func (s *DefaultSchedulerModel) BindTask(schedulerId uint, taskIds []uint, DB *gorm.DB) error {
	var st []SchedulerTask
	for _, taskId := range taskIds {
		st = append(st, SchedulerTask{
			SchedulerID: schedulerId,
			TaskID:      taskId,
		})
	}
	err := DB.Create(&st).Error
	return err
}
