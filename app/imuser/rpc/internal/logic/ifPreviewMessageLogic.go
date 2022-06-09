package logic

import (
	"context"

	"github.com/showurl/Path-IM-Server-OICQ/app/imuser/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/imuser/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IfPreviewMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIfPreviewMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IfPreviewMessageLogic {
	return &IfPreviewMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  是否预览消息
func (l *IfPreviewMessageLogic) IfPreviewMessage(in *pb.IfPreviewMessageReq) (*pb.IfPreviewMessageResp, error) {
	// todo: add your logic here and delete this line

	return &pb.IfPreviewMessageResp{}, nil
}
