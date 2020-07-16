package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ma-novel-crawler/service"
	"net/http"
)

type NovelController struct {
	NovelService service.INovelService
}

// @Description 生成单本小说txt
// @Tags article
// @Produce  json
// @Param url query string true " "
// @Success 200
// @Failure 500  "error info"
// @Router /novel-single-txt [get]
func (ctrl *NovelController) GenNovelSingleTxt(ctx *gin.Context) {
	url := ctx.Query("url")
	bookName,content,err := ctrl.NovelService.GetSingleNovel(url)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	contentType := "application/octet-stream;charset=UTF-8"
	extraHeaders := map[string]string{
		"Content-Disposition": fmt.Sprintf("attachment; filename=%s", bookName+".txt"),
	}

	contentLength := content.Len()
	ctx.DataFromReader(http.StatusOK, int64(contentLength), contentType, content, extraHeaders)
}
