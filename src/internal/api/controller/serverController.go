package controller

import (
	"ShortService/src/common"
	"ShortService/src/common/e"
	"ShortService/src/global"
	"ShortService/src/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ShortController struct {
	shortService service.IShortService
}

func NewShortController(service service.IShortService) *ShortController {
	return &ShortController{shortService: service}
}

// 解析端短链接
func (cc *ShortController) ParseUrl(ctx *gin.Context) {
	shortUrl := ctx.Param("shortCode")
	// 解析出原始长链接
	redictUrl, err := cc.shortService.ParseUrl(ctx, shortUrl)
	if err != nil {
		global.Log.Info("解析短链接:%s失败", shortUrl)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	// 重定向到redictUrl
	ctx.Redirect(http.StatusFound, redictUrl)
	global.Log.Info(fmt.Sprintf("解析短链接:%s, 原始链接:%s", shortUrl, redictUrl))
}

// 创建短链接
func (cc *ShortController) CreateShortUrl(ctx *gin.Context) {
	url := ctx.DefaultPostForm("url", "")

	shortUrl, err := cc.shortService.CreateShort(ctx, url)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, common.Result{
		Code: e.SUCCESS,
		Msg:  shortUrl,
	})
	//global.Log.Info("创建短链接:" + url)
}
