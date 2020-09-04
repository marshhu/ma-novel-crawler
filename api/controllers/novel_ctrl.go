package controllers

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marshhu/ma-novel-crawler/service"
	"net/http"
	"net/url"
)

type NovelController struct {
	NovelService service.INovelService
}

// @Description 生成单本小说txt
// @Tags novels
// @Produce  json
// @Param novelUrl query string true " "
// @Success 200
// @Failure 500  "error info"
// @Router /novels [get]
func (ctrl *NovelController) GenNovelSingleTxt(ctx *gin.Context) {
	novelUrl := ctx.Query("novelUrl")
	novel, err := ctrl.NovelService.GetNovelByUrl(novelUrl)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var buffer bytes.Buffer
	for _, chapter := range novel.Chapters {
		buffer.WriteString(chapter.Name + "\n\n")
		buffer.WriteString(chapter.Content + "\n\n")
	}
	contentType := "application/octet-stream;charset=UTF-8"

	extraHeaders := map[string]string{
		"Content-Disposition": fmt.Sprintf("attachment; filename=%s", url.QueryEscape(novel.Name+".txt")),
	}

	contentLength := buffer.Len()
	ctx.DataFromReader(http.StatusOK, int64(contentLength), contentType, &buffer, extraHeaders)
	ctx.JSON(http.StatusOK, novel)
}
