package storage

import (
	"context"

	// config "github.com/weridolin/simple-vedio-notifications/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var appConfig = config.GetAppConfig()
// var logger = config.GetLogger()

type StorageInterface interface {
	Save(info []interface{}) error
	Remove() error
}

type StorageParams interface {
}

type MongoBDStorageParams struct {
	MongoDbUri     string
	MongoDbName    string
	CollectionName string
}

type FileStorageParams struct {
	StorageFileRelativePath string
}

type MySqlStorageParams struct {
}

func NewStorage(ctx context.Context, storageType string, storageParams StorageParams) StorageInterface {
	var storage StorageInterface
	switch storageType {
	case "mongodb":
		{
			StorageInstance := NewMongoDBStorage(storageParams.(MongoBDStorageParams).MongoDbUri, storageParams.(MongoBDStorageParams).MongoDbName, storageParams.(MongoBDStorageParams).CollectionName, ctx)
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
			StorageInstance := NewFileStorage(storageParams.(FileStorageParams).StorageFileRelativePath, ctx)
			storage = StorageInstance
		}
	}
	return storage
}
