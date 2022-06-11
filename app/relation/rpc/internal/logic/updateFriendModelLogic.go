package logic

import (
	"context"
	"errors"
	relationmodel "github.com/showurl/Path-IM-Server-OICQ/app/relation/model"
	"strconv"

	"github.com/showurl/Path-IM-Server-OICQ/app/relation/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFriendModelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFriendModelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFriendModelLogic {
	return &UpdateFriendModelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateFriendModelLogic) UpdateFriendModel(in *pb.UpdateFriendModelReq) (*pb.UpdateFriendModelResp, error) {
	var updateMap = make(map[string]interface{})
	for k, v := range in.UpdateMap {
		switch k {
		case "name", "remark":
			updateMap[k] = v
		case "msg_opt":
			opt, _ := strconv.Atoi(v)
			updateMap["msg_opt"] = opt
		default:
			return &pb.UpdateFriendModelResp{
				FailedReason: "invalid field",
			}, errors.New("invalid field")
		}
	}
	err := l.svcCtx.Mysql.Model(&relationmodel.Friend{}).Where("send_id = ? AND receive_id = ?", in.SendUserId, in.RecvUserId).
		Updates(updateMap).Error
	if err != nil {
		l.Errorf("update friend model failed, err: %v", err)
		return &pb.UpdateFriendModelResp{
			FailedReason: "update group member model failed",
		}, err
	}
	return &pb.UpdateFriendModelResp{}, nil
}
