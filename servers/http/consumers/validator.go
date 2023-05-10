package consumers

import (
	"github.com/gin-gonic/gin"
	"github.com/weridolin/simple-vedio-notifications/database"
	"github.com/weridolin/simple-vedio-notifications/servers/common"
	"github.com/weridolin/simple-vedio-notifications/tools"
)

type EmailNotifierValidator struct {
	Sender   string                 `json:"sender" binding:"required,email" form:"sender"`
	PWD      string                 `json:"pwd" binding:"required" form:"pwd"`
	Receiver database.EmailReceiver `json:"receiver" binding:"required" form:"receiver"`
	Content  string                 `json:"content" binding:"required" form:"content"`
}

func (u *EmailNotifierValidator) Bind(c *gin.Context) error {
	err := tools.Bind(c, u)
	if err != nil {
		return err
	}
	return nil
}

type QueryEmailNotifierValidator struct {
	common.PaginationValidator
	UserID int `json:"user_id" form:"user_id"`
}

func (q *QueryEmailNotifierValidator) Bind(c *gin.Context) error {
	err := tools.Bind(c, q)
	if err != nil {
		return err
	}
	return nil
}

type BindEmailNotifierToTaskValidator struct {
	TaskID      uint   `json:"task_id" form:"task_id" query:"task_id" binding:"required"`
	NotifierIDs []uint `json:"notifier_ids" form:"notifier_ids" query:"notifier_ids" binding:"required"`
}

func (q *BindEmailNotifierToTaskValidator) Bind(c *gin.Context) error {
	err := tools.Bind(c, q)
	if err != nil {
		return err
	}
	return nil
}
