package config

import (
	"github.com/Path-IM/Path-IM-Server/common/xorm/global"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql         global.MysqlConfig
	Redis         redis.RedisConf
	RedisDB       int
	TokenSecret   string
	MsgGatewayRpc zrpc.RpcClientConf
}
