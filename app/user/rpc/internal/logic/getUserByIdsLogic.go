package logic

import (
	"context"
	"github.com/Path-IM/Path-IM-Server/common/xcache/dc"
	usermodel "github.com/showurl/Path-IM-Server-OICQ/app/user/model"

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
	mapping := dc.GetDbMapping(l.svcCtx.Redis, l.svcCtx.Mysql)
	var users []*usermodel.User
	var userBytes [][]byte
	err := mapping.ListByIds(&usermodel.User{}, &users, in.UserId)
	if err != nil {
		l.Errorf("get user by ids error: %v", err)
		return &pb.GetUserByIdsResp{}, err
	}
	for _, user := range users {
		userBytes = append(userBytes, user.Bytes())
	}
	return &pb.GetUserByIdsResp{Users: userBytes}, nil
}
