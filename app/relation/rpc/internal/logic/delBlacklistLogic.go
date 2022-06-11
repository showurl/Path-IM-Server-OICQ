package logic

import (
	"context"
	"github.com/Path-IM/Path-IM-Server/common/xcache/rc"
	"github.com/Path-IM/Path-IM-Server/common/xorm"
	relationmodel "github.com/showurl/Path-IM-Server-OICQ/app/relation/model"
	"gorm.io/gorm"

	"github.com/showurl/Path-IM-Server-OICQ/app/relation/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelBlacklistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelBlacklistLogic {
	return &DelBlacklistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelBlacklistLogic) DelBlacklist(in *pb.DelBlacklistReq) (*pb.DelBlacklistResp, error) {
	err := xorm.Transaction(l.svcCtx.Mysql,
		// 删除黑名单
		func(tx *gorm.DB) error {
			err := tx.Where("send_id = ? AND receive_id in (?)", in.SendUserId, in.RecvUserIds).Delete(&relationmodel.Blacklist{}).Error
			if err != nil {
				l.Errorf("delete blacklist error: %v", err)
				return err
			}
			return nil
		},
		// 清理黑名单缓存
		func(tx *gorm.DB) error {
			var keys []string
			// 删除缓存
			relation := rc.NewRelationMapping(l.svcCtx.Mysql, l.svcCtx.Redis)
			// 用户的所有黑名单
			relationKey1, _ := relation.Key(&relationmodel.Blacklist{}, "receive_id", map[string]interface{}{
				"send_id": in.SendUserId,
			}, rc.Order("created_at"))
			keys = append(keys, relationKey1)
			// 删除缓存
			err := l.svcCtx.Redis.Del(l.ctx, keys...).Err()
			if err != nil {
				l.Errorf("delete cache error: %v", err)
				return err
			}
			return nil
		},
	)
	if err != nil {
		return &pb.DelBlacklistResp{}, err
	}
	// 预热数据
	go func() {
		ctx := context.Background()
		_, err := NewGetBlacklistLogic(ctx, l.svcCtx).GetBlacklist(&pb.GetBlacklistReq{SendUserId: in.SendUserId})
		if err != nil {
			l.Errorf("GetBlacklist error: %v", err)
		}
	}()
	return &pb.DelBlacklistResp{}, nil
}
