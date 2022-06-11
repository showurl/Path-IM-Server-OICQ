# 上一章：[用户模块](../day02/user-rpc.md)

---

# 用户rpc

## 编写rpc

### 用户名查询用户(1)

> 这个很简单 就是sql查询

```go
package logic

func (l *GetUserByUsernameLogic) GetUserByUsername(in *pb.GetUserByUsernameReq) (*pb.GetUserByUsernameResp, error) {
	user := &usermodel.User{}
	err := l.svcCtx.Mysql.Model(user).Where("username = ?", in.Username).First(user).Error
	if err != nil {
		if xormerr.RecordNotFound(err) {
			err = nil
		} else {
			l.Errorf("mysql first error: %v", err)
			return nil, err
		}
	}
	return &pb.GetUserByUsernameResp{
		FailReason: "",
		User:       user.Bytes(),
	}, nil
}
```

### 用户信息查询(n)

> 用id查询用户 加缓存

```go
package logic

func (l *GetUserByIdsLogic) GetUserByIds(in *pb.GetUserByIdsReq) (*pb.GetUserByIdsResp, error) {
	mapping := dc.GetDbMapping(l.svcCtx.Redis, l.svcCtx.Mysql)
	var users []*usermodel.User
	var userBytes [][]byte
	err := mapping.ListByIds(&usermodel.User{}, &users, in.UserId)
	if err != nil {
		l.Errorf("get user by ids error: %v", err)
		return &pb.GetUserByIdsResp{}, err
	}
	for _, user := range users {
		userBytes = append(userBytes, user.Bytes())
	}
	return &pb.GetUserByIdsResp{Users: userBytes}, nil
}
```

### 用户信息修改

> 注意缓存的更新

```go
package main

func (l *UpdateUserLogic) UpdateUser(in *pb.UpdateUserReq) (*pb.UpdateUserResp, error) {
	var updateMap = make(map[string]interface{})
	// 过滤字段
	for k, v := range in.UpdateMaps {
		// 支持的字段
		switch k {
		case "nickname", "sign", "avatar", "province", "city", "district":
			updateMap[k] = v
		case "password":
			// 密码加密
			updateMap[k] = encrypt.Md5(v)
		default:
			l.Errorf("update user field error: %v", k)
			return &pb.UpdateUserResp{FailReason: ""}, errors.New("不支持的字段值")
		}
	}
	user := &usermodel.User{}
	user.Id = in.UserId
	err := xorm.Transaction(l.svcCtx.Mysql,
		func(tx *gorm.DB) error {
			// 更新用户信息
			err := tx.Model(user).Where("id = ?", in.UserId).Updates(updateMap).Error
			if err != nil {
				l.Errorf("update user error: %v", err)
				return err
			}
			return nil
		}, func(tx *gorm.DB) error {
			// 清理用户缓存
			err := user.FlushCache(tx, l.svcCtx.Redis)
			if err != nil {
				l.Errorf("flush user cache error: %v", err)
				return err
			}
			return nil
		},
	)
	if err != nil {
		l.Errorf("update user error: %v", err)
		return &pb.UpdateUserResp{FailReason: ""}, err
	}
	// 预热数据
	go user.Preheat(l.svcCtx.Mysql, l.svcCtx.Redis)
	return &pb.UpdateUserResp{}, nil
}
```

### 用户配置(kv存储)查询

```go
package logic

func (l *GetUserConfigLogic) GetUserConfig(in *pb.GetUserConfigReq) (*pb.GetUserConfigResp, error) {
	// 先从缓存中获取
	// 是否存在缓存
	key := fmt.Sprintf("%s%s", usermodel.RedisKeyConfig, in.UserId)
	exist, _ := l.svcCtx.Redis.Exists(l.ctx, key).Result()
	var configMap = make(map[string]string)
	if exist == 0 {
		// 不存在缓存
		// 数据库中查询
		var configs []*usermodel.Config
		err := l.svcCtx.Mysql.Model(&usermodel.Config{}).Where("user_id = ?", in.UserId).Find(&configs).Error
		if err != nil {
			l.Errorf("get user config error: %v", err)
			return &pb.GetUserConfigResp{FailReason: ""}, err
		}
		// 存入缓存
		var kvs []interface{}
		for _, config := range configs {
			kvs = append(kvs, config.K, config.V)
			if (len(in.Ks) > 0 && strUtils.IsContain(config.K, in.Ks)) || len(in.Ks) == 0 {
				configMap[config.K] = config.V
			}
		}
		err = l.svcCtx.Redis.HMSet(l.ctx, key, kvs...).Err()
		if err != nil {
			l.Errorf("set user config to redis error: %v", err)
			err = nil
		}
		return &pb.GetUserConfigResp{
			FailReason: "",
			Configs:    configMap,
		}, nil
	}
	resultMap, err := l.svcCtx.Redis.HGetAll(l.ctx, key).Result()
	if err != nil {
		return nil, err
	}
	for k, v := range resultMap {
		if (len(in.Ks) > 0 && strUtils.IsContain(k, in.Ks)) || len(in.Ks) == 0 {
			configMap[k] = v
		}
	}
	return &pb.GetUserConfigResp{Configs: configMap}, nil
}
```

### 用户配置修改

> 注意缓存的更新

