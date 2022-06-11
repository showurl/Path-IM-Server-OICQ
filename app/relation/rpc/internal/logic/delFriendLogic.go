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

type DelFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelFriendLogic {
	return &DelFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelFriendLogic) DelFriend(in *pb.DelFriendReq) (*pb.DelFriendResp, error) {
	err := xorm.Transaction(l.svcCtx.Mysql,
		// 删除黑名单
		func(tx *gorm.DB) error {
			err := tx.Where("send_id = ? AND receive_id in (?)", in.SendUserId, in.RecvUserIds).Delete(&relationmodel.Friend{}).Error
			if err != nil {
				l.Errorf("delete Friend error: %v", err)
				return err
			}
			return nil
		},
		// 清理黑名单缓存
		func(tx *gorm.DB) error {
			var keys []string
			// 删除缓存
			relation := rc.NewRelationMapping(l.svcCtx.Mysql, l.svcCtx.Redis)
			// 用户的所有好友
			relationKey1, _ := relation.Key(&relationmodel.Friend{}, "receive_id", map[string]interface{}{
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
		return &pb.DelFriendResp{}, err
	}
	// 预热数据
	go func() {
		ctx := context.Background()
		_, err := NewGetFriendIdsLogic(ctx, l.svcCtx).GetFriendIds(&pb.GetFriendIdsReq{SendUserId: in.SendUserId})
		if err != nil {
			l.Errorf("GetFriend error: %v", err)
		}
	}()
	return &pb.DelFriendResp{}, nil
}
