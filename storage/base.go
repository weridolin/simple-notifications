package storage

import (
	"context"

	config "github.com/weridolin/simple-vedio-notifications/configs"
)

var appConfig = config.GetAppConfig()
var logger = config.GetLogger()

type StorageInterface interface {
	Save(info []interface{}) error
	Remove() error
}

func NewStorage(ctx context.Context) StorageInterface {
	var storage StorageInterface
	if appConfig.StorageType == "mongodb" {
		storage = NewMongoDBStorage(appConfig.MongoDbUri, ctx)
	}
	return storage
}
