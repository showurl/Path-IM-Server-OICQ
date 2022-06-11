package svc

import (
	"github.com/Path-IM/Path-IM-Server/common/xcache"
	"github.com/Path-IM/Path-IM-Server/common/xcache/global"
	"github.com/Path-IM/Path-IM-Server/common/xorm"
	"github.com/go-redis/redis/v8"
	groupmodel "github.com/showurl/Path-IM-Server-OICQ/app/group/model"
	"github.com/showurl/Path-IM-Server-OICQ/app/group/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Mysql  *gorm.DB
	Redis  redis.UniversalClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	tx := xorm.GetClient(c.Mysql)
	err := tx.AutoMigrate(
		&groupmodel.Group{},
		&groupmodel.GroupMember{},
	)
	if err != nil {
		logx.Errorf("auto migrate error: %v", err)
		panic(err)
	}
	return &ServiceContext{
		Config: c,
		Mysql:  tx,
		Redis:  xcache.GetClient(c.Redis, global.DB(c.RedisDB)),
	}
}
