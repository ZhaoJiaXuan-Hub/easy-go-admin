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
// @router /api/system/account/login [post]
func Login(c *gin.Context) {
	var dto entity.LoginData
	_ = c.BindJSON(&dto)
	service.SystemAccountService().Login(c, dto)
}

// CreateAccount 创建管理员
// @Summary 创建管理员接口
// @Produce json
// @Description 后台创建管理员登陆接口
// @Param data body entity.SystemAccountCreateData true "data"
// @Success 200 {object} message.Response
// @router /api/system/account/create [post]
func CreateAccount(c *gin.Context) {
	var dto entity.SystemAccountCreateData
	_ = c.BindJSON(&dto)
	service.SystemAccountService().Create(c, dto)
}

// GetAccountDetail 获取登录用户信息
// @Summary 获取登录用户信息
// @Produce json
// @Description 后台获取登录用户信息
// @Security ApiKeyAuth
// @Success 200 {object} message.Response
// @router /api/system/account/getDetail [get]
func GetAccountDetail(c *gin.Context) {
	service.SystemAccountService().GetDetail(c)
}

// GetRouters 获取登录用户可用目录菜单
// @Summary 获取登录用户可用目录菜单
// @Produce json
// @Description 获取登录用户可用目录菜单
// @Security ApiKeyAuth
// @Success 200 {object} message.Response
// @router /api/system/account/getRouters [get]
func GetRouters(c *gin.Context) {
	service.SystemAccountService().GetRouters(c)
}
