package biquge

import (
	"fmt"
	"github.com/marshhu/ma-novel-crawler/fetcher"
	"github.com/marshhu/ma-novel-crawler/parser/biquge"
	"net/http"
	"testing"
)

func Test_ChapterDetailParser(t *testing.T) {
	url := "https://www.biquge.com.cn/book/43108/348013.html"
	status, contents, err := fetcher.Fetcher(url, "", 5)
	if err != nil {
		t.FailNow()
	}
	if status != http.StatusOK {
		t.FailNow()
	}

	chapterDetailParser := biquge.NewChapterDetailParser(nil)
	chapterDetailResult, err := chapterDetailParser.Parse(url, contents)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(chapterDetailResult)
}
