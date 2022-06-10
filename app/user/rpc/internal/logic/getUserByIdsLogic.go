package logic

import (
	"context"

	"github.com/showurl/Path-IM-Server-OICQ/app/user/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdsLogic {
	return &GetUserByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdsLogic) GetUserByIds(in *pb.GetUserByIdsReq) (*pb.GetUserByIdsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserByIdsResp{}, nil
}
