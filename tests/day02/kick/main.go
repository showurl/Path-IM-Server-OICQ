package main

import (
	"context"
	"fmt"
	onlinemessagerelayservice "github.com/Path-IM/Path-IM-Server/app/msg-gateway/cmd/wsrpc/onlineMessageRelayService"
	msggatewaypb "github.com/Path-IM/Path-IM-Server/app/msg-gateway/cmd/wsrpc/pb"
	"github.com/zeromicro/go-zero/zrpc"
)

func main() {
	service := onlinemessagerelayservice.NewOnlineMessageRelayService(zrpc.MustNewClient(zrpc.RpcClientConf{
		Endpoints: []string{"192.168.1.98:10001"},
	}))
	status, err := service.GetUsersOnlineStatus(context.Background(), &msggatewaypb.GetUsersOnlineStatusReq{UserIDList: []string{"1", "2"}})
	if err != nil {
		panic(err)
	}
	for _, userStatus := range status.StatusList {
		fmt.Printf("%s: \n", userStatus.UserID)
		for k, v := range userStatus.PlatformAddrMap {
			fmt.Printf("%s: %s\n", k, v)
			conns, err := service.KickUserConns(context.Background(), &msggatewaypb.KickUserConnsReq{UserID: userStatus.UserID, PlatformIDs: []string{
				k,
			}})
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s: %s\n", userStatus.UserID, conns)
		}
		fmt.Println()
	}
}
