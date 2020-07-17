package biquge

import (
	"github.com/antchfx/htmlquery"
	"ma-novel-crawler/parser"
	"net/url"
	"regexp"
	"strings"
)

type ChapterDetailParser struct {
}

func NewChapterDetailParser() *ChapterDetailParser {
	return &ChapterDetailParser{}
}
func (p *ChapterDetailParser) Parse(crawlerUrl string,contents []byte) (*parser.ParseResult,error) {
	_,err:= url.Parse(crawlerUrl)
	if err!= nil{
		return nil,err
	}
	result := &parser.ParseResult{}
	result.Requests = make(map[string]parser.UrlParser)

	root, _ := htmlquery.Parse(strings.NewReader(string(contents)))
	findNode := htmlquery.FindOne(root, "//div[@id='content']")
	var chapterContent string
	if findNode != nil {
		chapterContent = htmlquery.InnerText(findNode)
		//替换空格&nbsp
		nbspRg := regexp.MustCompile(`&nbsp;`)
		chapterContent = nbspRg.ReplaceAllString(chapterContent, "")
		//替换换行<br>为\n
		brRg := regexp.MustCompile(`<br>`)
		chapterContent = brRg.ReplaceAllString(chapterContent, "\n")
		//fmt.Println(chapterContent)
		result.Data = chapterContent
	}
	return result,nil
}

func (p *ChapterDetailParser) Serialize() (name string, args interface{}) {
	return "biquge_chapter_detail_parser", nil
}
