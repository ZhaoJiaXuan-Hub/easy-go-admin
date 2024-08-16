package entity

import (
	"easy-go-admin/config/util"
	"gorm.io/gorm"
)

func (SystemDict) TableName() string {
	return "system_dict"
}

// SystemDict 字典
type SystemDict struct {
	ID        uint             `gorm:"primaryKey;autoIncrement;comment:主键" json:"id"`
	Name      string           `gorm:"not null;comment:字典名称" json:"name"`
	Desc      string           `gorm:"comment:描述" json:"desc"`
	Code      string           `gorm:"comment:字典编码" json:"code"`
	Status    int              `gorm:"comment:字典状态;default:1" json:"status"`
	DictData  []SystemDictData `gorm:"ForeignKey:DictID" json:"dict_data"`                           // 新增字段，表示关联的 SystemDictData 切片
	DeletedAt gorm.DeletedAt   `gorm:"index;softDelete;comment:删除时间;default:NULL" json:"deleted_at"` // 软删除
	CreatedAt util.Htime       `gorm:"comment:创建时间;type:datetime" json:"created_at"`
	UpdatedAt util.Htime       `gorm:"comment:更新时间;type:datetime" json:"updated_at"`
}
