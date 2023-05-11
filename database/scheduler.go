package database

import (
	"github.com/weridolin/simple-vedio-notifications/tools"
	"gorm.io/gorm"
)

type Scheduler struct {
	User User `gorm:"foreignKey:UserID;OnDelete:CASCADE;AssociationForeignKey:ID"` // 外键约束
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

func (Scheduler) TableName() string {
	return "scheduler"
}

func CreateScheduler(user User, period string, platform, name, description string) error {
	scheduler := Scheduler{
		UserID:      user.ID,
		Period:      period,
		Type:        "custom",
		Platform:    platform,
		Name:        name,
		Description: description,
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
	// print()
	return scheduler, err
}

func FullQuerySchedulers(condition interface{}, page, size int) ([]*Scheduler, error) {
	db := GetDB()
	var schedulers []*Scheduler
	err := db.Model(&Scheduler{}).Preload("Tasks").Preload("Tasks.EmailNotifiers").Where(condition).Offset((page - 1) * size).Limit(size).Find(&schedulers).Error
	return schedulers, err
}

func BindTask(schedulerId uint, taskIds []uint) error {
	db := GetDB()
	var st []SchedulerTask
	for _, taskId := range taskIds {
		st = append(st, SchedulerTask{
			SchedulerID: schedulerId,
			TaskID:      taskId,
		})
	}
	err := db.Create(&st).Error
	return err
}
