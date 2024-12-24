package initialize

import (
	"ShortService/src/internal/router"
	"github.com/gin-gonic/gin"
)

func routerInit() *gin.Engine {
	r := gin.Default()
	allRouter := router.AllRouter

	{
		allRouter.ServerRouter.InitApiRouter(r)
	}
	return r
}
