package database

import (
	"github.com/weridolin/simple-vedio-notifications/tools"
	"gorm.io/gorm"
)

type Scheduler struct {
	User User `gorm:"foreignKey:UserID;OnDelete:CASCADE;AssociationForeignKey:ID"` // 外键约束
	gorm.Model
	Period string  `gorm:"comment:时间"`
	UserID uint    `gorm:"comment:用户ID"`
	Task   []*Task `gorm:"many2many:scheduler_task;"`  // scheduler和task是多对多的关系
	Type   string  `gorm:"comment:类型;default:builtin"` // builtin:内置, custom:自定义
	Active bool    `gorm:"default:true"`
}

func (Scheduler) TableName() string {
	return "scheduler"
}

func CreateScheduler(user User, period string) error {
	scheduler := Scheduler{
		UserID: user.ID,
		Period: period,
		Type:   "custom",
	}
	_, e := QueryScheduler(scheduler)
	if e != nil {
		db := GetDB()
		err := db.Create(&scheduler).Error
		return err
	}
	return tools.SchedulerIsExistError
}

func (s *Scheduler) Update(data interface{}) error {
	db := GetDB()
	err := db.Model(s).Updates(data).Error
	return err
}

func (s *Scheduler) Delete() error {
	db := GetDB()
	err := db.Delete(s).Error
	return err
}

func QueryScheduler(condition interface{}) (Scheduler, error) {
	db := GetDB()
	var scheduler Scheduler
	err := db.Where(condition).First(&scheduler).Error
	return scheduler, err
}

func QuerySchedulers(condition interface{}, page, size int) ([]Scheduler, error) {
	db := GetDB()
	var scheduler []Scheduler
	err := db.Where(condition).Find(&scheduler).Error
	return scheduler, err
}
