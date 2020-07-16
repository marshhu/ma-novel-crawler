package biquge

import (
	"fmt"
	"ma-novel-crawler/fetcher"
	"ma-novel-crawler/parser"
	"ma-novel-crawler/parser/biquge"
	"net/http"
	"testing"
)

func Test_ChapterListParser(t *testing.T) {
	request := parser.Request{Url: "https://www.biquge.com.cn/book/43108/", Parser: biquge.NewChapterListParser(), BaseUrl: "https://www.biquge.com.cn"}
	status, contents, err := fetcher.Fetcher(request.Url, "", 5)
	if err != nil {
		t.FailNow()
	}
	if status != http.StatusOK {
		t.FailNow()
	}
	result := request.Parser.Parse(contents, request.BaseUrl)
	fmt.Println("findCount:", len(result.LinkInfos))
	fmt.Println(result.LinkInfos)
}
