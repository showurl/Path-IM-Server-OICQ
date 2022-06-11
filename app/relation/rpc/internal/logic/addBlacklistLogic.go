package logic

import (
	"context"
	"github.com/Path-IM/Path-IM-Server/common/xcache/rc"
	"github.com/Path-IM/Path-IM-Server/common/xorm"
	xormerr "github.com/Path-IM/Path-IM-Server/common/xorm/err"
	relationmodel "github.com/showurl/Path-IM-Server-OICQ/app/relation/model"
	"gorm.io/gorm"
	"time"

	"github.com/showurl/Path-IM-Server-OICQ/app/relation/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddBlacklistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBlacklistLogic {
	return &AddBlacklistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddBlacklistLogic) AddBlacklist(in *pb.AddBlacklistReq) (*pb.AddBlacklistResp, error) {
	createdAt := time.Now().UnixMilli()
	var friends []*relationmodel.Blacklist
	for _, id := range in.RecvUserIds {
		friends = append(friends, &relationmodel.Blacklist{
			SendId:    in.SendUserId,
			ReceiveId: id,
			CreatedAt: createdAt,
		})
	}
	err := xorm.Transaction(l.svcCtx.Mysql,
		// 添加关系
		func(tx *gorm.DB) error {
			err := tx.Model(&relationmodel.Blacklist{}).Create(friends).Error
			if err != nil {
				l.Errorf("create group member error: %v", err)
				return err
			}
			return nil
		},
		// 清理群缓存
		func(tx *gorm.DB) error {
			var keys []string
			// 删除缓存
			relation := rc.NewRelationMapping(l.svcCtx.Mysql, l.svcCtx.Redis)
			// 用户的所有好友
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
		if xormerr.DuplicateError(err) {
			return &pb.AddBlacklistResp{FailedReason: "已经拉黑了"}, nil
		}
		return &pb.AddBlacklistResp{}, err
	}
	// 预热数据
	go func() {
		ctx := context.Background()
		_, err := NewGetBlacklistLogic(ctx, l.svcCtx).GetBlacklist(&pb.GetBlacklistReq{SendUserId: in.SendUserId})
		if err != nil {
			l.Errorf("GetBlacklist error: %v", err)
		}
	}()
	return &pb.AddBlacklistResp{}, nil
}
