package svc

import "github.com/showurl/Path-IM-Server-OICQ/app/imuser/rpc/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
