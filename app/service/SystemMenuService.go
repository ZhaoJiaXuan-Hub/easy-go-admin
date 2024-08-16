package service

import (
	"easy-go-admin/app/dao"
	"easy-go-admin/app/entity"
	"easy-go-admin/config/message"
	"github.com/gin-gonic/gin"
	Validate "github.com/go-playground/validator/v10"
)

type ISystemMenuService interface {
	GetMenuTree(c *gin.Context)
	Create(c *gin.Context, dto entity.SystemMenuCreateData)
}

type SystemMenuServiceImpl struct{}

func (s SystemMenuServiceImpl) Create(c *gin.Context, dto entity.SystemMenuCreateData) {
	// 创建参数验证
	err := Validate.New().Struct(dto)
	if err != nil {
		message.Fail(c, int(message.Code.BADREQUEST), message.Code.GetMessage(message.Code.BADREQUEST))
		return
	}
	createData := entity.SystemMenu{
		Title:      dto.Title,
		AlwaysShow: dto.AlwaysShow,
		Breadcrumb: dto.Breadcrumb,
		Hidden:     dto.Hidden,
		Icon:       dto.Icon,
		KeepAlive:  dto.KeepAlive,
		ParentID:   dto.ParentID,
		Permission: dto.Permission,
		Redirect:   dto.Redirect,
		ShowInTabs: dto.ShowInTabs,
		SvgIcon:    dto.SvgIcon,
		Type:       dto.Type,
		Component:  dto.Component,
		Sort:       dto.Sort,
		Status:     dto.Status,
		Path:       dto.Path,
	}
	err = dao.SystemMenuCreate(createData)
	if err != nil {
		message.Fail(c, int(message.Code.CREATERESOURCEFAILED), "创建菜单失败")
		return
	}
	message.Success(c, nil)
}

func SystemMenuService() ISystemMenuService {
	return &SystemMenuServiceImpl{}
}

// GetMenuTree 获取树状菜单数据
func (s SystemMenuServiceImpl) GetMenuTree(c *gin.Context) {
	menuData, err := dao.SystemMenuAll()
	if err != nil {
		message.Fail(c, int(message.Code.FAILURE), err.Error())
		return
	}
	uniqueMenus := make(map[uint]*entity.SystemMenu)
	for _, menu := range menuData {
		uniqueMenus[menu.ID] = &menu
	}
	// 对菜单进行树状归类
	treeMenuList := BuildNestedMenus(uniqueMenus)
	message.Success(c, treeMenuList)
}
