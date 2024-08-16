package controllerSystem

import (
	"easy-go-admin/app/entity"
	"easy-go-admin/app/service"
	"github.com/gin-gonic/gin"
)

// GetMenuTree 获取所有树状菜单数据
// @Summary 获取所有树状菜单数据
// @Produce json
// @Description 获取所有树状菜单数据
// @Success 200 {object} message.Response
// @router /api/system/menu/getMenuTree [get]
func GetMenuTree(c *gin.Context) {
	service.SystemMenuService().GetMenuTree(c)
}

// CreateMenu 创建菜单
// @Summary 创建菜单
// @Produce json
// @Description 创建菜单接口
// @Param data body entity.SystemMenuCreateData true "data"
// @Success 200 {object} message.Response
// @router /api/system/menu/create [post]
func CreateMenu(c *gin.Context) {
	var dto entity.SystemMenuCreateData
	_ = c.BindJSON(&dto)
	service.SystemMenuService().Create(c, dto)
}

func UpdateMenu() {

}
