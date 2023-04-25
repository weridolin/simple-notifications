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
	// router.POST("/login", Login)
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
