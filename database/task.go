package database

import "gorm.io/gorm"

type Task struct {
	Scheduler []*Scheduler `gorm:"many2many:scheduler_task;"` // scheduler和task是多对多的关系
	gorm.Model
	PlatForm      string           `gorm:"comment:平台"`
	Ups           Ups              `gorm:"comment:订阅的该平台的up主;type:json"` // 用json存储map
	EmailNotifier []*EmailNotifier `gorm:"many2many:email_notifier_tasks;"`
	User          User             `gorm:"foreignKey:UserID;OnDelete:CASCADE;AssociationForeignKey:ID"` // 外键约束
	UserID        uint             `gorm:"comment:用户ID"`
	Active        bool             `gorm:"default:true"`
	Name          string           `gorm:"comment:任务名称"`
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

func (t *Task) Delete() error {
	db := GetDB()
	err := db.Delete(&t).Error
	return err
}

func (t *Task) Update() error {
	db := GetDB()
	err := db.Model(&t).Updates(t).Error
	return err
}

func CreateTask(user User, ups Ups, platform, name string) error {
	db := GetDB()
	new := Task{
		PlatForm: platform,
		Ups:      ups,
		UserID:   user.ID,
		Name:     name,
	}
	err := db.Create(&new).Error
	return err
}

func GetTaskById(id int) (Task, error) {
	db := GetDB()
	var task Task
	err := db.Where("id = ?", id).First(&task).Error
	return task, err
}

func SearchTasks(user User, name string, page, size int, schedulerIds int) ([]Task, error) {
	db := GetDB()
	condition := map[string]interface{}{
		"user_id": user.ID, "active": true,
	}
	var task []Task
	if name != "" {
		condition["name"] = name
	}
	if schedulerIds != -1 {
		db.Model(&user).Where(condition).Association("Scheduler").Find(&task)
	}
	err := db.Where(condition).Find(&task).Error
	return task, err
}
