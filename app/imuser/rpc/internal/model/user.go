package model

type User struct {
	Username string `gorm:"column:username;type:varchar(255);not null;primary_key"`
	Password string `gorm:"column:password;type:varchar(255);not null"`
}

func (u *User) TableName() string {
	return "simple_user"
}
