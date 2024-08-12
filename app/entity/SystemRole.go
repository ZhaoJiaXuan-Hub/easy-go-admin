package entity

import "time"

func (SystemRole) TableName() string {
	return "system_role"
}

type SystemRole struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;comment:主键"`
	Name      string    `gorm:"not null;comment:角色名称"`
	Desc      string    `gorm:"comment:描述"`
	CreatedAt time.Time `gorm:"comment:创建时间"`
	UpdatedAt time.Time `gorm:"comment:更新时间"`
}
