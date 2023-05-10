package database

import (
	config "github.com/weridolin/simple-vedio-notifications/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB //指针
var logger = config.GetLogger()

//root:werido@8.131.78.84:3306/sitebackend?charset=utf8mb4

func init() {
	dsn := config.GetAppConfig().DBUri
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Panicln("db connect err: (Init) ", err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
