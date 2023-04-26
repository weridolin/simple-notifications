package schedulers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/weridolin/simple-vedio-notifications/database"
	"github.com/weridolin/simple-vedio-notifications/servers/common"
	"github.com/weridolin/simple-vedio-notifications/tools"
)

// func UnAuthorizationRouteRegister(router *gin.RouterGroup) {
// 	router.GET("", GetSchedulers)

// }

func AuthorizationRouteRegister(router *gin.RouterGroup) {
	router.GET("", GetSchedulers)
	router.POST("", AddScheduler)
	router.DELETE("/:id", DeleteScheduler)
}

func GetSchedulers(c *gin.Context) {
	user, _ := c.Get("user")
	validator := QuerySchedulerValidator{}
	if err := validator.Bind(c); err != nil {
		common.HttpResponse(c, http.StatusUnprocessableEntity, -1, "请求参数错误", tools.NewValidatorError(err))
		return
	}
	schedulers, err := database.QuerySchedulers(map[string]interface{}{"user_id": user.(database.User).ID}, validator.Page, validator.Size)
	if err != nil {
		common.HttpResponse(c, http.StatusBadGateway, -1, err.Error(), nil)
		return
	}
	serializer := SchedulerSerializer{}
	common.HttpResponse(c, http.StatusOK, 0, "获取成功", serializer.FromSchedulerModel(schedulers, user.(database.User)))
	return

}

func AddScheduler(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		common.HttpResponse(c, http.StatusUnauthorized, -1, "用户未登录", nil)
		return
	}
	SchedulerValidator := SchedulerValidator{}
	if err := SchedulerValidator.Bind(c); err != nil {
		common.HttpResponse(c, http.StatusUnprocessableEntity, -1, "请求参数错误", tools.NewValidatorError(err))
		return
	}
	fmt.Println("create a new scheduler", SchedulerValidator)
	err := database.CreateScheduler(user.(database.User), SchedulerValidator.Period)
	if err != nil {
		common.HttpResponse(c, http.StatusBadRequest, -1, err.Error(), nil)
	}
	common.HttpResponse(c, http.StatusOK, 0, "创建成功", nil)
	return
}

func DeleteScheduler(c *gin.Context) {
	user, _ := c.Get("user")
	schedulerId := c.Param("id")
	scheduler, err := database.QueryScheduler(map[string]interface{}{"id": schedulerId})
	if err != nil {
		common.HttpResponse(c, http.StatusNotFound, -1, err.Error(), nil)
		return
	}
	if scheduler.UserID != user.(database.User).ID {
		common.HttpResponse(c, http.StatusForbidden, -1, "当前账户没有权限", nil)
		return
	}
	delErr := scheduler.Delete()
	if delErr != nil {
		common.HttpResponse(c, http.StatusBadRequest, -1, delErr.Error(), nil)
		return
	}
	common.HttpResponse(c, http.StatusOK, 0, "删除成功", nil)
	return
}
