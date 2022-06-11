package groupmodel

import (
	"encoding/json"
	"github.com/Path-IM/Path-IM-Server/common/xcache/dc"
	"github.com/Path-IM/Path-IM-Server/common/xorm/global"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

func (u *Group) FuncInsert(rc redis.UniversalClient) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return tx.Create(u).Error
	}
}
func (u *Group) FlushCache(tx *gorm.DB, rc redis.UniversalClient) error {
	mapping := u.DbMapping(tx, rc)
	var keys []string
	// 群组模型
	{
		keys = append(keys, mapping.Key(u, map[string][]interface{}{
			"id": {u.Id},
		}))
	}
	return nil
}

// Preheat 预热缓存
func (u *Group) Preheat(tx *gorm.DB, rc redis.UniversalClient) {
	mapping := u.DbMapping(tx, rc)
	err := mapping.FirstById(&Group{}, u.Id)
	if err != nil {
		logx.Errorf("group.Preheat error: %s", err.Error())
	}
}

type Group struct {
	Id        string        `gorm:"column:id;primary_key;type:char(32);comment:'群组ID'" json:"id"`
	Name      string        `gorm:"column:name;type:varchar(64);comment:'群组名称'" json:"name"`
	CreatedAt int64         `gorm:"column:created_at;index;type:bigint(13);not null;default:0;comment:'创建时间-毫秒级时间戳'" json:"createdAt"`
	dbMapping *dc.DbMapping `gorm:"-"`
}

func (u *Group) DbMapping(tx *gorm.DB, rc redis.UniversalClient) *dc.DbMapping {
	if u.dbMapping == nil {
		u.dbMapping = dc.GetDbMapping(rc, tx)
	}
	return u.dbMapping
}

func (u *Group) BeforeCreate(_ *gorm.DB) error {
	if u.Id == "" {
		u.Id = global.GetID()
	}
	u.CreatedAt = time.Now().UnixMilli()
	return nil
}

func (u *Group) GetIdString() string {
	return u.Id
}

func (u *Group) TableName() string {
	return "groups"
}

func (u *Group) Bytes() []byte {
	buf, _ := json.Marshal(u)
	return buf
}
