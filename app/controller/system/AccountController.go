package controllerSystem

import (
	"easy-go-admin/app/entity"
	"easy-go-admin/app/service"
	"github.com/gin-gonic/gin"
)

// Login 管理员登陆接口
// @Summary 管理员登陆接口
// @Produce json
// @Description 管理员后台登陆接口
// @Param data body entity.LoginData true "data"
// @Success 200 {object} message.Response
// @router /system/account/login [post]
func Login(c *gin.Context) {
	var dto entity.LoginData
	_ = c.BindJSON(&dto)
	service.SystemAccountService().Login(c, dto)
}

// Create 创建管理员
// @Summary 创建管理员接口
// @Produce json
// @Description 后台创建管理员登陆接口
// @Param data body entity.SystemAccountCreateData true "data"
// @Success 200 {object} message.Response
// @router /system/account/create [post]
func Create(c *gin.Context) {
	var dto entity.SystemAccountCreateData
	_ = c.BindJSON(&dto)
	service.SystemAccountService().Create(c, dto)
}
