package storage

import (
	"context"

	config "github.com/weridolin/simple-vedio-notifications/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var appConfig = config.GetAppConfig()
var logger = config.GetLogger()

type StorageInterface interface {
	Save(info []interface{}) error
	Remove() error
}

func NewStorage(ctx context.Context) StorageInterface {
	var storage StorageInterface
	switch appConfig.StorageType {
	case "mongodb":
		{
			StorageInstance := NewMongoDBStorage(appConfig.MongoDbUri, appConfig.MongoDbName, "Result", ctx)
			StorageInstance.CreateIndex(mongo.IndexModel{
				Keys: map[string]int{"videoinfo.created": 1}}, nil)
			StorageInstance.CreateIndex(mongo.IndexModel{
				Keys: map[string]int{"upname": 1}}, nil)
			//全文搜索
			StorageInstance.CreateIndex(mongo.IndexModel{
				Keys: map[string]string{"videoinfo.title": "text"}}, nil)
			//创建一个唯一索引
			StorageInstance.CreateIndex(mongo.IndexModel{
				Keys:    map[string]int{"videoinfo.bvid": 1},
				Options: options.Index().SetUnique(true)}, nil)
			storage = StorageInstance

		}
	case "file":
		{
			StorageInstance := NewFileStorage(appConfig.StorageFileRelativePath, ctx)
			storage = StorageInstance
		}
	}
	return storage
}
