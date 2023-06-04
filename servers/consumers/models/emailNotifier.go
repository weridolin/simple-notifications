package models

import "gorm.io/gorm"

type EmailNotifierModel interface {
	Delete(id int, DB *gorm.DB) error
	Update(data map[string]interface{}, DB *gorm.DB) error
	Create(userId int, pwd, sender, content string, receiver []string, DB *gorm.DB) (EmailNotifier, error)
	Query(condition interface{}, DB *gorm.DB) (EmailNotifier, error)
	QueryAll(condition interface{}, page, size int, DB *gorm.DB) ([]*EmailNotifier, error)
}

type EmailNotifier struct {
	gorm.Model
	UserID   int
	Sender   string        `gorm:"comment:邮箱;not null"`
	PWD      string        `gorm:"comment:邮箱授权码;not null"`
	Receiver EmailReceiver `gorm:"comment:接收者;TYPE:json"`
	Content  string        `gorm:"comment:内容"`
}

type DefaultEmailNotifierModel struct {
	Table string
}

func (EmailNotifier) TableName() string {
	return "email_notifier"
}

func NewDefaultEmailNotifierModel(table string) EmailNotifierModel {
	return &DefaultEmailNotifierModel{
		Table: table,
	}
}

type EmailNotifierTask struct {
	gorm.Model
	EmailNotifierID int `gorm:"comment:邮件通知ID"`
	// EmailNotifier   EmailNotifier `gorm:"foreignKey:EmailNotifierID;OnDelete:CASCADE"`
	TaskID int `gorm:"comment:任务ID"`
}

func (EmailNotifierTask) TableName() string {
	return "email_notifier_tasks"
}

func (m *DefaultEmailNotifierModel) Delete(id int, DB *gorm.DB) error {
	err := DB.Table(m.Table).Delete(m.Table, map[string]interface{}{"id": id}).Error //硬删除，如果包括DeletedAt,软删除
	return err
}

func (m *DefaultEmailNotifierModel) Update(data map[string]interface{}, DB *gorm.DB) error {
	err := DB.Table(m.Table).Updates(data).Error
	return err
}

func (m *DefaultEmailNotifierModel) Create(userId int, pwd, sender, content string, receiver []string, DB *gorm.DB) (EmailNotifier, error) {
	new := EmailNotifier{
		UserID:   userId,
		Sender:   sender,
		PWD:      pwd,
		Receiver: receiver,
		Content:  content,
	}
	err := DB.Table(m.Table).Create(&new).Error
	return new, err
}

func (m *DefaultEmailNotifierModel) Query(condition interface{}, DB *gorm.DB) (EmailNotifier, error) {
	var emailNotifier EmailNotifier
	err := DB.Table(m.Table).Where(condition).First(&emailNotifier).Error
	return emailNotifier, err
}

func (m *DefaultEmailNotifierModel) QueryAll(condition interface{}, page, size int, DB *gorm.DB) ([]*EmailNotifier, error) {
	var notifiers []*EmailNotifier
	err := DB.Table(m.Table).Where(condition).Offset((page - 1) * size).Limit(size).Find(&notifiers).Error
	return notifiers, err
}

func BindEmailNotifierToTask(userId, taskId int, emailNotifierId []int, DB *gorm.DB) error {
	var emailNotifierTasks []EmailNotifierTask
	for _, id := range emailNotifierId {
		emailNotifierTasks = append(emailNotifierTasks, EmailNotifierTask{
			EmailNotifierID: id,
			TaskID:          taskId,
		})
	}
	err := DB.Create(&emailNotifierTasks).Error
	return err
}
