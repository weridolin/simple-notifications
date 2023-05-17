package storage

import (
	"context"
	"os"
)

type FileStorage struct {
	Path string
	Ctx  context.Context
}

func NewFileStorage(path string, ctx context.Context) *FileStorage {
	return &FileStorage{
		Path: path,
		Ctx:  ctx,
	}
}

func (s *FileStorage) Save(info interface{}) error {
	file, err := os.Open(s.Path)
	if err != nil {
		logger.Println("open file error -> ", err)
		return err
	}
	defer file.Close()
	file.WriteString(info.(string))
	logger.Println("save info:", info)
	return nil
}

func (s *FileStorage) Remove() error {
	logger.Println("remove info:")
	return nil
}