```go
package logic

func (l *UpdateUserConfigLogic) UpdateUserConfig(in *pb.UpdateUserConfigReq) (*pb.UpdateUserConfigResp, error) {
	var ks []string
	var existKs []string
	for k := range in.Configs {
		ks = append(ks, k)
	}
	// 查询配置项是否存在
	logic := NewGetUserConfigLogic(l.ctx, l.svcCtx)
	userConfigResp, err := logic.GetUserConfig(&pb.GetUserConfigReq{
		UserId: in.UserId,
		Ks:     ks,
	})
	if err != nil {
		l.Errorf("get user config error: %v", err)
		return &pb.UpdateUserConfigResp{FailReason: ""}, err
	}
	var updateMap = make(map[string]string)
	for k, v := range userConfigResp.Configs {
		if strUtils.IsContain(k, ks) {
			existKs = append(existKs, k)
			if v != in.Configs[k] {
				updateMap[k] = in.Configs[k]
			}
		}
	}
	// 不存在的key
	var notExistKs = strUtils.DifferenceString(ks, existKs)
	var insertMap = make(map[string]string)
	for _, k := range notExistKs {
		insertMap[k] = in.Configs[k]
	}
	if len(updateMap) == 0 && len(insertMap) == 0 {
		return &pb.UpdateUserConfigResp{}, nil
	}
	err = xorm.Transaction(l.svcCtx.Mysql,
		func(tx *gorm.DB) error {
			// 修改数据库
			for k, v := range updateMap {
				err := tx.Model(&usermodel.Config{}).Where("user_id = ?", in.UserId).Update(k, v).Error
				if err != nil {
					l.Errorf("update user config error: %v", err)
					return err
				}
			}
			// 插入数据库
			for k, v := range insertMap {
				err := tx.Create(&usermodel.Config{
					UserId: in.UserId,
					K:      k,
					V:      v,
				}).Error
				if err != nil {
					l.Errorf("insert user config error: %v", err)
					return err
				}
			}
			return nil
		},
		func(tx *gorm.DB) error {
			// 清理缓存
			key := fmt.Sprintf("%s%s", usermodel.RedisKeyConfig, in.UserId)
			err := l.svcCtx.Redis.Del(l.ctx, key).Err()
			if err != nil {
				l.Errorf("del redis key error: %v", err)
				return err
			}
			return nil
		},
	)
	if err != nil {
		l.Errorf("update user config error: %v", err)
		return &pb.UpdateUserConfigResp{FailReason: ""}, err
	}
	// 预热缓存
	_, err = logic.GetUserConfig(&pb.GetUserConfigReq{UserId: in.UserId})
	if err != nil {
		l.Errorf("get user config error: %v", err)
	}
	return &pb.UpdateUserConfigResp{}, nil
}

```

---

## 测试rpc

```go
package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Path-IM/Path-IM-Server/common/utils/encrypt"
	usermodel "github.com/showurl/Path-IM-Server-OICQ/app/user/model"
	"github.com/showurl/Path-IM-Server-OICQ/app/user/rpc/pb"
	"github.com/showurl/Path-IM-Server-OICQ/app/user/rpc/userservice"
	"github.com/zeromicro/go-zero/zrpc"
	"testing"
)

var (
	service userservice.UserService
	ctx     = context.Background()
)

func init() {
	service = userservice.NewUserService(zrpc.MustNewClient(zrpc.RpcClientConf{
		Endpoints: []string{"192.168.1.98:10013"},
	}))
}
func TestRegister(t *testing.T) {
	resp, err := service.Register(ctx, &pb.RegisterReq{
		Username: fmt.Sprintf("test-%d", 1),
		Password: "123456",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
}
func TestGetToken(t *testing.T) {
	resp, err := service.GetToken(ctx, &pb.GetTokenReq{
		UserId:   "test-1",
		Platform: "IOS",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
}
func TestGetUserByUsername(t *testing.T) {
	resp, err := service.GetUserByUsername(ctx, &pb.GetUserByUsernameReq{
		Username: "test-1",
	})
	if err != nil {
		t.Error(err)
	}
	user := &usermodel.User{}
	_ = json.Unmarshal(resp.User, user)
	t.Log(resp.String())
	t.Log(user.Password)
	t.Log(encrypt.Md5("123456"))
}
func TestGetUserByIds(t *testing.T) {
	resp, err := service.GetUserByIds(ctx, &pb.GetUserByIdsReq{
		UserId: []string{"3604f15eed0f94282e271ed583427eec", "test-1"},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
}
func TestGetUserConfig(t *testing.T) {
	resp, err := service.GetUserConfig(ctx, &pb.GetUserConfigReq{
		UserId: "3604f15eed0f94282e271ed583427eec",
		Ks:     []string{},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
	resp, err = service.GetUserConfig(ctx, &pb.GetUserConfigReq{
		UserId: "3604f15eed0f94282e271ed583427eec",
		Ks: []string{
			"k1", "k2",
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
}
func TestUpdateUserConfig(t *testing.T) {
	resp, err := service.UpdateUserConfig(ctx, &pb.UpdateUserConfigReq{
		UserId: "3604f15eed0f94282e271ed583427eec",
		Configs: map[string]string{
			"k1": "v1",
			"k2": "v2",
			"k3": "v3",
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
}

```

# 下一章：[群组模块](group-rpc.md)