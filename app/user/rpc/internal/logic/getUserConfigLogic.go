package logic

import (
	"context"

	"github.com/showurl/Path-IM-Server-OICQ/app/user/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserConfigLogic {
	return &GetUserConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserConfigLogic) GetUserConfig(in *pb.GetUserConfigReq) (*pb.GetUserConfigResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserConfigResp{}, nil
}
