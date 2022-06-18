package imuser

import (
	"context"
	"fmt"
	chatpb "github.com/Path-IM/Path-IM-Server/app/msg/cmd/rpc/pb"
	"github.com/Path-IM/Path-IM-Server/common/xerr"
	"github.com/Path-IM/Path-IM-Server/common/xorm/global"
	"github.com/showurl/Path-IM-Server-OICQ/app/api/internal/model"
	"github.com/showurl/Path-IM-Server-OICQ/app/api/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/api/internal/types"
	"gorm.io/gorm"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	user := &model.User{}
	err = l.svcCtx.DB().Model(user).Where("username = ?", req.Username).First(user).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, xerr.New(500, err.Error())
		}
		// 插入用户
		user.Username = req.Username
		user.Password = req.Password
		err = l.svcCtx.DB().Create(user).Error
		if err != nil {
			return nil, xerr.New(500, err.Error())
		}
	} else {
		if user.Password != req.Password {
			return nil, xerr.New(500, "密码错误")
		}
	}
	resp = &types.LoginResp{
		Token: fmt.Sprintf(`{
"id": "%s"
}`, user.Username),
		Uid: user.Username,
	}
	// 群里发消息
	go func() {
		// 发送群消息
		for {
			_, err := l.svcCtx.MsgRpc().SendMsg(context.Background(), &chatpb.SendMsgReq{MsgData: &chatpb.MsgData{
				ClientMsgID:      global.GetID(),
				ServerMsgID:      "",
				ConversationType: 2,
				SendID:           user.Username,
				ReceiveID:        "default_group",
				ContentType:      1001,
				Content:          []byte("hello 我来了"),
				AtUserIDList:     nil,
				ClientTime:       time.Now().UnixMilli(),
				ServerTime:       0,
				Seq:              0,
				OfflinePush: &chatpb.OfflinePush{
					Title:         "你收到一条消息",
					Desc:          "",
					Ex:            "",
					IOSPushSound:  "",
					IOSBadgeCount: false,
				},
				MsgOptions: &chatpb.MsgOptions{
					Persistent:         true,
					History:            true,
					Local:              true,
					UpdateUnreadCount:  true,
					UpdateConversation: true,
					NeedBeFriend:       false,
					OfflinePush:        true,
					SenderSync:         true,
				},
			}})
			if err == nil {
				break
			}
			logx.Errorf("send msg error: %s", err.Error())
		}
	}()
	return
}
