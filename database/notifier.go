package database

import "gorm.io/gorm"

type EmailNotifier struct {
	gorm.Model
	UserID   uint
	User     User          `gorm:"foreignKey:UserID;OnDelete:CASCADE"`
	Tasks    []*Task       `gorm:"many2many:email_notifier_tasks;"`
	Sender   string        `gorm:"comment:邮箱;not null"`
	PWD      string        `gorm:"comment:邮箱授权码;not null"`
	Receiver EmailReceiver `gorm:"comment:接收者;TYPE:json"`
	Content  string        `gorm:"comment:内容"`
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

func (e *EmailNotifier) Delete() error {
	db := GetDB()
	err := db.Delete(&e).Error
	return err
}

func (e *EmailNotifier) Update(data map[string]interface{}) error {
	db := GetDB()
	err := db.Model(&e).Updates(data).Error
	return err
}

func CreateEmailNotifier(user User, pwd, sender, content string, receiver []string) error {
	db := GetDB()
	new := EmailNotifier{
		UserID:   user.ID,
		Sender:   sender,
		PWD:      pwd,
		Receiver: receiver,
		Content:  content,
	}
	err := db.Create(&new).Error
	return err
}

func GetEmailNotifierById(id int) (EmailNotifier, error) {
	db := GetDB()
	var emailNotifier EmailNotifier
	err := db.Where("id = ?", id).First(&emailNotifier).Error
	return emailNotifier, err
}

func FullQueryEmailNotify(condition interface{}, page, size int) ([]*EmailNotifier, error) {
	db := GetDB()
	var notifiers []*EmailNotifier
	err := db.Model(&EmailNotifier{}).Preload("Tasks").Where(condition).Offset((page - 1) * size).Limit(size).Find(&notifiers).Error
	return notifiers, err
}

func BindEmailNotifierToTask(user User, taskId uint, emailNotifierId []uint) error {
	db := GetDB()
	var emailNotifierTasks []EmailNotifierTask
	for _, id := range emailNotifierId {
		emailNotifierTasks = append(emailNotifierTasks, EmailNotifierTask{
			EmailNotifierID: id,
			TaskID:          taskId,
		})
	}
	err := db.Create(&emailNotifierTasks).Error
	return err
}
