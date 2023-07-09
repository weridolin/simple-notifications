package svc

import (
	"github.com/weridolin/simple-vedio-notifications/servers/users/cmd/rest/internal/config"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/weridolin/simple-vedio-notifications/servers/users/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	Config    config.Config
	DB        *gorm.DB
	UserModel models.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	db, err := gorm.Open(mysql.Open("root:werido@tcp(8.131.78.84:3306)/simple_notification?charset=utf8mb4&parseTime=true"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix:   "auth_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	if err != nil {
		logx.Error(err)
	}
	//自动同步更新表结构
	db.AutoMigrate(&models.User{})

	return &ServiceContext{
		Config:    c,
		DB:        db,
		UserModel: models.NewUserModel("user"),
	}
}
