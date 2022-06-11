package relationmodel

import "encoding/json"

type Friend struct {
	SendId    string `gorm:"column:send_id;index:sr,unique;index;type:char(32);comment:'发送者ID'" json:"sendId"`
	ReceiveId string `gorm:"column:receive_id;index:sr,unique;index;type:char(32);comment:'接收者ID'" json:"receiveId"`
	CreatedAt int64  `gorm:"column:created_at;index;type:bigint(13);not null;default:0;comment:'创建时间-毫秒级时间戳'" json:"createdAt"`
	MsgOpt    int    `gorm:"column:msg_opt;index;type:tinyint(1);not null;default:0;comment:'消息接收选项,0正常,1收但不通知,2不接收'" json:"msgOpt"`
	Remark    string `gorm:"column:remark;type:varchar(64);comment:'备注'" json:"remark"`
}

func (g *Friend) TableName() string {
	return "relation_friends"
}

func (g *Friend) Bytes() []byte {
	buf, _ := json.Marshal(g)
	return buf
}
