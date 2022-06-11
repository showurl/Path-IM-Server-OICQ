package logic

import (
	"context"
	"github.com/Path-IM/Path-IM-Server/common/xcache/rc"
	"github.com/Path-IM/Path-IM-Server/common/xorm"
	groupmodel "github.com/showurl/Path-IM-Server-OICQ/app/group/model"
	"github.com/showurl/Path-IM-Server-OICQ/app/group/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/group/rpc/pb"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGroupMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteGroupMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGroupMemberLogic {
	return &DeleteGroupMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteGroupMemberLogic) DeleteGroupMember(in *pb.DeleteGroupMemberReq) (*pb.DeleteGroupMemberResp, error) {
	err := xorm.Transaction(l.svcCtx.Mysql,
		// 删除群成员
		func(tx *gorm.DB) error {
			err := tx.Where("group_id = ? AND user_id = ?", in.GroupId, in.Member).Delete(&groupmodel.GroupMember{}).Error
			if err != nil {
				l.Errorf("delete group member error: %v", err)
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
		return &pb.DeleteGroupMemberResp{}, err
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
	return &pb.DeleteGroupMemberResp{}, nil
}
