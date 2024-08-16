package entity

import (
	"gorm.io/gorm"
	"time"
)

func (SystemRole) TableName() string {
	return "system_role"
}

// SystemRole 角色
type SystemRole struct {
	ID          uint           `gorm:"primaryKey;autoIncrement;comment:主键" json:"id"`
	Name        string         `gorm:"not null;comment:角色名称" json:"name"`
	Description string         `gorm:"comment:描述" json:"description"`
	Code        string         `gorm:"comment:角色编码" json:"code"`
	Status      int            `gorm:"comment:角色状态;default:1" json:"status"`
	Menus       []*SystemMenu  `gorm:"many2many:system_role_menu" json:"menus"`
	DeletedAt   gorm.DeletedAt `gorm:"index;softDelete;comment:删除时间;default:NULL" json:"deleted_at"` // 软删除
	CreatedAt   time.Time      `gorm:"comment:创建时间" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"comment:更新时间" json:"updated_at"`
}
