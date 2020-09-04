package service

import (
	"github.com/marshhu/ma-novel-crawler/service"
	"testing"
)

func Test_GetSingleNovel(t *testing.T) {
	url := "https://www.biquge.com.cn/book/44060/"
	novelService := service.NovelService{}
	novel, err := novelService.GetSingleNovel(url)
	if err != nil {
		t.FailNow()
	}
	if novel.Name != "万族之劫" {
		t.FailNow()
	}
	if novel.Author != "老鹰吃小鸡" {
		t.FailNow()
	}
}
