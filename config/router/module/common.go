package routerModule

import (
	"easy-go-admin/app/controller"
	controllerSystem "easy-go-admin/app/controller/system"
	"easy-go-admin/app/middleware"
	"github.com/gin-gonic/gin"
)

// SetupCommonRouter 设置全局路由
func SetupCommonRouter(base *gin.RouterGroup) {
	index := base.Group("/index")
	index.GET("/index", controller.Index)

	api := base.Group("/api")
	system := api.Group("/system")

	systemAccount := system.Group("/account")
	systemAccount.POST("/login", controllerSystem.Login)
	// 鉴权中间件
	systemAccount.Use(middleware.AuthMiddleware())
	{
		systemAccount.POST("/create", controllerSystem.CreateAccount)
		systemAccount.GET("/getDetail", controllerSystem.GetAccountDetail)
		systemAccount.GET("/getRouters", controllerSystem.GetRouters)
	}

	dict := system.Group("/dict")
	dict.GET("/dictData", controllerSystem.DictData)

	systemMenu := system.Group("/menu")
	// 鉴权中间件
	systemMenu.Use(middleware.AuthMiddleware())
	{
		systemMenu.GET("/getMenuTree", controllerSystem.GetMenuTree)
		systemMenu.POST("/save", controllerSystem.SaveMenu)
		systemMenu.GET("/detail", controllerSystem.GetMenuDetail)
	}
}
