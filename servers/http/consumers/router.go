package consumers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	config "github.com/weridolin/simple-vedio-notifications/configs"
	"github.com/weridolin/simple-vedio-notifications/database"
	"github.com/weridolin/simple-vedio-notifications/servers/common"
	"github.com/weridolin/simple-vedio-notifications/tools"
)

var logger = config.GetLogger()

func EmailRouterRegister(router *gin.RouterGroup) {
	router.GET("", GetEmailNotifier)
	router.POST("", CreateEmailNotifier)
	router.DELETE("/:id", DeleteEmailNotifier)
	router.PUT("/:id", UpdateEmailNotifier)
	router.POST("/bind", BindEmailNotifierToTask)

}

func CreateEmailNotifier(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		common.HttpResponse(c, http.StatusUnauthorized, -1, "用户未登录", nil)
		return
	}
	EmailNotifierValidator := EmailNotifierValidator{}
	if err := EmailNotifierValidator.Bind(c); err != nil {
		common.HttpResponse(c, http.StatusUnprocessableEntity, -1, "请求参数错误", tools.NewValidatorError(err))
		return
	}

	err := database.CreateEmailNotifier(user.(database.User), EmailNotifierValidator.PWD,
		EmailNotifierValidator.Sender, EmailNotifierValidator.Content, EmailNotifierValidator.Receiver)
	if err != nil {
		common.HttpResponse(c, http.StatusInternalServerError, -1, err.Error(), nil)
		return
	}
	common.HttpResponse(c, http.StatusOK, 0, "创建成功", nil)
	return
}

func DeleteEmailNotifier(c *gin.Context) {
	user, _ := c.Get("user")
	notifierID, _ := strconv.Atoi(c.Param("id"))
	logger.Println("delete email notifier id --> ", notifierID)
	notifier, err := database.GetEmailNotifierById(notifierID)
	if err != nil {
		common.HttpResponse(c, http.StatusNotFound, -1, err.Error(), nil)
		return
	}
	if notifier.UserID != user.(database.User).ID {
		common.HttpResponse(c, http.StatusForbidden, -1, "当前账户没有权限", nil)
		return
	}
	delErr := notifier.Delete()
	if delErr != nil {
		common.HttpResponse(c, http.StatusBadRequest, -1, delErr.Error(), nil)
		return
	}
	common.HttpResponse(c, http.StatusOK, 0, "删除成功", nil)
	return
}

func UpdateEmailNotifier(c *gin.Context) {
	user, _ := c.Get("user")
	notifierId, _ := strconv.Atoi(c.Param("id"))
	notifier, err := database.GetEmailNotifierById(notifierId)
	if err != nil {
		common.HttpResponse(c, http.StatusNotFound, -1, err.Error(), nil)
		return
	}
	if notifier.UserID != user.(database.User).ID {
		common.HttpResponse(c, http.StatusForbidden, -1, "当前账户没有权限", nil)
		return
	}
	EmailNotifierValidator := EmailNotifierValidator{}
	if err := EmailNotifierValidator.Bind(c); err != nil {
		common.HttpResponse(c, http.StatusUnprocessableEntity, -1, "请求参数错误", tools.NewValidatorError(err))
		return
	}

	err = notifier.Update(map[string]interface{}{"pwd": EmailNotifierValidator.PWD, "sender": EmailNotifierValidator.Sender,
		"content": EmailNotifierValidator.Content, "receiver": EmailNotifierValidator.Receiver})
	if err != nil {
		common.HttpResponse(c, http.StatusInternalServerError, -1, err.Error(), nil)
		return
	}
	common.HttpResponse(c, http.StatusOK, 0, "更新成功", nil)
	return
}

func GetEmailNotifier(c *gin.Context) {
	user, _ := c.Get("user")
	validator := QueryEmailNotifierValidator{}
	if err := validator.Bind(c); err != nil {
		common.HttpResponse(c, http.StatusUnprocessableEntity, -1, "请求参数错误", tools.NewValidatorError(err))
		return
	}
	notifiers, err := database.FullQueryEmailNotify(&database.EmailNotifier{UserID: user.(database.User).ID}, validator.Page, validator.Size)
	if err != nil {
		common.HttpResponse(c, http.StatusInternalServerError, -1, err.Error(), nil)
		return
	}
	serializer := EmailNotifierSerializer{}
	common.HttpResponse(c, http.StatusOK, 0, "获取成功", serializer.FromSchedulerModels(notifiers, user.(database.User)))
	return
}

func BindEmailNotifierToTask(c *gin.Context) {
	user, _ := c.Get("user")
	validator := BindEmailNotifierToTaskValidator{}
	if err := validator.Bind(c); err != nil {
		common.HttpResponse(c, http.StatusUnprocessableEntity, -1, "请求参数错误", tools.NewValidatorError(err))
		return
	}
	err := database.BindEmailNotifierToTask(user.(database.User), validator.TaskID, validator.NotifierIDs)
	if err != nil {
		common.HttpResponse(c, http.StatusInternalServerError, -1, err.Error(), nil)
		return
	}
	common.HttpResponse(c, http.StatusOK, 0, "绑定成功", nil)
	return
}
