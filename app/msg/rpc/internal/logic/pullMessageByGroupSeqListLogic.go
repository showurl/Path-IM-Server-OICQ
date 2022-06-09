package logic

import (
	"context"
	"github.com/Path-IM/Path-IM-Server/common/types"
	"github.com/showurl/Path-IM-Server-OICQ/app/msg/rpc/internal/repository"

	"github.com/showurl/Path-IM-Server-OICQ/app/msg/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/msg/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PullMessageByGroupSeqListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewPullMessageByGroupSeqListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PullMessageByGroupSeqListLogic {
	return &PullMessageByGroupSeqListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

func (l *PullMessageByGroupSeqListLogic) PullMessageByGroupSeqList(in *pb.PullMsgByGroupSeqListReq) (*pb.PullMsgListResp, error) {
	resp := new(pb.PullMsgListResp)
	msgList, err := l.rep.GetMsgByGroupSeqList(in.GroupID, in.SeqList)
	if err != nil {
		l.Error("PullMessageBySeqList data error ", err.Error())
		resp.ErrCode = types.ErrCodeFailed
		resp.ErrMsg = err.Error()
		return resp, nil
	}
	resp.ErrCode = types.ErrCodeOK
	resp.ErrMsg = ""
	resp.List = msgList
	return resp, nil
}
