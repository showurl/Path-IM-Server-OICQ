package logic

import (
	"context"
	"github.com/Path-IM/Path-IM-Server/common/xcache/rc"
	"github.com/Path-IM/Path-IM-Server/common/xorm"
	xormerr "github.com/Path-IM/Path-IM-Server/common/xorm/err"
	groupmodel "github.com/showurl/Path-IM-Server-OICQ/app/group/model"
	"gorm.io/gorm"
	"time"

	"github.com/showurl/Path-IM-Server-OICQ/app/group/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/group/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddGroupMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddGroupMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGroupMemberLogic {
	return &AddGroupMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddGroupMemberLogic) AddGroupMember(in *pb.AddGroupMemberReq) (*pb.AddGroupMemberResp, error) {
	var groupMembers []*groupmodel.GroupMember
	groupMembers = append(groupMembers, &groupmodel.GroupMember{
		GroupID:  in.GroupId,
		UserID:   in.Member,
		JoinedAt: time.Now().UnixMilli(),
		MsgOpt:   0,
		Remark:   "",
	})
	err := xorm.Transaction(l.svcCtx.Mysql,
		// 添加群成员
		func(tx *gorm.DB) error {
			err := tx.Model(&groupmodel.GroupMember{}).Create(groupMembers).Error
			if err != nil {
				l.Errorf("create group member error: %v", err)
				return err
			}
			return nil
		},
		// 清理群缓存
		func(tx *gorm.DB) error {
			var keys []string
			// 删除缓存
			relation := rc.NewRelationMapping(l.svcCtx.Mysql, l.svcCtx.Redis)
			// 用户的所有群
			relationKey1, _ := relation.Key(&groupmodel.GroupMember{}, "group_id", map[string]interface{}{
				"user_id": in.Member,
			}, rc.Order("joined_at"))
			keys = append(keys, relationKey1)
			// 群里的所有用户
			relationKey2, _ := relation.Key(&groupmodel.GroupMember{}, "user_id", map[string]interface{}{
				"group_id": in.GroupId,
			}, rc.Order("joined_at"))
			keys = append(keys, relationKey2)
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
		if xormerr.DuplicateError(err) {
			return &pb.AddGroupMemberResp{FailedReason: "用户已经在群里了"}, nil
		}
		return &pb.AddGroupMemberResp{}, err
	}
	// 预热数据
	go func() {
		ctx := context.Background()
		go func() {
			_, err := NewGetGroupMemberLogic(ctx, l.svcCtx).GetGroupMember(&pb.GetGroupMemberReq{GroupId: in.GroupId})
			if err != nil {
				l.Errorf("get group member error: %v", err)
			}
		}()
		go func() {
			_, err := NewGetGroupIdsLogic(ctx, l.svcCtx).GetGroupIds(&pb.GetGroupIdsReq{UserId: in.Member})
			if err != nil {
				l.Errorf("get group ids error: %v", err)
			}
		}()
	}()
	return &pb.AddGroupMemberResp{}, nil
}
