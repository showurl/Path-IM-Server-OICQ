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

type GetGroupMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupMemberLogic {
	return &GetGroupMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupMemberLogic) GetGroupMember(in *pb.GetGroupMemberReq) (*pb.GetGroupMemberResp, error) {
	relation := rc.NewRelationMapping(l.svcCtx.Mysql, l.svcCtx.Redis)
	var userIds []string
	err := relation.List(&userIds, 0, 0, &groupmodel.GroupMember{}, "user_id", map[string]interface{}{
		"group_id": in.GroupId,
	}, rc.Order("joined_at"))
	if err != nil {
		if global.RedisErrorNotExists == err {
			err = nil
		} else {
			l.Errorf("GetGroupMember error: %s", err.Error())
			return &pb.GetGroupMemberResp{}, err
		}
	}
	return &pb.GetGroupMemberResp{Members: userIds}, nil
}
