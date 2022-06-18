package svc

import (
	"github.com/Path-IM/Path-IM-Server/common/xorm"
	"github.com/showurl/Path-IM-Server-OICQ/app/imuser/rpc/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	db     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
func (s *ServiceContext) DB() *gorm.DB {
	if s.db == nil {
		s.db = xorm.GetClient(s.Config.Mysql)
	}
	return s.db
}
