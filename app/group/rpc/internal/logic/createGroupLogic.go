package logic

import (
	"context"
	"errors"
	"github.com/Path-IM/Path-IM-Server/common/xcache/dc"
	"github.com/Path-IM/Path-IM-Server/common/xcache/rc"
	"github.com/Path-IM/Path-IM-Server/common/xorm"
	"github.com/Path-IM/Path-IM-Server/common/xorm/global"
	groupmodel "github.com/showurl/Path-IM-Server-OICQ/app/group/model"
	"gorm.io/gorm"
	"time"

	"github.com/showurl/Path-IM-Server-OICQ/app/group/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/group/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateGroupLogic) CreateGroup(in *pb.CreateGroupReq) (*pb.CreateGroupResp, error) {
	if len(in.Members) == 0 {
		return &pb.CreateGroupResp{}, errors.New("members is empty")
	}
	group := &groupmodel.Group{
		Id:   global.GetID(),
		Name: in.Name,
	}
	var groupMembers []*groupmodel.GroupMember
	for _, member := range in.Members {
		groupMembers = append(groupMembers, &groupmodel.GroupMember{
			GroupID:  group.Id,
			UserID:   member,
			JoinedAt: time.Now().UnixMilli(),
			MsgOpt:   0,
			Remark:   "",
		})
	}
	err := xorm.Transaction(l.svcCtx.Mysql,
		group.FuncInsert(l.svcCtx.Redis),
		func(tx *gorm.DB) error {
			// 添加群成员
			err := tx.Model(&groupmodel.GroupMember{}).Create(groupMembers).Error
			if err != nil {
				l.Errorf("create group member error: %v", err)
				return err
			}
			return nil
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
			// 用户的所有群
			for _, member := range in.Members {
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
		return &pb.CreateGroupResp{}, err
	}
	// 预热数据
	go func() {
		ctx := context.Background()
		go func() {
			_, err := NewGetGroupLogic(ctx, l.svcCtx).GetGroup(&pb.GetGroupReq{GroupIds: []string{group.Id}})
			if err != nil {
				l.Errorf("get group error: %v", err)
			}
		}()
		go func() {
			_, err := NewGetGroupMemberLogic(ctx, l.svcCtx).GetGroupMember(&pb.GetGroupMemberReq{GroupId: group.Id})
			if err != nil {
				l.Errorf("get group member error: %v", err)
			}
		}()
		{
			for _, member := range in.Members {
				go func(member string) {
					_, err := NewGetGroupIdsLogic(ctx, l.svcCtx).GetGroupIds(&pb.GetGroupIdsReq{UserId: member})
					if err != nil {
						l.Errorf("get group ids error: %v", err)
					}
				}(member)
			}
		}
	}()
	return &pb.CreateGroupResp{GroupId: group.Id}, nil
}
