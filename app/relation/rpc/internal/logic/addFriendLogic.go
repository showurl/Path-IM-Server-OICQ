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

type AddFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFriendLogic {
	return &AddFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddFriendLogic) AddFriend(in *pb.AddFriendReq) (*pb.AddFriendResp, error) {
	createdAt := time.Now().UnixMilli()
	var friends []*relationmodel.Friend
	for _, id := range in.RecvUserIds {
		friends = append(friends, &relationmodel.Friend{
			SendId:    in.SendUserId,
			ReceiveId: id,
			CreatedAt: createdAt,
			MsgOpt:    0,
			Remark:    "",
		})
	}
	err := xorm.Transaction(l.svcCtx.Mysql,
		// 添加关系
		func(tx *gorm.DB) error {
			err := tx.Model(&relationmodel.Friend{}).Create(friends).Error
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
		if xormerr.DuplicateError(err) {
			return &pb.AddFriendResp{FailedReason: "已经成为好友了"}, nil
		}
		return &pb.AddFriendResp{}, err
	}
	// 预热数据
	go func() {
		ctx := context.Background()
		_, err := NewGetFriendIdsLogic(ctx, l.svcCtx).GetFriendIds(&pb.GetFriendIdsReq{SendUserId: in.SendUserId})
		if err != nil {
			l.Errorf("GetFriendIds error: %v", err)
		}
	}()
	return &pb.AddFriendResp{}, nil
}
