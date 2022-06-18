package config

import (
	"github.com/Path-IM/Path-IM-Server/common/xorm/global"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Mysql  global.MysqlConfig
	MsgRpc zrpc.RpcClientConf
}
