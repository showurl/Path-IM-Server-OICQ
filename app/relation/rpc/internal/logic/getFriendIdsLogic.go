package logic

import (
	"context"
	"github.com/Path-IM/Path-IM-Server/common/xcache/global"
	"github.com/Path-IM/Path-IM-Server/common/xcache/rc"
	relationmodel "github.com/showurl/Path-IM-Server-OICQ/app/relation/model"

	"github.com/showurl/Path-IM-Server-OICQ/app/relation/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendIdsLogic {
	return &GetFriendIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendIdsLogic) GetFriendIds(in *pb.GetFriendIdsReq) (*pb.GetFriendIdsResp, error) {
	relation := rc.NewRelationMapping(l.svcCtx.Mysql, l.svcCtx.Redis)
	var receiveIds []string
	err := relation.List(&receiveIds, 0, 0, &relationmodel.Friend{}, "receive_id", map[string]interface{}{
		"send_id": in.SendUserId,
	}, rc.Order("created_at"))
	if err != nil {
		if global.RedisErrorNotExists == err {
			err = nil
		} else {
			l.Errorf("get blacklist error: %v", err)
			return &pb.GetFriendIdsResp{}, err
		}
	}
	return &pb.GetFriendIdsResp{FriendIds: receiveIds}, nil
}
