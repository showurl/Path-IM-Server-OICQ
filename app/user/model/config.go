package usermodel

const (
	RedisKeyConfig = "user:user_configs:"
)

type Config struct {
	UserId string `gorm:"column:user_id;index;type:char(32);comment:用户id;not null;" json:"userId"`
	K      string `gorm:"column:k;index;type:varchar(255);comment:配置键;not null;" json:"k"`
	V      string `gorm:"column:v;type:varchar(255);comment:配置值;not null;" json:"v"`
}

func (u *Config) TableName() string {
	return "user_configs"
}
