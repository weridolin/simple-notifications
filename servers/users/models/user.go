package models

import (
	"errors"

	"github.com/lib/pq"
	"github.com/weridolin/simple-vedio-notifications/tools"
	"gorm.io/gorm"
)

type UserModel interface {
	Create(username, email, password string, DB *gorm.DB) (*User, error)
	Delete(id int, DB *gorm.DB) error
	QueryUser(condition interface{}, DB *gorm.DB) (User, error)
	Update(u User, DB *gorm.DB) error
	CheckPWd(count, password string, DB *gorm.DB) (*User, error)
}

type User struct {
	gorm.Model
	Username     string         `gorm:"uniqueIndex;not null;comment:用户名;size:256" json:"username" binding:"alphanum,min=4,max=255" form:"username"`
	Password     string         `gorm:"not null;comment:密码" json:"password" binding:"required,min=4,max=255" form:"password"`
	Email        string         `gorm:"comment:邮箱" json:"email" binding:"email" form:"email"`
	Phone        string         `gorm:"comment:手机号" json:"phone" form:"phone"`
	Avatar       string         `gorm:"comment:头像连接" json:"avatar" form:"avatar"`
	Role         pq.StringArray `gorm:"comment:角色;type:json" json:"role" form:"role" `
	IsSuperAdmin bool           `gorm:"default:false" json:"is_super_admin" binding:"-"`
	Deleted      bool           `gorm:"default:false" json:"-" binding:"-"`
	Age          int            `gorm:"comment:年龄" json:"age"  form:"age"`
	Gender       int8           `gorm:"comment:性别" json:"gender" form:"gender"`
}

type DefaultUserModel struct {
	Table string `gorm:"-" json:"-" binding:"-"`
}

func NewUserModel(table string) UserModel {
	return DefaultUserModel{
		Table: table,
	}
}

func (u DefaultUserModel) Create(username, email, password string, DB *gorm.DB) (*User, error) {
	user := User{
		Username: username,
		Email:    email,
		Password: tools.GetMD5Hash(password),
	}
	DB.Table(u.Table).Where("username = ? or email = ? ", username, email).First(&user)
	if user.ID != 0 {
		return nil, errors.New("用户名或邮箱已存在")
	} else {
		DB.Create(&user)
		return &user, nil
	}
}

func (m DefaultUserModel) QueryUser(condition interface{}, DB *gorm.DB) (User, error) {
	var user User
	err := DB.Table(m.Table).Where(condition).First(&user).Error
	return user, err
}

func (m DefaultUserModel) Delete(id int, DB *gorm.DB) error {
	err := DB.Table(m.Table).Delete(map[string]interface{}{"id": id}).Error
	return err
}

func (m DefaultUserModel) Update(u User, DB *gorm.DB) error {
	err := DB.Table(m.Table).Updates(u).Error
	return err
}

func (m DefaultUserModel) CheckPWd(count, password string, DB *gorm.DB) (*User, error) {
	var user User
	var err error
	if tools.IsEmail(count) {
		user, err = m.QueryUser(&User{Email: count}, DB)
		if err != nil {
			return nil, errors.New("邮箱不存在")
		}
	} else {
		user, err = m.QueryUser(&User{Username: count}, DB)
		if err != nil {
			return nil, errors.New("用户名不存在")
		}
	}
	if user.Password != tools.GetMD5Hash(password) {
		return nil, errors.New("密码错误")
	}

	return &user, nil
}
