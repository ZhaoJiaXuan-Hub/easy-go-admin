package controllerSystem

import (
	"easy-go-admin/app/entity"
	"easy-go-admin/app/service"
	"fmt"
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

// SaveMenu 保存菜单
// @Summary 保存菜单
// @Produce json
// @Description 保存菜单接口
// @Param data body entity.SystemMenuSaveData true "data"
// @Success 200 {object} message.Response
// @router /api/system/menu/save [post]
func SaveMenu(c *gin.Context) {
	var dto entity.SystemMenuSaveData
	_ = c.BindJSON(&dto)
	service.SystemMenuService().Save(c, dto)
}

// GetMenuDetail 获取菜单详情
// @Summary 获取菜单详情
// @Produce json
// @Description 获取菜单详情接口
// @Param data body entity.SystemMenuDetailData true "data"
// @Success 200 {object} message.Response
// @router /api/system/menu/detail [get]
func GetMenuDetail(c *gin.Context) {
	var dto entity.SystemMenuDetailData
	// 获取获取请求参数
	_ = c.ShouldBindQuery(&dto)
	fmt.Println(c)
	service.SystemMenuService().Detail(c, dto)
}

func deleteMenu(c *gin.Context) {

}
