package logic

import (
	"context"

	"github.com/showurl/Path-IM-Server-OICQ/app/relation/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFriendLogic {
	return &IsFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsFriendLogic) IsFriend(in *pb.IsFriendReq) (*pb.IsFriendResp, error) {
	idsResp, err := NewGetFriendIdsLogic(l.ctx, l.svcCtx).GetFriendIds(&pb.GetFriendIdsReq{SendUserId: in.SendUserId})
	resp := &pb.IsFriendResp{}
	if err != nil {
		l.Errorf("GetFriendIds error: %s", err)
		return resp, err
	}
	friendMap := make(map[string]interface{})
	for _, friend := range idsResp.FriendIds {
		friendMap[friend] = nil
	}
	for _, id := range in.RecvUserIds {
		_, ok := friendMap[id]
		resp.List = append(resp.List, &pb.IsFriendResp_Item{
			UserId:   id,
			IsFriend: ok,
		})
	}
	return resp, nil
}
