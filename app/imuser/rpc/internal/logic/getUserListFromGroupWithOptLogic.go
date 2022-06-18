package logic

import (
	"context"
	"github.com/showurl/Path-IM-Server-OICQ/app/imuser/rpc/internal/model"
	"github.com/showurl/Path-IM-Server-OICQ/app/imuser/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/imuser/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListFromGroupWithOptLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListFromGroupWithOptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListFromGroupWithOptLogic {
	return &GetUserListFromGroupWithOptLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  获取群成员列表 通过消息接收选项
func (l *GetUserListFromGroupWithOptLogic) GetUserListFromGroupWithOpt(in *pb.GetUserListFromGroupWithOptReq) (*pb.GetUserListFromGroupWithOptResp, error) {
	var uids []string
	l.svcCtx.DB().Model(&model.User{}).Pluck("username", &uids)
	var uidList []*pb.UserIDOpt
	for _, uid := range uids {
		uidList = append(uidList, &pb.UserIDOpt{
			UserID: uid,
			Opts:   pb.RecvMsgOpt_ReceiveMessage,
		})
	}
	return &pb.GetUserListFromGroupWithOptResp{
		CommonResp:    &pb.CommonResp{},
		UserIDOptList: uidList,
	}, nil
}
