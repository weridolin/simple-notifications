package database

import (
	"fmt"

	config "github.com/weridolin/simple-vedio-notifications/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB //指针

//root:werido@8.131.78.84:3306/sitebackend?charset=utf8mb4

func init() {
	dsn := config.GetAppConfig().DBUri
	fmt.Println("dsn: ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("db connect err: (Init) ", err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
