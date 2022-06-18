package svc

import (
	"github.com/Path-IM/Path-IM-Server/app/msg/cmd/rpc/chat"
	"github.com/Path-IM/Path-IM-Server/common/xorm"
	"github.com/showurl/Path-IM-Server-OICQ/app/api/internal/config"
	"github.com/showurl/Path-IM-Server-OICQ/app/api/internal/model"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	db     *gorm.DB
	msgRpc chat.Chat
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
func (s *ServiceContext) DB() *gorm.DB {
	if s.db == nil {
		s.db = xorm.GetClient(s.Config.Mysql)
		err := s.db.AutoMigrate(&model.User{})
		if err != nil {
			panic(err)
		}
	}
	return s.db
}
func (s *ServiceContext) MsgRpc() chat.Chat {
	if s.msgRpc == nil {
		s.msgRpc = chat.NewChat(zrpc.MustNewClient(s.Config.MsgRpc))
	}
	return s.msgRpc
}
