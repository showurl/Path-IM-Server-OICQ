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

type GetBlacklistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBlacklistLogic {
	return &GetBlacklistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBlacklistLogic) GetBlacklist(in *pb.GetBlacklistReq) (*pb.GetBlacklistResp, error) {
	relation := rc.NewRelationMapping(l.svcCtx.Mysql, l.svcCtx.Redis)
	var receiveIds []string
	err := relation.List(&receiveIds, 0, 0, &relationmodel.Blacklist{}, "receive_id", map[string]interface{}{
		"send_id": in.SendUserId,
	}, rc.Order("created_at"))
	if err != nil {
		if global.RedisErrorNotExists == err {
			err = nil
		} else {
			l.Errorf("get blacklist error: %v", err)
			return &pb.GetBlacklistResp{}, err
		}
	}
	return &pb.GetBlacklistResp{BlacklistIds: receiveIds}, nil
}
