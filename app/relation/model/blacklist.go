package relationmodel

import "encoding/json"

type Blacklist struct {
	SendId    string `gorm:"column:send_id;index:sr,unique;index;type:char(32);comment:'发送者ID'" json:"sendId"`
	ReceiveId string `gorm:"column:receive_id;index:sr,unique;index;type:char(32);comment:'接收者ID'" json:"receiveId"`
	CreatedAt int64  `gorm:"column:created_at;index;type:bigint(13);not null;default:0;comment:'创建时间-毫秒级时间戳'" json:"createdAt"`
}

func (g *Blacklist) TableName() string {
	return "relation_blacklists"
}

func (g *Blacklist) Bytes() []byte {
	buf, _ := json.Marshal(g)
	return buf
}
