package biquge

import (
	"fmt"
	"ma-novel-crawler/fetcher"
	"ma-novel-crawler/parser/biquge"
	"net/http"
	"testing"
)

func Test_HomeParser(t *testing.T) {
	url :="https://www.biquge.com.cn/"
	status, contents, err := fetcher.Fetcher(url, "", 5)
	if err != nil {
		t.FailNow()
	}
	if status != http.StatusOK {
		t.FailNow()
	}
	parser := biquge.NewHomeParse()
	result,err := parser.Parse(url,contents)
	if err!=nil{
		t.FailNow()
	}
	fmt.Println(result)
	if len(result.Requests) != 10 {
		t.FailNow()
	}
}
