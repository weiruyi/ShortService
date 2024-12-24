package main

import (
	"ShortService/src/global"
	"ShortService/src/initialize"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	router := initialize.GlobalInit()

	// 设置运行环境
	gin.SetMode(global.Config.Server.Level)

	router.Run(":" + global.Config.Server.Port)
}
