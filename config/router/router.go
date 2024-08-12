package router

import (
	"easy-go-admin/app/middleware"
	"easy-go-admin/config/router/module"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// CreateHttpHandler 创建路由
func CreateHttpHandler() *gin.Engine {
	root := gin.New()
	// 跨域中间件
	root.Use(middleware.CorsMiddleware())
	// 日志中间件
	root.Use(middleware.LoggerMiddleware())
	root.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	base := root.Group("/")
	// 全局路由
	routerModule.SetupCommonRouter(base)
	return root
}
