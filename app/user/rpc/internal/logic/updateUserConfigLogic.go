package logic

import (
	"context"
	"fmt"
	strUtils "github.com/Path-IM/Path-IM-Server-Demo/common/utils/str"
	"github.com/Path-IM/Path-IM-Server/common/xorm"
	usermodel "github.com/showurl/Path-IM-Server-OICQ/app/user/model"
	"gorm.io/gorm"

	"github.com/showurl/Path-IM-Server-OICQ/app/user/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserConfigLogic {
	return &UpdateUserConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

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
