package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Path-IM/Path-IM-Server/common/utils/encrypt"
	usermodel "github.com/showurl/Path-IM-Server-OICQ/app/user/model"
	"github.com/showurl/Path-IM-Server-OICQ/app/user/rpc/pb"
	"github.com/showurl/Path-IM-Server-OICQ/app/user/rpc/userservice"
	"github.com/zeromicro/go-zero/zrpc"
	"testing"
)

var (
	service userservice.UserService
	ctx     = context.Background()
)

func init() {
	service = userservice.NewUserService(zrpc.MustNewClient(zrpc.RpcClientConf{
		Endpoints: []string{"192.168.1.98:10013"},
	}))
}
func TestRegister(t *testing.T) {
	resp, err := service.Register(ctx, &pb.RegisterReq{
		Username: fmt.Sprintf("test-%d", 1),
		Password: "123456",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
}
func TestGetToken(t *testing.T) {
	resp, err := service.GetToken(ctx, &pb.GetTokenReq{
		UserId:   "test-1",
		Platform: "IOS",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
}
func TestGetUserByUsername(t *testing.T) {
	resp, err := service.GetUserByUsername(ctx, &pb.GetUserByUsernameReq{
		Username: "test-1",
	})
	if err != nil {
		t.Error(err)
	}
	user := &usermodel.User{}
	_ = json.Unmarshal(resp.User, user)
	t.Log(resp.String())
	t.Log(user.Password)
	t.Log(encrypt.Md5("123456"))
}
func TestGetUserByIds(t *testing.T) {
	resp, err := service.GetUserByIds(ctx, &pb.GetUserByIdsReq{
		UserId: []string{"3604f15eed0f94282e271ed583427eec", "test-1"},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
}
func TestGetUserConfig(t *testing.T) {
	resp, err := service.GetUserConfig(ctx, &pb.GetUserConfigReq{
		UserId: "3604f15eed0f94282e271ed583427eec",
		Ks:     []string{},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
	resp, err = service.GetUserConfig(ctx, &pb.GetUserConfigReq{
		UserId: "3604f15eed0f94282e271ed583427eec",
		Ks: []string{
			"k1", "k2",
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
}
func TestUpdateUserConfig(t *testing.T) {
	resp, err := service.UpdateUserConfig(ctx, &pb.UpdateUserConfigReq{
		UserId: "3604f15eed0f94282e271ed583427eec",
		Configs: map[string]string{
			"k1": "v1",
			"k2": "v2",
			"k3": "v3",
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.String())
}
