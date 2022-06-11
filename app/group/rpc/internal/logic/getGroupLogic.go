package logic

import (
	"context"
	"github.com/Path-IM/Path-IM-Server/common/xcache/dc"
	groupmodel "github.com/showurl/Path-IM-Server-OICQ/app/group/model"

	"github.com/showurl/Path-IM-Server-OICQ/app/group/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/group/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupLogic {
	return &GetGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupLogic) GetGroup(in *pb.GetGroupReq) (*pb.GetGroupResp, error) {
	mapping := dc.GetDbMapping(l.svcCtx.Redis, l.svcCtx.Mysql)
	group := &groupmodel.Group{}
	var groups []*groupmodel.Group
	err := mapping.ListByIds(group, &groups, in.GroupIds)
	if err != nil {
		l.Errorf("get group failed, err: %v", err)
		return &pb.GetGroupResp{}, err
	}
	var resp [][]byte
	for _, g := range groups {
		resp = append(resp, g.Bytes())
	}
	return &pb.GetGroupResp{Groups: resp}, nil
}
