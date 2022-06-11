package logic

import (
	"context"
	"github.com/Path-IM/Path-IM-Server/common/xcache/global"
	"github.com/Path-IM/Path-IM-Server/common/xcache/rc"
	groupmodel "github.com/showurl/Path-IM-Server-OICQ/app/group/model"

	"github.com/showurl/Path-IM-Server-OICQ/app/group/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/group/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupIdsLogic {
	return &GetGroupIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupIdsLogic) GetGroupIds(in *pb.GetGroupIdsReq) (*pb.GetGroupIdsResp, error) {
	relation := rc.NewRelationMapping(l.svcCtx.Mysql, l.svcCtx.Redis)
	var groupIds []string
	err := relation.List(&groupIds, 0, 0, &groupmodel.GroupMember{}, "group_id", map[string]interface{}{
		"user_id": in.UserId,
	}, rc.Order("joined_at"))
	if err != nil {
		if global.RedisErrorNotExists == err {
			err = nil
		} else {
			l.Errorf("GetGroupIds error: %s", err)
			return &pb.GetGroupIdsResp{}, err
		}
	}
	return &pb.GetGroupIdsResp{GroupIds: groupIds}, nil
}
