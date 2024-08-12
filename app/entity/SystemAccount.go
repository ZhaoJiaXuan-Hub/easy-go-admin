package entity

import "time"

func (SystemAccount) TableName() string {
	return "system_account"
}

type SystemAccount struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;comment:主键"`
	UserName  string    `gorm:"not null;comment:用户名"`
	NickName  string    `gorm:"comment:昵称"`
	Avatar    string    `gorm:"comment:头像"`
	Password  string    `gorm:"not null;comment:密码"`
	Salt      string    `gorm:"not null;comment:密码盐"`
	Email     string    `gorm:"comment:邮箱"`
	Status    int       `gorm:"comment:状态"`
	DeletedAt time.Time `gorm:"comment:删除时间"`
	CreatedAt time.Time `gorm:"comment:创建时间"`
	UpdatedAt time.Time `gorm:"comment:更新时间"`
}

type JwtSystemAccountData struct {
	Id       uint   `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
}

type LoginData struct {
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SystemAccountCreateData struct {
	UserName string `json:"user_name" validate:"required"`
	NickName string `json:"nick_name" validate:"required"`
	Avatar   string `json:"avatar" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
