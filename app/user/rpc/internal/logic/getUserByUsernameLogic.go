package logic

import (
	"context"
	xormerr "github.com/Path-IM/Path-IM-Server/common/xorm/err"
	usermodel "github.com/showurl/Path-IM-Server-OICQ/app/user/model"

	"github.com/showurl/Path-IM-Server-OICQ/app/user/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByUsernameLogic {
	return &GetUserByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByUsernameLogic) GetUserByUsername(in *pb.GetUserByUsernameReq) (*pb.GetUserByUsernameResp, error) {
	user := &usermodel.User{}
	err := l.svcCtx.Mysql.Model(user).Where("username = ?", in.Username).First(user).Error
	if err != nil {
		if xormerr.RecordNotFound(err) {
			err = nil
		} else {
			l.Errorf("mysql first error: %v", err)
			return nil, err
		}
	}
	return &pb.GetUserByUsernameResp{
		FailReason: "",
		User:       user.Bytes(),
	}, nil
}
