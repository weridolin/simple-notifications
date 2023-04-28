package schedulers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/weridolin/simple-vedio-notifications/servers/common"
	"github.com/weridolin/simple-vedio-notifications/tools"
)

func init() {
	validate := validator.New()
	validate.RegisterValidation("cronValidate", tools.CronValidator)
}

type SchedulerValidator struct {
	Period string `json:"period" binding:"required" example:"0 0 0 * * *" form:"period" validate:"cronValidate"` // cron 表达式
	Task   []int  `json:"task" form:"task"`
}

func (u *SchedulerValidator) Bind(c *gin.Context) error {
	err := tools.Bind(c, u)
	if err != nil {
		return err
	}
	return nil
}

type QuerySchedulerValidator struct {
	common.PaginationValidator
	UserID int `json:"user_id" form:"user_id" query:"user_id"`
}

func (u *QuerySchedulerValidator) Bind(c *gin.Context) error {
	err := tools.Bind(c, u)
	if err != nil {
		return err
	}
	return nil
}

type TaskValidator struct {
	PlatForm   string                 `json:"platform" binding:"required" example:"youku" form:"platform"`
	Ups        map[string]interface{} `json:"ups" binding:"required" form:"ups"`
	Schedulers []int                  `json:"schedulers" binding:"required" form:"schedulers"`
	Name       string                 `json:"name"  form:"name"`
}

func (u *TaskValidator) Bind(c *gin.Context) error {
	err := tools.Bind(c, u)
	if err != nil {
		return err
	}
	return nil
}

type QueryTaskValidator struct {
	common.PaginationValidator
	SchedulerID int    `json:"scheduler_id" form:"scheduler_id" query:"scheduler_id" default:"-1"`
	UserID      int    `json:"user_id" form:"user_id" query:"user_id" default:"-1"`
	Name        string `json:"name" form:"name" query:"name" default:""`
}

func (u *QueryTaskValidator) Bind(c *gin.Context) error {
	err := tools.Bind(c, u)
	if err != nil {
		return err
	}
	return nil
}

type BindTasksToSchedulerValidator struct {
	SchedulerID []int `json:"scheduler_id" form:"scheduler_id" query:"scheduler_id" binding:"required"`
	TaskID      []int `json:"task_id" form:"task_id" query:"task_id" binding:"required"`
}

func (u *BindTasksToSchedulerValidator) Bind(c *gin.Context) error {
	err := tools.Bind(c, u)
	if err != nil {
		return err
	}
	return nil
}
