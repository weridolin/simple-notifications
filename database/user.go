package database

import (
	"errors"

	"gorm.io/gorm"
)

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

func (u *User) Register() (*User, error) {
	DB.Where("username = ? and email = ? ", u.Username, u.Email).First(&u)
	if u.ID != 0 {
		return nil, errors.New("用户名或邮箱已存在")
	} else {
		DB.Create(&u)
		return u, nil
	}
}

// func (u *User) QueryFirst() (User, error) {
// 	DB.Where("username = ?", u.Username).First(&u)
// 	if u.ID == 0 {
// 		return *u, errors.New("用户不存在")
// 	} else {
// 		return *u, nil
// 	}
// }

func (u *User) Delete() error {
	DB.Where("id = ?", u.ID).Delete(&u)
	return nil
}

func (u *User) Update() error {
	DB.Model(&u).Updates(u)
	return nil
}
