package entity

import (
	"easy-go-admin/config/util"
	"gorm.io/gorm"
)

func (SystemDept) TableName() string {
	return "system_dept"
}

// SystemDept 部门
type SystemDept struct {
	ID          uint           `gorm:"primaryKey;autoIncrement;comment:主键" json:"id"`
	ParentID    string         `gorm:"type:varchar(20);comment:上级部门" json:"parent_id"`
	Name        string         `gorm:"type:varchar(255);comment:部门名称" json:"name"`
	Sort        int            `gorm:"type:int;comment:排序" json:"sort"`
	Status      int            `gorm:"type:int;comment:状态" json:"status"`
	Description string         `gorm:"type:varchar(255);comment:描述" json:"description"`
	DeletedAt   gorm.DeletedAt `gorm:"index;softDelete;comment:删除时间;default:NULL" json:"deleted_at"` // 软删除
	CreatedAt   util.Htime     `gorm:"comment:创建时间;type:datetime" json:"created_at"`
	UpdatedAt   util.Htime     `gorm:"comment:更新时间;type:datetime" json:"updated_at"`
}
