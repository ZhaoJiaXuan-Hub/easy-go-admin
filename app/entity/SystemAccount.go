package entity

import (
	"easy-go-admin/config/util"
	"gorm.io/gorm"
)

func (SystemAccount) TableName() string {
	return "system_account"
}

// SystemAccount 管理员用户
type SystemAccount struct {
	ID        uint           `gorm:"primaryKey;autoIncrement;comment:主键" json:"id"`
	UserName  string         `gorm:"not null;comment:用户名" json:"username"`
	NickName  string         `gorm:"comment:昵称" json:"nickname"`
	Avatar    string         `gorm:"comment:头像;default:NULL" json:"avatar"`
	Password  string         `gorm:"not null;comment:密码" json:"password"`
	Salt      string         `gorm:"not null;comment:密码盐" json:"salt"`
	Email     string         `gorm:"comment:邮箱" json:"email"`
	Status    int            `gorm:"comment:状态;default:1" json:"status"`
	DeptID    uint           `gorm:"comment:部门" json:"dept_id"`
	Dept      *SystemDept    `gorm:"foreignKey:DeptID" json:"dept"`
	Roles     []*SystemRole  `gorm:"many2many:system_account_role" json:"roles"`
	DeletedAt gorm.DeletedAt `gorm:"index;softDelete;comment:删除时间;default:NULL" json:"deleted_at"` // 软删除
	CreatedAt util.Htime     `gorm:"comment:创建时间;type:datetime" json:"created_at"`
	UpdatedAt util.Htime     `gorm:"comment:更新时间;type:datetime" json:"updated_at"`
}

type JwtSystemAccountData struct {
	Id       uint          `json:"id"`
	UserName string        `json:"username"`
	NickName string        `json:"nickname"`
	Avatar   string        `json:"avatar"`
	Email    string        `json:"email"`
	Roles    []*SystemRole `json:"roles"`
	Dept     *SystemDept   `json:"dept"`
}

type LoginData struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SystemAccountCreateData struct {
	UserName string  `json:"username" validate:"required"`
	NickName string  `json:"nickname" validate:"required"`
	Avatar   *string `json:"avatar,omitempty" validate:"omitempty"`
	Email    string  `json:"email" validate:"required"`
	Password string  `json:"password" validate:"required"`
}
