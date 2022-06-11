package tests

import (
	"context"
	"github.com/showurl/Path-IM-Server-OICQ/app/relation/rpc/relationservice"
	"github.com/zeromicro/go-zero/zrpc"
	"testing"
)

var (
	service relationservice.RelationService
	ctx     = context.Background()
)

func init() {
	service = relationservice.NewRelationService(zrpc.MustNewClient(zrpc.RpcClientConf{
		Endpoints: []string{"192.168.1.98:10015"},
	}))
}
func TestAddBlacklist(t *testing.T) {
	resp, err := service.AddBlacklist(ctx, &relationservice.AddBlacklistReq{
		SendUserId:  "bbc64e190c7107574a217d7c7c279c7c",
		RecvUserIds: []string{"3604f15eed0f94282e271ed583427eec", "bbc64e190c7107574a217d7c7c279c7c"},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.String())
}
func TestDelBlacklist(t *testing.T) {
	resp, err := service.DelBlacklist(ctx, &relationservice.DelBlacklistReq{
		SendUserId:  "bbc64e190c7107574a217d7c7c279c7c",
		RecvUserIds: []string{"3604f15eed0f94282e271ed583427eec", "bbc64e190c7107574a217d7c7c279c7c"},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.String())
}

func TestGetBlacklist(t *testing.T) {
	resp, err := service.GetBlacklist(ctx, &relationservice.GetBlacklistReq{
		SendUserId: "bbc64e190c7107574a217d7c7c279c7c",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.String())
	resp2, err := service.GetBlacklistModel(ctx, &relationservice.GetBlacklistModelReq{
		SendUserId:   "bbc64e190c7107574a217d7c7c279c7c",
		BlacklistIds: resp.BlacklistIds,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp2.String())
	resp3, err := service.IsBlacklist(ctx, &relationservice.IsBlacklistReq{
		SendUserId:  "bbc64e190c7107574a217d7c7c279c7c",
		RecvUserIds: append(resp.BlacklistIds, "123"),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp3.String())
}
func TestAddFriend(t *testing.T) {
	resp, err := service.AddFriend(ctx, &relationservice.AddFriendReq{
		SendUserId:  "bbc64e190c7107574a217d7c7c279c7c",
		RecvUserIds: []string{"3604f15eed0f94282e271ed583427eec", "bbc64e190c7107574a217d7c7c279c7c"},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.String())
}
func TestDelFriend(t *testing.T) {
	resp, err := service.DelFriend(ctx, &relationservice.DelFriendReq{
		SendUserId:  "bbc64e190c7107574a217d7c7c279c7c",
		RecvUserIds: []string{"3604f15eed0f94282e271ed583427eec", "bbc64e190c7107574a217d7c7c279c7c"},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.String())
}
func TestUpdateFriendModel(t *testing.T) {
	resp, err := service.UpdateFriendModel(ctx, &relationservice.UpdateFriendModelReq{
		SendUserId: "bbc64e190c7107574a217d7c7c279c7c",
		RecvUserId: "3604f15eed0f94282e271ed583427eec",
		UpdateMap: map[string]string{
			"remark":  "这是用户2的备注",
			"msg_opt": "1",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.String())
}
func TestGetFriend(t *testing.T) {
	resp, err := service.GetFriendIds(ctx, &relationservice.GetFriendIdsReq{SendUserId: "bbc64e190c7107574a217d7c7c279c7c"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.String())
	resp1, err := service.GetFriendModel(ctx, &relationservice.GetFriendModelReq{
		SendUserId:  "bbc64e190c7107574a217d7c7c279c7c",
		RecvUserIds: resp.FriendIds,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp1.String())
	resp2, err := service.IsFriend(ctx, &relationservice.IsFriendReq{
		SendUserId:  "bbc64e190c7107574a217d7c7c279c7c",
		RecvUserIds: append(resp.FriendIds, "123"),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp2.String())
}
