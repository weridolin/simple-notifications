package clients

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var appConfig = config.GetAppConfig()
// var logger = config.GetLogger()

type MongoBD struct {
	Uri    string
	Ctx    *context.Context
	Client *mongo.Client
}

func NewMongoDB(uri string, ctx context.Context) *MongoBD {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMaxPoolSize(5))
	if err != nil {
		logx.Info("connect mongodb failed:", err)
	}
	return &MongoBD{
		Uri:    uri,
		Client: client,
		Ctx:    &ctx,
	}
}

// func (m *MongoBD) CreateIndex(info interface{}) error {
// 	...
// }
