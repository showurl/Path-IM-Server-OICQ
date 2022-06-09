package main

import (
	"fmt"
	"github.com/Path-IM/Path-IM-Server/app/msg-gateway/cmd/wsrpc/pb"
	"github.com/Path-IM/Path-IM-Server/common/types"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	chatpb "github.com/showurl/Path-IM-Server-OICQ/app/msg/rpc/pb"
	"log"
	"sync"
	"time"
)

var (
	wsMap    = make(map[string]*websocket.Conn)
	mapLock  = sync.Mutex{}
	WsAddr   = "ws://192.168.1.98:11000"
	Token    = "123"
	Platform = "IOS"
	//sendMsgLock = sync.Mutex{}
)

func Ws(uid string) *websocket.Conn {
	mapLock.Lock()
	defer mapLock.Unlock()
	if ws, ok := wsMap[uid]; !ok || ws == nil {
		var (
			err error
			w   *websocket.Conn
		)
		url := fmt.Sprintf("%s?token=%s&userID=%s&platform=%s", WsAddr, Token, uid, Platform)
		w, _, err = websocket.DefaultDialer.Dial(
			url,
			nil)
		if err != nil {
			log.Printf("url: %v", url)
			log.Printf("dial: %v", err)
			time.Sleep(time.Second)
			return Ws(uid)
		}
		wsMap[uid] = w
		return w
	} else {
		return ws
	}
}
func main() {
	ws := Ws("1")
	for {
		typ, body, err := ws.ReadMessage()
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		if typ == websocket.BinaryMessage {
			bodyResp := &pb.BodyResp{}
			_ = proto.Unmarshal(body, bodyResp)
			switch bodyResp.ReqIdentifier {
			case types.WSPushMsg, types.WSGroupPushMsg:
				msgData := &chatpb.MsgData{}
				_ = proto.Unmarshal(bodyResp.Data, msgData)
				fmt.Printf(fmt.Sprintf(`
被推送的
%s
%s
`, string(msgData.Content), msgData.String()))

			}
		}
	}
}
