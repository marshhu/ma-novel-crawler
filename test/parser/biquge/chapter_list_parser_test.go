package biquge

import (
	"fmt"
	"ma-novel-crawler/fetcher"
	"ma-novel-crawler/parser/biquge"
	"net/http"
	"testing"
)

func Test_ChapterListParser(t *testing.T) {
	url := "https://www.biquge.com.cn/book/43108/"
	status, contents, err := fetcher.Fetcher(url, "", 5)
	if err != nil {
		t.FailNow()
	}
	if status != http.StatusOK {
		t.FailNow()
	}
	parser := biquge.NewChapterListParser()
	result, err := parser.Parse(url, contents)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}
