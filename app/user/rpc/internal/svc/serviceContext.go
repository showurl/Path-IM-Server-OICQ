package svc

import (
	onlinemessagerelayservice "github.com/Path-IM/Path-IM-Server/app/msg-gateway/cmd/wsrpc/onlineMessageRelayService"
	"github.com/Path-IM/Path-IM-Server/common/xcache"
	"github.com/Path-IM/Path-IM-Server/common/xcache/global"
	"github.com/Path-IM/Path-IM-Server/common/xorm"
	"github.com/go-redis/redis/v8"
	usermodel "github.com/showurl/Path-IM-Server-OICQ/app/user/model"
	"github.com/showurl/Path-IM-Server-OICQ/app/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	Mysql   *gorm.DB
	Redis   redis.UniversalClient
	wsLogic onlinemessagerelayservice.OnlineMessageRelayService
}

func NewServiceContext(c config.Config) *ServiceContext {
	tx := xorm.GetClient(c.Mysql)
	err := tx.AutoMigrate(
		&usermodel.User{},
		&usermodel.Config{},
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
func (s *ServiceContext) WsLogic() onlinemessagerelayservice.OnlineMessageRelayService {
	if s.wsLogic == nil {
		s.wsLogic = onlinemessagerelayservice.NewOnlineMessageRelayService(zrpc.MustNewClient(s.Config.MsgGatewayRpc))
	}
	return s.wsLogic
}
