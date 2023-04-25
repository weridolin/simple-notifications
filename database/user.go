package database

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/weridolin/simple-vedio-notifications/tools"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"uniqueIndex;not null;comment:用户名;size:256" json:"username" binding:"alphanum,min=4,max=255" form:"username"`
	Password     string `gorm:"not null;comment:密码" json:"password" binding:"required,min=4,max=255" form:"password"`
	Email        string `gorm:"comment:邮箱" json:"email" binding:"email" form:"email"`
	Phone        string `gorm:"comment:手机号" json:"phone" form:"phone"`
	Avatar       string `gorm:"comment:头像连接" json:"avatar" form:"avatar"`
	Role         int    `gorm:"comment:角色" json:"role" form:"role"`
	IsSuperAdmin bool   `gorm:"default:false" json:"is_super_admin" binding:"-"`
	Deleted      bool   `gorm:"default:false" json:"-" binding:"-"`
	Age          int    `gorm:"comment:年龄" json:"age"  form:"age"`
	Gender       int8   `gorm:"comment:性别" json:"gender" form:"gender"`
}

func (u *User) Create() (*User, error) {
	DB.Where("username = ? and email = ? ", u.Username, u.Email).First(&u)
	if u.ID != 0 {
		return nil, errors.New("用户名或邮箱已存在")
	} else {
		DB.Create(&u)
		return u, nil
	}
}

func NewUserValidator() User {
	return User{}
}

func (u *User) Bind(c *gin.Context) error {
	err := tools.Bind(c, u)
	if err != nil {
		return err
	}
	// 处理下密码
	u.Password = tools.GetMD5Hash(u.Password)
	return nil
}

func QueryFirst(condition interface{}) (User, error) {
	db := GetDB()
	var user User
	err := db.Where(condition).First(&user).Error
	return user, err
}

func (u *User) Delete() error {
	DB.Where("id = ?", u.ID).Delete(&u)
	return nil
}

func (u *User) Update() error {
	DB.Model(&u).Updates(u)
	return nil
}
