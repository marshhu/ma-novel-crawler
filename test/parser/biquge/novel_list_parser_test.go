package biquge

import (
	"fmt"
	"github.com/marshhu/ma-novel-crawler/fetcher"
	"github.com/marshhu/ma-novel-crawler/parser/biquge"
	"net/http"
	"testing"
)

func Test_NovelListParser(t *testing.T) {
	url := "https://www.biquge.com.cn/xuanhuan/"
	status, contents, err := fetcher.Fetcher(url, "", 5)
	if err != nil {
		t.FailNow()
	}
	if status != http.StatusOK {
		t.FailNow()
	}
	parser := biquge.NewNovelListParser()
	result, err := parser.Parse(url, contents)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)

}
