package logic

import (
	"context"
	"fmt"
	strUtils "github.com/Path-IM/Path-IM-Server-Demo/common/utils/str"
	usermodel "github.com/showurl/Path-IM-Server-OICQ/app/user/model"

	"github.com/showurl/Path-IM-Server-OICQ/app/user/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserConfigLogic {
	return &GetUserConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

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
