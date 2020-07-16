package biquge

import (
	"github.com/antchfx/htmlquery"
	"ma-novel-crawler/parser"
	"strings"
)

type NovelListParser struct {
}

func NewNovelListParser() *NovelListParser {
	return &NovelListParser{}
}

func (p *NovelListParser) Parse(contents []byte, url string) parser.ParseResult {
	result := parser.ParseResult{}
	root, _ := htmlquery.Parse(strings.NewReader(string(contents)))
	list := htmlquery.Find(root, "//div[@id='hotcontent']/div/div/dl/dt")
	for _, row := range list {
		var link, title string
		linkNode := htmlquery.FindOne(row, "./a")
		if linkNode != nil {
			link = htmlquery.SelectAttr(linkNode, "href")
			if !strings.HasPrefix(link, "http") && !strings.HasPrefix(link, "https") {
				link = url + link
			}
			title = htmlquery.InnerText(linkNode)
		}

		request := parser.Request{Url: link, Parser: NewChapterListParser(), BaseUrl: url}
		result.Requests = append(result.Requests, request)
		linkInfo := parser.LinkInfo{Url: link, Info: title}
		result.LinkInfos = append(result.LinkInfos, linkInfo)
	}

	list = htmlquery.Find(root, "//div[@id='newscontent']/div/ul/li")
	for _, row := range list {
		var link, title string
		spanNodes := htmlquery.Find(row, "./span")
		if len(spanNodes) > 0 {
			linkNode := htmlquery.FindOne(spanNodes[0], "./a")
			if linkNode != nil {
				link = htmlquery.SelectAttr(linkNode, "href")
				if !strings.HasPrefix(link, "http") && !strings.HasPrefix(link, "https") {
					link = url + link
				}
				title = htmlquery.InnerText(linkNode)
			}
		}
		request := parser.Request{Url: link, Parser: parser.NilParser{}}
		result.Requests = append(result.Requests, request)
		linkInfo := parser.LinkInfo{Url: link, Info: title}
		result.LinkInfos = append(result.LinkInfos, linkInfo)
	}

	return result
}

func (p *NovelListParser) Serialize() (name string, args interface{}) {
	return "biquge_novel_list_parser", nil
}
