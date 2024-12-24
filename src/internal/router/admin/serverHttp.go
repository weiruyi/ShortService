package admin

import (
	"ShortService/src/common"
	"ShortService/src/common/e"
	"ShortService/src/global"
	"ShortService/src/internal/api/controller"
	"ShortService/src/internal/repository/dao"
	"ShortService/src/internal/service"
	"net/http"

	//"ShortService/src/internal/api/controller"
	//"ShortService/src/internal/repository/dao"
	//"ShortService/src/internal/service"
	"github.com/gin-gonic/gin"
	//"take-out/internal/service"
	//"take-out/middle"
)

type ServerRouter struct{}

func (cr *ServerRouter) InitApiRouter(parent *gin.Engine) {

	//// 依赖注入
	serverCtrl := controller.NewShortController(
		service.NewShortService(
			service.NewSequenceService(dao.NewSequenceDao(global.DB)),
			service.NewShortUrlService(dao.NewShortUrlsDao(global.DB)),
		),
	)
	{

		parent.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, common.Result{
				Code: e.SUCCESS,
				Msg:  "Welcome to the short URL system!",
			})
		})

		//创建
		parent.POST("/create", serverCtrl.CreateShortUrl)

		//解析
		parent.GET("/:shortCode", serverCtrl.ParseUrl)

	}
}
