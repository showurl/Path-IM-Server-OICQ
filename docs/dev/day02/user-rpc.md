# 上一章：[服务拆分](split-service.md)

---

# 用户模块

## 定义proto

```shell
#!/bin/bash
goctl rpc protoc *.proto -v --go_out=.. --go-grpc_out=..  --zrpc_out=.. --style=goZero --home ../../../../../goctl/home
```

## 定义数据库模型

### 用户表

```shell
mkdir -p app/user/model
```

```go
package model

type User struct {
	Id           string `gorm:"column:id;primary_key;type:char(32);comment:主键;" json:"id"`
	Username     string `gorm:"column:username;type:char(32);index:,unique;not null;comment:用户名;" json:"username"`
	Password     string `gorm:"column:password;type:char(64);not null;comment:密码;" json:"password"`
	Nickname     string `gorm:"column:nickname;type:varchar(64);not null;default:'';comment:昵称;index;" json:"nickname"`
	Sign         string `gorm:"column:sign;type:varchar(128);not null;default:'';comment:签名;" json:"sign"`
	Avatar       string `gorm:"column:avatar;type:varchar(255);not null;default:'';comment:头像;" json:"avatar"`
	Province     string `gorm:"column:province;type:varchar(64);not null;default:'';comment:省份;" json:"province"`
	City         string `gorm:"column:city;type:varchar(64);not null;default:'';comment:城市;" json:"city"`
	District     string `gorm:"column:district;type:varchar(64);not null;default:'';comment:区县;" json:"district"`
	RegisterTime int64  `gorm:"column:register_time;type:bigint(13);not null;default:0;comment:注册时间-毫秒级时间戳;" json:"registerTime"`
	IsMale       bool   `gorm:"column:is_male;index;comment:是否是男性;" json:"isMale"`
}
```

### 用户配置表

```go
package model

type Config struct {
	UserId string `gorm:"column:user_id;index;type:char(32);comment:用户id;not null;" json:"userId"`
	K      string `gorm:"column:k;index;type:varchar(255);index;comment:配置键;not null;" json:"k"`
	V      string `gorm:"column:v;type:varchar(255);comment:配置值;not null;" json:"v"`
}
```

## 配置数据库

### config

```go
package config

type Config struct {
	zrpc.RpcServerConf
	Mysql   global.MysqlConfig
	Redis   redis.RedisConf
	RedisDB int
}
```

### service context

```go
package svc

type ServiceContext struct {
	Config config.Config
	Mysql  *gorm.DB
	Redis  redis.UniversalClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Mysql:  xorm.GetClient(c.Mysql),
		Redis:  xcache.GetClient(c.Redis, global.DB(c.RedisDB)),
	}
}
```

## 写逻辑层

### 用户注册

```go
package logic

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	user := &usermodel.User{
		Username: in.Username,
		Password: in.Password,
	}
	err := xorm.Transaction(l.svcCtx.Mysql,
		// 插入用户
		user.FuncInsert(l.svcCtx.Redis),
		// 清除用户相关的缓存
		func(tx *gorm.DB) error {
			err := user.FlushCache(l.svcCtx.Mysql, l.svcCtx.Redis)
			return err
		})
	if err != nil {
		return &pb.RegisterResp{
			FailReason: "操作数据库错误：" + err.Error(),
			UserId:     "",
		}, xerr.New(500, err.Error())
	}
	// 预热数据
	go user.Preheat(l.svcCtx.Mysql, l.svcCtx.Redis)
	return &pb.RegisterResp{
		FailReason: "",
		UserId:     user.Id,
	}, nil
}
```

### 用户获取token

> token 使用 [JWT](https://jwt.io/) 算法生成

> 存储到 redis 中，key 为 `token:${platform}:${user_id}`

> value 为 hashMap，存储的内容为：

    - `token`: 状态(正常、被踢、被封禁)

```go
package logic

func (l *GetTokenLogic) GetToken(in *pb.GetTokenReq) (*pb.GetTokenResp, error) {
	claim := jwtUtils.BuildClaims(in.UserId, in.Platform)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedString, err := token.SignedString([]byte(l.svcCtx.Config.TokenSecret))
	if err != nil {
		l.Errorf("jwt signed string error: %v", err)
		return nil, err
	}
	redisKey := global.MergeKey("token", in.Platform, in.UserId)
	// 踢出原来的token和socket连接
	tokenMap, err := l.svcCtx.Redis.HGetAll(l.ctx, redisKey).Result()
	if err != nil {
		l.Errorf("redis hgetall error: %v", err)
		return nil, err
	}
	if len(tokenMap) > 0 {
		for token, status := range tokenMap {
			switch status {
			case "0":
				// 改为过期
				err = l.svcCtx.Redis.HSet(l.ctx, redisKey, token, "1").Err()
				if err != nil {
					l.Errorf("redis hset error: %v", err)
					return nil, err
				}
				// 断开连接
				_, err = l.svcCtx.WsLogic().KickUserConns(l.ctx, &msggatewaypb.KickUserConnsReq{
					UserID:      in.UserId,
					PlatformIDs: []string{in.Platform},
				})
				if err != nil {
					l.Errorf("kick user conns error: %v", err)
					return nil, err
				}
			case "1":
				// 删除
				e := l.svcCtx.Redis.HDel(l.ctx, redisKey, token).Err()
				if e != nil {
					l.Errorf("redis hdel error: %v", e)
				}
			}
		}
	}
	// 存储新的token
	err = l.svcCtx.Redis.HSet(l.ctx, redisKey, signedString, "0").Err()
	if err != nil {
		l.Errorf("redis hset error: %v", err)
		return nil, err
	}
	return &pb.GetTokenResp{
		Token: signedString,
	}, nil
}

```

---

# 下一章：[用户rpc](../day03/user-rpc.md)