package svc

import (
	"github.com/weridolin/simple-vedio-notifications/servers/consumers/cmd/rest/internal/config"
	"github.com/weridolin/simple-vedio-notifications/servers/consumers/models"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	Config             config.Config
	DB                 *gorm.DB
	EmailNotifierModel models.EmailNotifierModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.SetUp(c.Logger.LogConf)
	db, err := gorm.Open(mysql.Open(c.DBUri), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix:   "auth_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})

	if err != nil {
		logx.Error(err)
	}
	//自动同步更新表结构
	db.AutoMigrate(&models.EmailNotifier{}, &models.EmailNotifierTask{})
	return &ServiceContext{
		Config:             c,
		DB:                 db,
		EmailNotifierModel: models.NewDefaultEmailNotifierModel("email_notifier"),
	}
}
