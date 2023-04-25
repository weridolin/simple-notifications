package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/weridolin/simple-vedio-notifications/database"
	"github.com/weridolin/simple-vedio-notifications/tools"
)

func RouteRegister(router *gin.RouterGroup) {
	router.POST("/register", Register)
	router.POST("/login", Login)
}

func Register(c *gin.Context) {
	UserValidator := database.NewUserValidator()
	if err := UserValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, tools.NewValidatorError(err)) // 422 请求参数错误
		return
	}
	fmt.Println("create a new user", UserValidator)

	if user, err := UserValidator.Create(); err != nil {
		c.JSON(http.StatusUnprocessableEntity, tools.NewError("database", err))
		return
	} else {
		serializer := UserResponseSerializer{}
		c.JSON(http.StatusCreated, gin.H{"user": serializer.FromUserModel(user)})
		return
	}
}

func Login(c *gin.Context) {
	validator := LoginRequestValidator{}
	if err := validator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, tools.NewValidatorError(err)) // 422 请求参数错误
		return
	}
	fmt.Println("login a user", validator)
	serializer := UserResponseWithTokenSerializer{}
	user, err := validator.CheckPWd()
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.NewError("database", err))
		return
	}
	accessToken := database.GenToken(user)
	userInfo := UserResponseSerializer{}
	userInfo.FromUserModel(&user)
	serializer = UserResponseWithTokenSerializer{
		UserResponseSerializer: userInfo,
		AccessToken:            accessToken,
	}
	c.JSON(http.StatusOK, gin.H{"data": serializer})
	return
}
