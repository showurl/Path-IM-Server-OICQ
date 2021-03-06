# 上一章：[群组](group-rpc.md)

---

# 关系模块

## proto

```protobuf
syntax = "proto3";

option go_package = "./pb";

package relation;

message AddFriendReq {
  string SendUserId = 1;
  repeated string RecvUserIds = 2;
}
message AddFriendResp {
  string FailedReason = 1;
}
message DelFriendReq {
  string SendUserId = 1;
  repeated string RecvUserIds = 2;
}
message DelFriendResp {
  string FailedReason = 1;
}
message IsFriendReq {
  string SendUserId = 1;
  repeated string RecvUserIds = 2;
}
message IsFriendResp {
  message Item {
    string UserId = 1;
    bool IsFriend = 2;
  }
  repeated Item list = 1;
}
message GetFriendModelReq {
  string SendUserId = 1;
  repeated string RecvUserIds = 2;
}
message GetFriendModelResp {
  repeated bytes models = 1;
}
message UpdateFriendModelReq {
  string SendUserId = 1;
  string RecvUserId = 2;
  map<string, string> updateMap = 3;
}
message UpdateFriendModelResp {
  string FailedReason = 1;
}
message GetFriendIdsReq {
  string SendUserId = 1;
}
message GetFriendIdsResp {
  repeated string FriendIds = 1;
}
message AddBlacklistReq {
  string SendUserId = 1;
  repeated string RecvUserIds = 2;
}
message AddBlacklistResp {
  string FailedReason = 1;
}
message DelBlacklistReq {
  string SendUserId = 1;
  repeated string RecvUserIds = 2;
}
message DelBlacklistResp {
  string FailedReason = 1;
}
message IsBlacklistReq {
  string SendUserId = 1;
  repeated string RecvUserIds = 2;
}
message IsBlacklistResp {
  message Item {
    string UserId = 1;
    bool IsBlacklist = 2;
  }
  repeated Item list = 1;
}
message GetBlacklistReq {
  string SendUserId = 1;
}
message GetBlacklistResp {
  repeated string BlacklistIds = 1;
}
message GetBlacklistModelReq {
  string SendUserId = 1;
  repeated string BlacklistIds = 2;
}
message GetBlacklistModelResp {
  repeated bytes models = 1;
}

service relationService {
  rpc AddFriend(AddFriendReq) returns (AddFriendResp);
  rpc DelFriend(DelFriendReq) returns (DelFriendResp);
  rpc IsFriend(IsFriendReq) returns (IsFriendResp);
  rpc GetFriendModel(GetFriendModelReq) returns (GetFriendModelResp);
  rpc UpdateFriendModel(UpdateFriendModelReq) returns (UpdateFriendModelResp);
  rpc GetFriendIds(GetFriendIdsReq) returns (GetFriendIdsResp);
  rpc AddBlacklist(AddBlacklistReq) returns (AddBlacklistResp);
  rpc DelBlacklist(DelBlacklistReq) returns (DelBlacklistResp);
  rpc IsBlacklist(IsBlacklistReq) returns (IsBlacklistResp);
  rpc GetBlacklist(GetBlacklistReq) returns (GetBlacklistResp);
  rpc GetBlacklistModel(GetBlacklistModelReq) returns (GetBlacklistModelResp);
}
```

## 编写逻辑

> 基本都是增删改查 跳过

## 单元测试

```go
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

```

---

# 下一章：