package users

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/weridolin/simple-vedio-notifications/database"
	"github.com/weridolin/simple-vedio-notifications/tools"
)

type LoginRequestValidator struct {
	Count    string `form:"count" json:"count" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (u *LoginRequestValidator) Bind(c *gin.Context) error {
	err := tools.Bind(c, u)
	if err != nil {
		return err
	}
	return nil
}

func (u *LoginRequestValidator) CheckPWd() (database.User, error) {
	var user database.User
	var err error
	if tools.IsEmail(u.Count) {
		user, err = database.QueryFirst(&database.User{Email: u.Count})
		if err != nil {
			return database.User{}, errors.New("用户不存在")
		}
	} else {
		user, err = database.QueryFirst(&database.User{Username: u.Count})
		if err != nil {
			return database.User{}, errors.New("用户不存在")
		}
	}
	if user.Password != tools.GetMD5Hash(u.Password) {
		return database.User{}, errors.New("密码错误")
	}

	// 生成JWT TOKEN

	return user, nil
}
