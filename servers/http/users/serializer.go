package users

import (
	"github.com/weridolin/simple-vedio-notifications/database"
)

type UserResponseSerializer struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Email    string   `json:"email"`
	Phone    string   `json:"phone"`
	Avatar   string   `json:"avatar"`
	Age      int      `json:"age"`
	Gender   int8     `josn:"gender"`
	Roles    []string `json:"roles"`
}

type UserResponseWithTokenSerializer struct {
	UserResponseSerializer
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (u *UserResponseSerializer) FromUserModel(m *database.User) UserResponseSerializer {
	user := UserResponseSerializer{
		Username: m.Username,
		Email:    m.Email,
		Phone:    m.Phone,
		Avatar:   m.Avatar,
		Age:      m.Age,
		Gender:   m.Gender,
		// Roles:    m.Roles,
	}
	return user
}
