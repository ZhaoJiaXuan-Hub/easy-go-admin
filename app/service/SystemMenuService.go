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
	Save(c *gin.Context, dto entity.SystemMenuSaveData)
	Detail(c *gin.Context, dto entity.SystemMenuDetailData)
}

type SystemMenuServiceImpl struct{}

func (s SystemMenuServiceImpl) Detail(c *gin.Context, dto entity.SystemMenuDetailData) {
	// 获取详情参数验证
	err := Validate.New().Struct(dto)
	if err != nil {
		message.Fail(c, int(message.Code.BADREQUEST), message.Code.GetMessage(message.Code.BADREQUEST))
		return
	}
	menu, err := dao.SystemMenuDetail(dto.ID)
	if err != nil {
		message.Fail(c, int(message.Code.FAILURE), err.Error())
		return
	}
	message.Success(c, menu)
}

func (s SystemMenuServiceImpl) Save(c *gin.Context, dto entity.SystemMenuSaveData) {
	// 保存参数验证
	err := Validate.New().Struct(dto)
	if err != nil {
		message.Fail(c, int(message.Code.BADREQUEST), message.Code.GetMessage(message.Code.BADREQUEST))
		return
	}
	saveData := entity.SystemMenu{
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
	// 检查 id 是否为 0 来决定是创建还是更新
	if dto.ID != 0 {
		saveData.ID = dto.ID
		err = dao.SystemMenuUpdate(saveData)
	} else {
		err = dao.SystemMenuCreate(saveData)
	}
	if err != nil {
		message.Fail(c, int(message.Code.CREATERESOURCEFAILED), err.Error())
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
