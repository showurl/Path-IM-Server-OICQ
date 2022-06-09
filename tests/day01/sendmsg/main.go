package main

import (
	"encoding/json"
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
	wsMap       = make(map[string]*websocket.Conn)
	mapLock     = sync.Mutex{}
	WsAddr      = "ws://192.168.1.98:11000"
	Token       = "123"
	Platform    = "IOS"
	sendMsgLock = sync.Mutex{}
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

func SendMsg(body *pb.BodyReq) error {
	sendMsgLock.Lock()
	defer sendMsgLock.Unlock()
	buf, _ := proto.Marshal(body)
	err := Ws(body.SendID).WriteMessage(websocket.BinaryMessage, buf)
	return err
}

func SendMsyByMsgData(msgData *chatpb.MsgData) error {
	req := &chatpb.SendMsgReq{
		MsgData: msgData,
	}
	data, _ := proto.Marshal(req)
	body := &pb.BodyReq{
		ReqIdentifier: types.WSSendMsg,
		Token:         Token,
		SendID:        msgData.SendID,
		Data:          data,
	}
	return SendMsg(body)
}

func main() {
	Ws("2")
	buf, _ := json.Marshal(map[string]string{
		"Text": "给1群发消息",
	})
	SendMsyByMsgData(&chatpb.MsgData{
		ClientMsgID:      "1",
		ConversationType: types.GroupChatType,
		SendID:           "2",
		ReceiveID:        "group1",
		ContentType:      0,
		Content:          buf,
		AtUserIDList:     nil,
		OfflinePush: &chatpb.OfflinePush{
			Title:         "2给你发了一条消息",
			Desc:          "hello, I'm Path-IM 2",
			Ex:            "时间：2018-01-01 12:00:00",
			IOSPushSound:  "xx",
			IOSBadgeCount: true,
		},
		ClientTime: time.Now().UnixMilli(),
		ServerTime: 0,
		Seq:        0,
		MsgOptions: &chatpb.MsgOptions{
			Persistent:         true,
			History:            true,
			UpdateUnreadCount:  true,
			UpdateConversation: true,
		},
	})
}
