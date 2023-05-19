package storage

import (
	"context"

	"github.com/weridolin/simple-vedio-notifications/clients"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoBDStorage struct {
	Client         *clients.MongoBD
	CollectionName string
	dbName         string
}

func NewMongoDBStorage(uri string, dbName, CollectionName string, ctx context.Context) *MongoBDStorage {
	s := &MongoBDStorage{
		Client:         clients.NewMongoDB(uri, ctx),
		dbName:         dbName,
		CollectionName: CollectionName,
	}
	return s
}

func (s *MongoBDStorage) CreateIndex(fields mongo.IndexModel, options *options.CreateIndexesOptions) error {
	collection := s.Client.Client.Database(s.dbName).Collection(s.CollectionName)
	collection.Indexes().CreateOne(context.Background(), fields, options)
	return nil
}

func (s *MongoBDStorage) Save(info []interface{}) error {
	logger.Println("save info:", info)
	collection := s.Client.Client.Database(s.dbName).Collection(s.CollectionName)
	// todo insert or update?
	_, err := collection.InsertMany(context.Background(), info)
	if nil != err {
		logger.Println("insert result into mongodb error -> ", err)
	}
	return nil
}

func (s *MongoBDStorage) Remove() error {
	logger.Println("remove info:")
	return nil
}
