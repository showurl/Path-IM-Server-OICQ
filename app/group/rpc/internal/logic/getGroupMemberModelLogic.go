package logic

import (
	"context"
	groupmodel "github.com/showurl/Path-IM-Server-OICQ/app/group/model"

	"github.com/showurl/Path-IM-Server-OICQ/app/group/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/group/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupMemberModelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupMemberModelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupMemberModelLogic {
	return &GetGroupMemberModelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupMemberModelLogic) GetGroupMemberModel(in *pb.GetGroupMemberModelReq) (*pb.GetGroupMemberModelResp, error) {
	var groupMembers []*groupmodel.GroupMember
	err := l.svcCtx.Mysql.Model(&groupmodel.GroupMember{}).Where("user_id = ? AND group_id in (?)", in.UserId, in.GroupIds).Find(&groupMembers).Error
	if err != nil {
		l.Errorf("GetGroupMemberModel error: %s", err.Error())
		return &pb.GetGroupMemberModelResp{}, err
	}
	var resp [][]byte
	for _, groupMember := range groupMembers {
		resp = append(resp, groupMember.Bytes())
	}
	return &pb.GetGroupMemberModelResp{GroupMembers: resp}, nil
}
