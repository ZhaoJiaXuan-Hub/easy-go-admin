package service

import "github.com/gin-gonic/gin"

type ISystemRoleService interface {
	GetRoleList(c *gin.Context)
}
