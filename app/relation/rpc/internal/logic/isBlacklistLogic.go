package logic

import (
	"context"

	"github.com/showurl/Path-IM-Server-OICQ/app/relation/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsBlacklistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsBlacklistLogic {
	return &IsBlacklistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsBlacklistLogic) IsBlacklist(in *pb.IsBlacklistReq) (*pb.IsBlacklistResp, error) {
	idsResp, err := NewGetBlacklistLogic(l.ctx, l.svcCtx).GetBlacklist(&pb.GetBlacklistReq{SendUserId: in.SendUserId})
	resp := &pb.IsBlacklistResp{}
	if err != nil {
		l.Errorf("GetFriendIds error: %s", err)
		return resp, err
	}
	blackMap := make(map[string]interface{})
	for _, friend := range idsResp.BlacklistIds {
		blackMap[friend] = nil
	}
	for _, id := range in.RecvUserIds {
		_, ok := blackMap[id]
		resp.List = append(resp.List, &pb.IsBlacklistResp_Item{
			UserId:      id,
			IsBlacklist: ok,
		})
	}
	return resp, nil
}
