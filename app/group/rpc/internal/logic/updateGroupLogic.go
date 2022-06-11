package logic

import (
	"context"
	"github.com/Path-IM/Path-IM-Server/common/xcache/dc"
	"github.com/Path-IM/Path-IM-Server/common/xorm"
	groupmodel "github.com/showurl/Path-IM-Server-OICQ/app/group/model"
	"gorm.io/gorm"

	"github.com/showurl/Path-IM-Server-OICQ/app/group/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/group/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupLogic {
	return &UpdateGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGroupLogic) UpdateGroup(in *pb.UpdateGroupReq) (*pb.UpdateGroupResp, error) {
	group := &groupmodel.Group{
		Id: in.GroupId,
	}
	err := xorm.Transaction(l.svcCtx.Mysql,
		func(tx *gorm.DB) error {
			return tx.Model(group).Where("id = ?", group.Id).Update("name", in.Name).Error
		},
		func(tx *gorm.DB) error {
			var keys []string
			// 删除缓存
			mapping := dc.GetDbMapping(l.svcCtx.Redis, l.svcCtx.Mysql)
			// 群组的
			keys = append(keys, mapping.Key(&groupmodel.Group{}, map[string][]interface{}{
				"id": {group.Id},
			}))
			return nil
		},
	)
	if err != nil {
		return &pb.UpdateGroupResp{}, err
	}
	return &pb.UpdateGroupResp{}, nil
}
