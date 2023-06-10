package storage

import (
	"context"
	"os"
	"path"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileStorage struct {
	Path string
	Ctx  context.Context
}

func NewFileStorage(relativePath string, ctx context.Context) *FileStorage {
	// 获取文件完全路径
	root, _ := os.Getwd()
	absPath := path.Join(root, relativePath)

	return &FileStorage{
		Path: absPath,
		Ctx:  ctx,
	}
}

func (s *FileStorage) Save(info []interface{}) error {
	file, err := os.Open(s.Path)
	if err != nil {
		// logger.Println("open file error -> ", err)
		logx.Error("open file error -> ", err)
		return err
	}
	defer file.Close()
	for _, v := range info {
		file.WriteString(v.(string))
	}
	logx.Info("save info:", info)
	return nil
}

func (s *FileStorage) Remove() error {
	logx.Info("remove info:")
	return nil
}
