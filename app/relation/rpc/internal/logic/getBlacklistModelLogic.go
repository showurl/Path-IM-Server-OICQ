package logic

import (
	"context"
	relationmodel "github.com/showurl/Path-IM-Server-OICQ/app/relation/model"

	"github.com/showurl/Path-IM-Server-OICQ/app/relation/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBlacklistModelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBlacklistModelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBlacklistModelLogic {
	return &GetBlacklistModelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBlacklistModelLogic) GetBlacklistModel(in *pb.GetBlacklistModelReq) (*pb.GetBlacklistModelResp, error) {
	var blacklists []*relationmodel.Blacklist
	err := l.svcCtx.Mysql.Model(&relationmodel.Blacklist{}).Where("send_id = ? AND receive_id in (?)", in.SendUserId, in.BlacklistIds).Find(&blacklists).Error
	if err != nil {
		l.Errorf("get blacklist error: %v", err)
		return &pb.GetBlacklistModelResp{}, err
	}
	var resp [][]byte
	for _, blacklist := range blacklists {
		resp = append(resp, blacklist.Bytes())
	}
	return &pb.GetBlacklistModelResp{Models: resp}, nil
}
