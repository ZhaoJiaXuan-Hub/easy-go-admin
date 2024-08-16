package entity

import (
	"easy-go-admin/config/util"
	"gorm.io/gorm"
)

func (SystemMenu) TableName() string {
	return "system_menu"
}

// SystemMenu 菜单权限
type SystemMenu struct {
	ID         uint           `gorm:"primaryKey;autoIncrement;comment:主键" json:"id"`
	ParentID   uint           `gorm:"type:varchar(20);comment:父级菜单ID" json:"parentId"`
	Path       string         `gorm:"type:varchar(255);comment:URL路径" json:"path"`
	Component  string         `gorm:"type:varchar(255);comment:前端组件路径" json:"component"`
	Redirect   string         `gorm:"type:varchar(255);comment:重定向路径" json:"redirect"`
	Type       int            `gorm:"type:int;comment:菜单类型" json:"type"`
	Title      string         `gorm:"type:varchar(255);comment:标题" json:"title"`
	SvgIcon    string         `gorm:"type:varchar(255);comment:SVG图标路径" json:"svgIcon"`
	Icon       string         `gorm:"type:varchar(255);comment:字体图标名称" json:"icon"`
	Permission string         `gorm:"type:varchar(255);comment:权限标识" json:"permission"`
	KeepAlive  bool           `gorm:"comment:是否缓存保留" json:"keepAlive"`
	Hidden     bool           `gorm:"comment:是否隐藏" json:"hidden"`
	Sort       int            `gorm:"comment:排序" json:"sort"`
	ActiveMenu string         `gorm:"type:varchar(255);comment:菜单激活引用" json:"activeMenu"`
	Breadcrumb bool           `gorm:"comment:面包屑显示" json:"breadcrumb"`
	Status     int            `gorm:"type:int;comment:状态;default:1" json:"status"`
	ShowInTabs bool           `gorm:"comment:标签页显示" json:"showInTabs"`
	AlwaysShow bool           `gorm:"comment:是否总显示" json:"alwaysShow"`
	Affix      bool           `gorm:"comment:是否固定" json:"affix"`
	Children   []*SystemMenu  `gorm:"-" json:"children"`
	RolesInfo  []*SystemRole  `gorm:"many2many:system_role_menu" json:"rolesInfo"`
	Roles      []string       `gorm:"-" json:"roles"`
	DeletedAt  gorm.DeletedAt `gorm:"index;softDelete;comment:删除时间;default:NULL" json:"deleted_at"` // 软删除
	CreatedAt  util.Htime     `gorm:"comment:创建时间;type:datetime" json:"created_at"`
	UpdatedAt  util.Htime     `gorm:"comment:更新时间;type:datetime" json:"updated_at"`
}

// SystemMenuCreateData 创建结构体
type SystemMenuCreateData struct {
	Title      string `json:"title"`
	AlwaysShow bool   `json:"alwaysShow"`
	Breadcrumb bool   `json:"breadcrumb"`
	Hidden     bool   `json:"hidden"`
	Icon       string `json:"icon"`
	KeepAlive  bool   `json:"keepAlive"`
	ParentID   uint   `json:"parentId"`
	Path       string `json:"path"`
	Permission string `json:"permission"`
	Redirect   string `json:"redirect"`
	ShowInTabs bool   `json:"showInTabs"`
	SvgIcon    string `json:"svgIcon"`
	Type       int    `json:"type"`
	Component  string `json:"component"`
	Sort       int    `json:"sort"`
	Status     int    `json:"status"`
}

// SystemMenuUpdateData 更新结构体
type SystemMenuUpdateData struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	AlwaysShow bool   `json:"alwaysShow"`
	Breadcrumb bool   `json:"breadcrumb"`
	Hidden     bool   `json:"hidden"`
	Icon       string `json:"icon"`
	KeepAlive  bool   `json:"keepAlive"`
	ParentID   uint   `json:"parentId"`
	Path       string `json:"path"`
	Permission string `json:"permission"`
	Redirect   string `json:"redirect"`
	ShowInTabs bool   `json:"showInTabs"`
	SvgIcon    string `json:"svgIcon"`
	Type       int    `json:"type"`
	Component  string `json:"component"`
	Sort       int    `json:"sort"`
	Status     int    `json:"status"`
}

// SystemMenuDetailData 详情结构体
type SystemMenuDetailData struct {
	ID uint `json:"id"`
}

// AfterFind is a GORM callback that runs after the model is found.
// This function will populate the Roles field based on the RolesInfo.
func (m *SystemMenu) AfterFind(tx *gorm.DB) error {
	m.Roles = make([]string, len(m.RolesInfo))
	for i, role := range m.RolesInfo {
		m.Roles[i] = role.Code
	}
	return nil
}
