package initialize

import (
	"ShortService/src/config"
	"ShortService/src/global"
	"ShortService/src/logger"
	"github.com/gin-gonic/gin"
)

func GlobalInit() *gin.Engine {
	// 配置文件初始化
	global.Config = config.InitLoadConfig()
	//// Log初始化
	global.Log = logger.NewMySlog(global.Config.Log.Level, global.Config.Log.FilePath)
	//// Gorm初始化
	global.DB = InitDatabase(global.Config.DataSource.Dsn())
	//// Redis初始化
	global.Redis = initRedis()
	//布隆过滤器
	global.Bloom = initBloom()
	// Router初始化
	router := routerInit()
	return router
}
