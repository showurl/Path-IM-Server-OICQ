package logic

import (
	"context"
	relationmodel "github.com/showurl/Path-IM-Server-OICQ/app/relation/model"

	"github.com/showurl/Path-IM-Server-OICQ/app/relation/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendModelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendModelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendModelLogic {
	return &GetFriendModelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendModelLogic) GetFriendModel(in *pb.GetFriendModelReq) (*pb.GetFriendModelResp, error) {
	var friends []*relationmodel.Friend
	err := l.svcCtx.Mysql.Model(&relationmodel.Friend{}).Where("send_id = ? AND receive_id in (?)", in.SendUserId, in.RecvUserIds).Find(&friends).Error
	if err != nil {
		l.Errorf("get friend error: %v", err)
		return &pb.GetFriendModelResp{}, err
	}
	var resp [][]byte
	for _, friend := range friends {
		resp = append(resp, friend.Bytes())
	}
	return &pb.GetFriendModelResp{Models: resp}, nil
}
