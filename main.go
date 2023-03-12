package main

import (
	"go-gin-video/initialization"
	"go-gin-video/router"
	"go.uber.org/zap"
)

// @title go-gin-videoSwaggerApi
// @version 1.0
// @description go-gin-videoSwaggerApi
// @contact.name loveyu233
// @host http://127.0.0.1:9898
func main() {
	// 初始化所有连接和配置
	initialization.InitAll()

	// 初始化路由
	if err := router.InitRouter(); err != nil {
		zap.L().Error("路由初始化错误: " + err.Error())
	}
}
