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
