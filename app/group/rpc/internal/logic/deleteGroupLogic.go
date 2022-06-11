package logic

import (
	"context"
	"github.com/Path-IM/Path-IM-Server/common/xcache/dc"
	"github.com/Path-IM/Path-IM-Server/common/xcache/rc"
	"github.com/Path-IM/Path-IM-Server/common/xorm"
	groupmodel "github.com/showurl/Path-IM-Server-OICQ/app/group/model"
	"github.com/showurl/Path-IM-Server-OICQ/app/group/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/group/rpc/pb"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGroupLogic {
	return &DeleteGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteGroupLogic) DeleteGroup(in *pb.DeleteGroupReq) (*pb.DeleteGroupResp, error) {
	group := &groupmodel.Group{
		Id: in.GroupId,
	}
	// 群成员
	resp, err := NewGetGroupMemberLogic(l.ctx, l.svcCtx).GetGroupMember(&pb.GetGroupMemberReq{GroupId: in.GetGroupId()})
	if err != nil {
		l.Errorf("get group member error: %v", err)
		return &pb.DeleteGroupResp{}, err
	}
	err = xorm.Transaction(l.svcCtx.Mysql,
		func(tx *gorm.DB) error {
			return tx.Model(group).Where("id = ?", group.Id).Delete(&groupmodel.Group{}).Error
		},
		func(tx *gorm.DB) error {
			return tx.Where("group_id = ?", group.Id).Delete(&groupmodel.GroupMember{}).Error
		},
		func(tx *gorm.DB) error {
			var keys []string
			// 删除缓存
			mapping := dc.GetDbMapping(l.svcCtx.Redis, l.svcCtx.Mysql)
			relation := rc.NewRelationMapping(l.svcCtx.Mysql, l.svcCtx.Redis)
			// 群组的
			keys = append(keys, mapping.Key(&groupmodel.Group{}, map[string][]interface{}{
				"id": {group.Id},
			}))
			for _, member := range resp.Members {
				relationKey, _ := relation.Key(&groupmodel.GroupMember{}, "group_id", map[string]interface{}{
					"user_id": member,
				}, rc.Order("joined_at"))
				keys = append(keys, relationKey)
			}
			// 群里的所有用户
			relationKey, _ := relation.Key(&groupmodel.GroupMember{}, "user_id", map[string]interface{}{
				"group_id": group.Id,
			}, rc.Order("joined_at"))
			keys = append(keys, relationKey)
			// 删除缓存
			err := l.svcCtx.Redis.Del(l.ctx, keys...).Err()
			if err != nil {
				l.Errorf("delete cache error: %v", err)
				return err
			}
			return nil
		},
	)
	if err != nil {
		return &pb.DeleteGroupResp{}, err
	}
	return &pb.DeleteGroupResp{}, nil
}
