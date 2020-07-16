package service

import (
	"bytes"
	"errors"
	"ma-novel-crawler/fetcher"
	"ma-novel-crawler/parser"
	"ma-novel-crawler/parser/biquge"
	"ma-novel-crawler/parser/model"
	"net/http"
	"net/url"
)

type NovelService struct {

}

func(s *NovelService)GetSingleNovel(novelUrl string) (bookName string,content *bytes.Buffer,err error){
	u,err := url.Parse(novelUrl)
	if err !=nil{
		return "",nil,err
	}
	request := parser.Request{Url: novelUrl, Parser: biquge.NewChapterListParser(), BaseUrl: u.Scheme+"://"+u.Host}
	status, contents, err := fetcher.Fetcher(request.Url, "", 5)
	if err != nil {
		return "",nil,errors.New("访问站点失败")
	}
	if status != http.StatusOK {
		return "",nil,errors.New("访问站点失败")
	}
	bookParseResult := request.Parser.Parse(contents, request.BaseUrl)
	bookInfo,ok := bookParseResult.Data.(model.BookInfo)
	if !ok{
		return "",nil,errors.New("解析小说信息失败")
	}
	bookName = bookInfo.BookName
	buf := new(bytes.Buffer)
	for index,chapterReq := range bookParseResult.Requests{
		status, contents, err := fetcher.Fetcher(chapterReq.Url, "", 5)
		if err != nil {
			return "",nil,errors.New("访问章节站点失败")
		}
		if status != http.StatusOK {
			return "",nil,errors.New("访问章节站点失败")
		}
		chapterParseResult := chapterReq.Parser.Parse(contents, chapterReq.BaseUrl)
		chapterContent := chapterParseResult.Data.(string)
		chapterInfo := bookParseResult.LinkInfos[index]
		chapter :=chapterInfo.Info.(model.BookChapter)

		buf.WriteString(chapter.Name+"\n\n")
		buf.WriteString(chapterContent+"\n")
		break
	}

	return bookName,buf,nil
}
