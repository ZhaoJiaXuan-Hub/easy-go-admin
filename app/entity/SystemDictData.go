package entity

import (
	"easy-go-admin/config/util"
	"gorm.io/gorm"
)

func (SystemDictData) TableName() string {
	return "system_dict_data"
}

// SystemDictData 字典数据
type SystemDictData struct {
	ID        uint           `gorm:"primaryKey;autoIncrement;comment:主键" json:"id"`
	Label     string         `gorm:"not null;comment:字典名" json:"label"`
	Value     string         `gorm:"comment:字典值" json:"value"`
	Status    int            `gorm:"comment:字典状态;default:1" json:"status"`
	DictID    uint           `gorm:"not null;comment:所属字典ID" json:"dict_id"`                       // 新增字段，存储所属字典的ID
	DeletedAt gorm.DeletedAt `gorm:"index;softDelete;comment:删除时间;default:NULL" json:"deleted_at"` // 软删除
	CreatedAt util.Htime     `gorm:"comment:创建时间;type:datetime" json:"created_at"`
	UpdatedAt util.Htime     `gorm:"comment:更新时间;type:datetime" json:"updated_at"`
}
