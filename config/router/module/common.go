package routerModule

import (
	"easy-go-admin/app/controller"
	controllerSystem "easy-go-admin/app/controller/system"
	"github.com/gin-gonic/gin"
)

// SetupCommonRouter 设置全局路由
func SetupCommonRouter(base *gin.RouterGroup) {
	index := base.Group("/index")
	index.GET("/index", controller.Index)

	system := base.Group("/system")
	systemAccount := system.Group("/account")
	systemAccount.POST("/login", controllerSystem.Login)
	systemAccount.POST("/create", controllerSystem.Create)
}
