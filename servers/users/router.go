package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/weridolin/simple-vedio-notifications/database"
	"github.com/weridolin/simple-vedio-notifications/servers/common"
	"github.com/weridolin/simple-vedio-notifications/tools"
)

func UnAuthorizationRouteRegister(router *gin.RouterGroup) {
	router.POST("/register", Register)
	router.POST("/login", Login)
}

func AuthorizationRouteRegister(router *gin.RouterGroup) {
	router.POST("/logout", Logout)
}

func Register(c *gin.Context) {
	UserValidator := database.NewUserValidator()
	if err := UserValidator.Bind(c); err != nil {
		common.HttpResponse(c, http.StatusUnprocessableEntity, -1, "请求参数错误", tools.NewValidatorError(err))
		return
	}
	fmt.Println("create a new user", UserValidator)

	if user, err := UserValidator.Create(); err != nil {
		common.HttpResponse(c, http.StatusBadRequest, -1, "用户名或者邮箱已经存在", tools.NewError("database", err))
		return
	} else {
		serializer := UserResponseSerializer{}
		common.HttpResponse(c, http.StatusOK, 0, "注册成功", serializer.FromUserModel(user))
		return
	}
}

func Login(c *gin.Context) {
	validator := LoginRequestValidator{}
	if err := validator.Bind(c); err != nil {
		common.HttpResponse(c, http.StatusUnprocessableEntity, -1, "请求参数错误", tools.NewValidatorError(err))
		return
	}
	fmt.Println("login a user", validator)
	serializer := UserResponseWithTokenSerializer{}
	user, err := validator.CheckPWd()
	if err != nil {
		common.HttpResponse(c, http.StatusBadRequest, -1, err.Error(), nil)
		return
	}
	accessToken := database.GenToken(*user)
	userInfo := UserResponseSerializer{}
	userInfo.FromUserModel(user)
	serializer = UserResponseWithTokenSerializer{
		UserResponseSerializer: userInfo,
		AccessToken:            accessToken,
	}
	common.HttpResponse(c, http.StatusOK, 0, "登录成功", serializer)
	return
}

func Logout(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		common.HttpResponse(c, http.StatusUnauthorized, -1, "用户未登录", nil)
		return
	}
	fmt.Println("logout a user", user)
	common.HttpResponse(c, http.StatusOK, 0, "登出成功", nil)
	return
}
