package logic

import (
	"context"
	msggatewaypb "github.com/Path-IM/Path-IM-Server/app/msg-gateway/cmd/wsrpc/pb"
	"github.com/Path-IM/Path-IM-Server/common/xcache/global"
	"github.com/golang-jwt/jwt/v4"

	jwtUtils "github.com/Path-IM/Path-IM-Server-Demo/common/utils/jwt"
	"github.com/showurl/Path-IM-Server-OICQ/app/user/rpc/internal/svc"
	"github.com/showurl/Path-IM-Server-OICQ/app/user/rpc/pb"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenLogic {
	return &GetTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTokenLogic) GetToken(in *pb.GetTokenReq) (*pb.GetTokenResp, error) {
	claim := jwtUtils.BuildClaims(in.UserId, in.Platform)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedString, err := token.SignedString([]byte(l.svcCtx.Config.TokenSecret))
	if err != nil {
		l.Errorf("jwt signed string error: %v", err)
		return nil, err
	}
	redisKey := global.MergeKey("token", in.Platform, in.UserId)
	// 踢出原来的token和socket连接
	tokenMap, err := l.svcCtx.Redis.HGetAll(l.ctx, redisKey).Result()
	if err != nil {
		l.Errorf("redis hgetall error: %v", err)
		return nil, err
	}
	if len(tokenMap) > 0 {
		for token, status := range tokenMap {
			switch status {
			case "0":
				// 改为过期
				err = l.svcCtx.Redis.HSet(l.ctx, redisKey, token, "1").Err()
				if err != nil {
					l.Errorf("redis hset error: %v", err)
					return nil, err
				}
				// 断开连接
				_, err = l.svcCtx.WsLogic().KickUserConns(l.ctx, &msggatewaypb.KickUserConnsReq{
					UserID:      in.UserId,
					PlatformIDs: []string{in.Platform},
				})
				if err != nil {
					l.Errorf("kick user conns error: %v", err)
					return nil, err
				}
			case "1":
				// 删除
				e := l.svcCtx.Redis.HDel(l.ctx, redisKey, token).Err()
				if e != nil {
					l.Errorf("redis hdel error: %v", e)
				}
			}
		}
	}
	// 存储新的token
	err = l.svcCtx.Redis.HSet(l.ctx, redisKey, signedString, "0").Err()
	if err != nil {
		l.Errorf("redis hset error: %v", err)
		return nil, err
	}
	return &pb.GetTokenResp{
		Token: signedString,
	}, nil
}
