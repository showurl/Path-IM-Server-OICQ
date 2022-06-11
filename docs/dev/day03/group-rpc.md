# 上一章：[用户模块](user-rpc.md)

---

# 群组模块

## 编写proto file

```protobuf
syntax = "proto3";

option go_package = "./pb";

package group;
message CreateGroupReq {
  string name = 1;
  repeated string members = 2;
}
message CreateGroupResp {
  string FailedReason = 1;
  string groupId = 2;
}
message GetGroupReq {
  repeated string groupIds = 1;
}
message GetGroupResp {
  repeated bytes groups = 1;
}
message UpdateGroupReq {
  string groupId = 1;
  string name = 2;
}
message UpdateGroupResp {
  string FailedReason = 1;
}
message DeleteGroupReq {
  string groupId = 1;
}
message DeleteGroupResp {
  string FailedReason = 1;
}
message GetGroupMemberReq {
  string groupId = 1;
}
message GetGroupMemberResp {
  repeated string members = 1;
}
message AddGroupMemberReq {
  string groupId = 1;
  string member = 2;
}
message AddGroupMemberResp {
  string FailedReason = 1;
}
message DeleteGroupMemberReq {
  string groupId = 1;
  string member = 2;
}
message DeleteGroupMemberResp {
  string FailedReason = 1;
}
service groupService {
  rpc CreateGroup(CreateGroupReq) returns (CreateGroupResp);
  rpc GetGroup(GetGroupReq) returns (GetGroupResp);
  rpc UpdateGroup(UpdateGroupReq) returns (UpdateGroupResp);
  rpc DeleteGroup(DeleteGroupReq) returns (DeleteGroupResp);
  rpc GetGroupMember(GetGroupMemberReq) returns (GetGroupMemberResp);
  rpc AddGroupMember(AddGroupMemberReq) returns (AddGroupMemberResp);
  rpc DeleteGroupMember(DeleteGroupMemberReq) returns (DeleteGroupMemberResp);
}
```

## 编写model

```go
package groupmodel

type Group struct {
	Id        string `gorm:"column:id;primary_key;type:char(32);comment:'群组ID'" json:"id"`
	Name      string `gorm:"column:name;type:varchar(64);comment:'群组名称'" json:"name"`
	CreatedAt int64  `gorm:"column:created_at;index;type:bigint(13);not null;default:0;comment:'创建时间-毫秒级时间戳'" json:"createdAt"`
}
type GroupMember struct {
	GroupID  string `gorm:"column:group_id;index:gm,unique;index;type:char(32);comment:'群组ID'" json:"groupId"`
	UserID   string `gorm:"column:user_id;index:gm,unique;index;type:char(32);comment:'用户ID'" json:"userId"`
	JoinedAt int64  `gorm:"column:joined_at;index;type:bigint(13);not null;default:0;comment:'加入时间-毫秒级时间戳'" json:"joinedAt"`
	MsgOpt   int    `gorm:"column:msg_opt;index;type:tinyint(1);not null;default:0;comment:'消息接收选项,0正常,1收但不通知,2不接收'" json:"msgOpt"`
	Remark   string `gorm:"column:remark;type:varchar(64);comment:'群备注'" json:"remark"`
}
```

## 编写rpc

> 简单的增删改查 没有特殊的处理

## 测试rpc

```go
package tests

import (
	"context"
	"github.com/showurl/Path-IM-Server-OICQ/app/group/rpc/groupservice"
	"github.com/showurl/Path-IM-Server-OICQ/app/group/rpc/pb"
	"github.com/zeromicro/go-zero/zrpc"
	"testing"
)

var (
	service groupservice.GroupService
	ctx     = context.Background()
)

func init() {
	service = groupservice.NewGroupService(zrpc.MustNewClient(zrpc.RpcClientConf{
		Endpoints: []string{"192.168.1.98:10014"},
	}))
}
func TestCreateGroup(t *testing.T) {
	resp, err := service.CreateGroup(ctx, &pb.CreateGroupReq{
		Name:    "第一个群哦",
		Members: []string{"3604f15eed0f94282e271ed583427eec"},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
}
func TestAddGroupMember(t *testing.T) {
	resp, err := service.AddGroupMember(ctx, &pb.AddGroupMemberReq{
		GroupId: "bea24ed2d0cc41ae8bf06b69bbf028e6",
		Member:  "3604f15eed0f94282e271ed583427eec",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
}
func TestDeleteGroup(t *testing.T) {
	resp, err := service.DeleteGroup(ctx, &pb.DeleteGroupReq{
		GroupId: "bea24ed2d0cc41ae8bf06b69bbf028e6",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
}
func TestDeleteGroupMember(t *testing.T) {
	resp, err := service.DeleteGroupMember(ctx, &pb.DeleteGroupMemberReq{
		GroupId: "bea24ed2d0cc41ae8bf06b69bbf028e6",
		Member:  "3604f15eed0f94282e271ed583427eec",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
}
func TestGetGroupIds(t *testing.T) {
	resp, err := service.GetGroupIds(ctx, &pb.GetGroupIdsReq{
		UserId: "3604f15eed0f94282e271ed583427eec",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
}
func TestGetGroup(t *testing.T) {
	resp, err := service.GetGroup(ctx, &pb.GetGroupReq{
		GroupIds: []string{"bea24ed2d0cc41ae8bf06b69bbf028e6"},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
}
func TestGetGroupMember(t *testing.T) {
	resp, err := service.GetGroupMember(ctx, &pb.GetGroupMemberReq{
		GroupId: "bea24ed2d0cc41ae8bf06b69bbf028e6",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
}
func TestGetGroupMemberModel(t *testing.T) {
	resp, err := service.GetGroupMemberModel(ctx, &pb.GetGroupMemberModelReq{
		GroupIds: []string{"bea24ed2d0cc41ae8bf06b69bbf028e6"},
		UserId:   "3604f15eed0f94282e271ed583427eec",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
}
func TestUpdateGroup(t *testing.T) {
	resp, err := service.UpdateGroup(ctx, &pb.UpdateGroupReq{
		GroupId: "bea24ed2d0cc41ae8bf06b69bbf028e6",
		Name:    "改个名吧",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
}

```

---

# 下一章：[关系模块](relation-rpc.md)