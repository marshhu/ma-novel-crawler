package router

import (
	"github.com/gin-gonic/gin"
	"github.com/marshhu/ma-novel-crawler/api"
)

func userBizRouters(eng *gin.Engine) {
	apiV1 := eng.Group("/api/v1")

	apiV1.GET("/novels", api.CtrlFactoryInstance.NovelCtrl.GenNovelSingleTxt)
}
