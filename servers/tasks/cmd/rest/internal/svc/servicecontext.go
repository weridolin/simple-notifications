package svc

import (
	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
