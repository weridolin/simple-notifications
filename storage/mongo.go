package storage

import (
	"context"

	"github.com/weridolin/simple-vedio-notifications/clients"
)

type MongoBDStorage struct {
	Client *clients.MongoBD
}

func NewMongoDBStorage(uri string, ctx context.Context) *MongoBDStorage {
	s := &MongoBDStorage{
		Client: clients.NewMongoDB(uri, ctx),
	}
	return s
}

func (s *MongoBDStorage) Save(info []interface{}) error {
	logger.Println("save info:", info)
	collection := s.Client.Client.Database("Notification").Collection("Result")
	collection.InsertMany(context.Background(), info)
	return nil
}

func (s *MongoBDStorage) Remove() error {
	logger.Println("remove info:")
	return nil
}
