package logic

import (
	"context"
	"errors"
	groupmodel "github.com/showurl/Path-IM-Server-OICQ/app/group/model"
	"strconv"

	"github.com/showurl/Path-IM-Server-OICQ/app/group/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/group/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupMemberModelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGroupMemberModelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupMemberModelLogic {
	return &UpdateGroupMemberModelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGroupMemberModelLogic) UpdateGroupMemberModel(in *pb.UpdateGroupMemberModelReq) (*pb.UpdateGroupMemberModelResp, error) {
	var updateMap = make(map[string]interface{})
	for k, v := range in.UpdateMap {
		switch k {
		case "name", "remark":
			updateMap[k] = v
		case "msg_opt":
			opt, _ := strconv.Atoi(v)
			updateMap["msg_opt"] = opt
		default:
			return &pb.UpdateGroupMemberModelResp{
				FailedReason: "invalid field",
			}, errors.New("invalid field")
		}
	}
	err := l.svcCtx.Mysql.Model(&groupmodel.GroupMember{}).Where("group_id = ? AND user_id = ?", in.GroupId, in.UserId).
		Updates(updateMap).Error
	if err != nil {
		l.Errorf("update group member model failed, err: %v", err)
		return &pb.UpdateGroupMemberModelResp{
			FailedReason: "update group member model failed",
		}, err
	}
	return &pb.UpdateGroupMemberModelResp{}, nil
}
