package schedulers

import (
	"github.com/gin-gonic/gin"
	"github.com/weridolin/simple-vedio-notifications/servers/common"
	"github.com/weridolin/simple-vedio-notifications/tools"
)

type SchedulerValidator struct {
	Period string `json:"period" binding:"required,cronValidate" example:"0 0 0 * * *" form:"period"`
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
