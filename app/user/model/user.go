package usermodel

import (
	"github.com/Path-IM/Path-IM-Server/common/utils/encrypt"
	"github.com/Path-IM/Path-IM-Server/common/xcache/dc"
	"github.com/Path-IM/Path-IM-Server/common/xorm/global"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

func (u *User) FuncInsert(rc redis.UniversalClient) func(tx *gorm.DB) error {
	return func(tx *gorm.DB) error {
		return tx.Create(u).Error
	}
}
func (u *User) FlushCache(tx *gorm.DB, rc redis.UniversalClient) error {
	mapping := u.DbMapping(tx, rc)
	var keys []string
	// 用户模型
	{
		keys = append(keys, mapping.Key(u, map[string][]interface{}{
			"id": {u.Id},
		}), mapping.Key(u, map[string][]interface{}{
			"username": {u.Username},
		}))
	}
	return nil
}

// Preheat 预热缓存
func (u *User) Preheat(tx *gorm.DB, rc redis.UniversalClient) {
	mapping := u.DbMapping(tx, rc)
	err := mapping.FirstById(&User{}, u.Id)
	if err != nil {
		logx.Errorf("user.Preheat error: %s", err.Error())
	}
}

type User struct {
	Id           string        `gorm:"column:id;primary_key;type:char(32);comment:主键;" json:"id"`
	Username     string        `gorm:"column:username;type:char(32);index:,unique;not null;comment:用户名;" json:"username"`
	Password     string        `gorm:"column:password;type:char(64);not null;comment:密码;" json:"password"`
	Nickname     string        `gorm:"column:nickname;type:varchar(64);not null;default:'';comment:昵称;index;" json:"nickname"`
	Sign         string        `gorm:"column:sign;type:varchar(128);not null;default:'';comment:签名;" json:"sign"`
	Avatar       string        `gorm:"column:avatar;type:varchar(255);not null;default:'';comment:头像;" json:"avatar"`
	Province     string        `gorm:"column:province;type:varchar(64);not null;default:'';comment:省份;" json:"province"`
	City         string        `gorm:"column:city;type:varchar(64);not null;default:'';comment:城市;" json:"city"`
	District     string        `gorm:"column:district;type:varchar(64);not null;default:'';comment:区县;" json:"district"`
	RegisterTime int64         `gorm:"column:register_time;type:bigint(13);not null;default:0;comment:注册时间-毫秒级时间戳;" json:"registerTime"`
	IsMale       bool          `gorm:"column:is_male;index;comment:是否是男性;" json:"isMale"`
	dbMapping    *dc.DbMapping `gorm:"-"`
}

func (u *User) DbMapping(tx *gorm.DB, rc redis.UniversalClient) *dc.DbMapping {
	if u.dbMapping == nil {
		u.dbMapping = dc.GetDbMapping(rc, tx)
	}
	return u.dbMapping
}

func (u *User) BeforeCreate(_ *gorm.DB) error {
	if u.Id == "" {
		u.Id = global.GetID()
	}
	if u.Password != "" {
		u.Password = encrypt.Md5(u.Password)
	}
	return nil
}

func (u *User) GetIdString() string {
	return u.Id
}

func (u *User) TableName() string {
	return "users"
}
