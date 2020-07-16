package biquge

import (
	"github.com/antchfx/htmlquery"
	"ma-novel-crawler/parser"
	"strings"
)

type HomeParser struct {
}

func NewHomeParse() *HomeParser {
	return &HomeParser{}
}

func (p *HomeParser) Parse(contents []byte, url string) parser.ParseResult {
	result := parser.ParseResult{}
	root, _ := htmlquery.Parse(strings.NewReader(string(contents)))
	list := htmlquery.Find(root, "//div[@class='nav']/ul/li")
	for _, row := range list {
		linkNode := htmlquery.FindOne(row, "./a")
		if linkNode != nil {
			link := htmlquery.SelectAttr(linkNode, "href")
			if link == "/" {
				continue
			}
			if !strings.HasPrefix(link, "http") && !strings.HasPrefix(link, "https") {
				link = url + link
			}
			request := parser.Request{Url: link, Parser: NewNovelListParser(), BaseUrl: url}
			result.Requests = append(result.Requests, request)
			linkInfo := parser.LinkInfo{Url: link, Info: htmlquery.InnerText(linkNode)}
			result.LinkInfos = append(result.LinkInfos, linkInfo)
		}
	}
	return result
}

func (p *HomeParser) Serialize() (name string, args interface{}) {
	return "biquge_home_parser", nil
}
