package logic

import (
	"context"
	"strconv"

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
	// todo: add your logic here and delete this line
	var uidList []*pb.UserIDOpt
	for i := 1; i <= 2000; i++ {
		uidList = append(uidList, &pb.UserIDOpt{
			UserID: strconv.Itoa(i),
			Opts:   pb.RecvMsgOpt_ReceiveMessage,
		})
	}
	for i := 2001; i <= 5000; i++ {
		uidList = append(uidList, &pb.UserIDOpt{
			UserID: strconv.Itoa(i),
			Opts:   pb.RecvMsgOpt_ReceiveNotNotifyMessage,
		})
	}
	for i := 5001; i <= 10000; i++ {
		uidList = append(uidList, &pb.UserIDOpt{
			UserID: strconv.Itoa(i),
			Opts:   pb.RecvMsgOpt_NotReceiveMessage,
		})
	}
	return &pb.GetUserListFromGroupWithOptResp{
		CommonResp:    &pb.CommonResp{},
		UserIDOptList: uidList,
	}, nil
}
