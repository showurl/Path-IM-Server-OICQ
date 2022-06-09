package logic

import (
	"context"

	"github.com/showurl/Path-IM-Server-OICQ/app/imuser/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/imuser/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSingleConversationRecvMsgOptsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSingleConversationRecvMsgOptsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSingleConversationRecvMsgOptsLogic {
	return &GetSingleConversationRecvMsgOptsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  获取单聊会话的消息接收选项
func (l *GetSingleConversationRecvMsgOptsLogic) GetSingleConversationRecvMsgOpts(in *pb.GetSingleConversationRecvMsgOptsReq) (*pb.GetSingleConversationRecvMsgOptsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetSingleConversationRecvMsgOptsResp{}, nil
}
