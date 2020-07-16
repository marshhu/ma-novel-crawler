package router

import (
	"github.com/gin-gonic/gin"
	"ma-novel-crawler/api"
)

func userBizRouters(eng *gin.Engine) {
	apiV1 := eng.Group("/api/v1")

	apiV1.GET("/novel-single-txt", api.CtrlFactoryInstance.NovelCtrl.GenNovelSingleTxt)
}
