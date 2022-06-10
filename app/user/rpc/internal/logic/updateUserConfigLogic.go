package logic

import (
	"context"

	"github.com/showurl/Path-IM-Server-OICQ/app/user/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserConfigLogic {
	return &UpdateUserConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserConfigLogic) UpdateUserConfig(in *pb.UpdateUserConfigReq) (*pb.UpdateUserConfigResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateUserConfigResp{}, nil
}
