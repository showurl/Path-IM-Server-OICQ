package groupmodel

import "encoding/json"

type GroupMember struct {
	GroupID  string `gorm:"column:group_id;index:gm,unique;index;type:char(32);comment:'群组ID'" json:"groupId"`
	UserID   string `gorm:"column:user_id;index:gm,unique;index;type:char(32);comment:'用户ID'" json:"userId"`
	JoinedAt int64  `gorm:"column:joined_at;index;type:bigint(13);not null;default:0;comment:'加入时间-毫秒级时间戳'" json:"joinedAt"`
	MsgOpt   int    `gorm:"column:msg_opt;index;type:tinyint(1);not null;default:0;comment:'消息接收选项,0正常,1收但不通知,2不接收'" json:"msgOpt"`
	Remark   string `gorm:"column:remark;type:varchar(64);comment:'群备注'" json:"remark"`
}

func (g *GroupMember) TableName() string {
	return "group_members"
}

func (g *GroupMember) Bytes() []byte {
	buf, _ := json.Marshal(g)
	return buf
}
