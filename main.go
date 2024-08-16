package main

import (
	"context"
	"easy-go-admin/config"
	"easy-go-admin/config/database"
	"easy-go-admin/config/log"
	"easy-go-admin/config/redis"
	"easy-go-admin/config/router"
	_ "easy-go-admin/docs"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// main 启动程序
// @title easyAdmin后台管理系统
// @version 1.0
// @description 后台管理系统api接口文档
// @securityDefinitions.apikey ApikeyAuth
// @in header
// @name jason
func main() {
	Log := log.Log()
	gin.SetMode(config.Config.Server.Model)
	route := router.CreateHttpHandler()
	srv := &http.Server{
		Addr:    ":" + config.Config.Server.Address,
		Handler: route,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			Log.Info("监听端口:", err)
		}
		Log.Info("监听端口:", config.Config.Server.Address)
	}()
	fmt.Printf("系统项目启动成功\n")
	fmt.Printf("BaseUrl: http://localhost:%s\n", config.Config.Server.Address)
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	Log.Info("关闭服务 ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		Log.Fatal("关闭服务:", err)
	}
	Log.Info("服务已关闭")
}

// init 初始化
func init() {
	fmt.Println("初始化数据库")
	err := database.InitDB()
	if err != nil {
		fmt.Println("初始化数据库失败:", err)
	} else {
		fmt.Println("初始化数据库成功")
	}
	err = redis.InitRedis()
	if err != nil {
		fmt.Println("初始化redis失败:", err)
	} else {
		fmt.Println("初始化redis成功")
	}
}
