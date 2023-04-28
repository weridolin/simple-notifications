package schedulers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/weridolin/simple-vedio-notifications/database"
	"github.com/weridolin/simple-vedio-notifications/servers/common"
	"github.com/weridolin/simple-vedio-notifications/tools"
)

func SchedulerRouteRegister(router *gin.RouterGroup) {
	router.GET("", GetSchedulers)
	router.POST("", AddScheduler)
	router.DELETE("/:id", DeleteScheduler)
	router.PUT("/:id", UpdateScheduler)
}

func TasksRouteRegister(router *gin.RouterGroup) {
	router.GET("", GetTasks)
	router.POST("", CreateTask)
	router.DELETE("/:id", DeleteTask)
	router.PUT("/:id", UpdateTask)
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

	err := database.CreateScheduler(user.(database.User), SchedulerValidator.Period)
	if err != nil {
		common.HttpResponse(c, http.StatusBadRequest, -1, err.Error(), nil)
		return
	}
	common.HttpResponse(c, http.StatusOK, 0, "创建成功", nil)
	return
}

func DeleteScheduler(c *gin.Context) {
	user, _ := c.Get("user")
	schedulerId, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(schedulerId, ">>>>>>>")
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

func UpdateScheduler(c *gin.Context) {
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
	SchedulerValidator := SchedulerValidator{}
	if err := SchedulerValidator.Bind(c); err != nil {
		common.HttpResponse(c, http.StatusUnprocessableEntity, -1, "请求参数错误", tools.NewValidatorError(err))
		return
	}

	err = scheduler.Update(map[string]interface{}{"period": SchedulerValidator.Period})
	if err != nil {
		common.HttpResponse(c, http.StatusBadRequest, -1, err.Error(), nil)
		return
	}
	common.HttpResponse(c, http.StatusOK, 0, "更新成功", nil)
	return
}

// task api

func CreateTask(c *gin.Context) {
	fmt.Println("create a new task...")
	user, _ := c.Get("user")
	taskValidator := TaskValidator{}
	if err := taskValidator.Bind(c); err != nil {
		common.HttpResponse(c, http.StatusUnprocessableEntity, -1, "请求参数错误", tools.NewValidatorError(err))
		return
	}
	err := database.CreateTask(user.(database.User), taskValidator.Ups, taskValidator.PlatForm, taskValidator.Name)
	if err != nil {
		common.HttpResponse(c, http.StatusBadRequest, -1, err.Error(), nil)
		return
	}
	common.HttpResponse(c, http.StatusOK, 0, "创建成功", nil)
	return
}

func DeleteTask(c *gin.Context) {
	user, _ := c.Get("user")
	taskId, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		common.HttpResponse(c, http.StatusBadRequest, -1, "任务id格式错误", nil)
		return
	}
	task, err := database.GetTaskById(taskId)
	if err != nil {
		common.HttpResponse(c, http.StatusNotFound, -1, "任务不存在!", nil)
		return
	}
	if task.UserID != user.(database.User).ID {
		common.HttpResponse(c, http.StatusForbidden, -1, "当前账户没有权限", nil)
		return
	}
	delErr := task.Delete()
	if delErr != nil {
		common.HttpResponse(c, http.StatusBadRequest, -1, delErr.Error(), nil)
		return
	}
	common.HttpResponse(c, http.StatusOK, 0, "删除成功", nil)
	return
}

func UpdateTask(c *gin.Context) {
	user, _ := c.Get("user")
	taskId, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		common.HttpResponse(c, http.StatusBadRequest, -1, "任务id格式错误", nil)
		return
	}
	task, err := database.GetTaskById(taskId)
	if err != nil {
		common.HttpResponse(c, http.StatusNotFound, -1, err.Error(), nil)
		return
	}
	if task.UserID != user.(database.User).ID {
		common.HttpResponse(c, http.StatusForbidden, -1, "当前账户没有权限", nil)
		return
	}
	taskValidator := TaskValidator{}
	if err := taskValidator.Bind(c); err != nil {
		common.HttpResponse(c, http.StatusUnprocessableEntity, -1, "请求参数错误", tools.NewValidatorError(err))
		return
	}
	task.Ups = taskValidator.Ups
	task.PlatForm = taskValidator.PlatForm
	task.Name = taskValidator.Name
	err = task.Update()
	if err != nil {
		common.HttpResponse(c, http.StatusBadRequest, -1, err.Error(), nil)
		return
	}
	common.HttpResponse(c, http.StatusOK, 0, "更新成功", TaskSerializer{}.FromTaskModel(task, user.(database.User)))
	return
}

func BindScheduler(c *gin.Context) {
	// user, _ := c.Get("user")
	// validator := BindTasksToSchedulerValidator{}
	// if err := validator.Bind(c); err != nil {
	// 	common.HttpResponse(c, http.StatusUnprocessableEntity, -1, "请求参数错误", tools.NewValidatorError(err))
	// 	return
	// }

	// return
}

func GetTasks(c *gin.Context) {
	user, _ := c.Get("user")
	validator := QueryTaskValidator{}
	if err := validator.Bind(c); err != nil {
		common.HttpResponse(c, http.StatusUnprocessableEntity, -1, "请求参数错误", tools.NewValidatorError(err))
		return
	}
	tasks, err := database.SearchTasks(user.(database.User), validator.Name, validator.Page, validator.Size, validator.SchedulerID)
	if err != nil {
		common.HttpResponse(c, http.StatusBadGateway, -1, err.Error(), nil)
		return
	}
	serializer := TaskSerializer{}
	common.HttpResponse(c, http.StatusOK, 0, "获取成功", serializer.FromTaskModels(tasks, user.(database.User)))
	return
}
