package logic

import (
	"context"
	"errors"
	"github.com/Path-IM/Path-IM-Server/common/utils/encrypt"
	"github.com/Path-IM/Path-IM-Server/common/xorm"
	usermodel "github.com/showurl/Path-IM-Server-OICQ/app/user/model"
	"gorm.io/gorm"

	"github.com/showurl/Path-IM-Server-OICQ/app/user/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

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
