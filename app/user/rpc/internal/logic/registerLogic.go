package logic

import (
	"context"
	"github.com/Path-IM/Path-IM-Server/common/xerr"
	"github.com/Path-IM/Path-IM-Server/common/xorm"
	"github.com/showurl/Path-IM-Server-OICQ/app/user/model"
	"gorm.io/gorm"

	"github.com/showurl/Path-IM-Server-OICQ/app/user/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	user := &usermodel.User{
		Username: in.Username,
		Password: in.Password,
	}
	err := xorm.Transaction(l.svcCtx.Mysql,
		// 插入用户
		user.FuncInsert(l.svcCtx.Redis),
		// 清除用户相关的缓存
		func(tx *gorm.DB) error {
			err := user.FlushCache(l.svcCtx.Mysql, l.svcCtx.Redis)
			return err
		})
	if err != nil {
		return &pb.RegisterResp{
			FailReason: "操作数据库错误：" + err.Error(),
			UserId:     "",
		}, xerr.New(500, err.Error())
	}
	// 预热数据
	go user.Preheat(l.svcCtx.Mysql, l.svcCtx.Redis)
	return &pb.RegisterResp{
		FailReason: "",
		UserId:     user.Id,
	}, nil
}
