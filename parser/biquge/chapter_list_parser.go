package biquge

import (
	"github.com/antchfx/htmlquery"
	"ma-novel-crawler/parser"
	"net/url"
	"strings"
)

type ChapterListParser struct {
}

func NewChapterListParser() *ChapterListParser {
	return &ChapterListParser{}
}

func (p *ChapterListParser) Parse(crawlerUrl string,contents []byte) (*parser.ParseResult,error) {
	u,err:= url.Parse(crawlerUrl)
	if err!= nil{
		return nil,err
	}
	result := &parser.ParseResult{}
	result.Requests = make(map[string]parser.UrlParser)

	root, _ := htmlquery.Parse(strings.NewReader(string(contents)))
	//获取小说信息
	bookInfo := parser.BookInfo{}
	findNode := htmlquery.FindOne(root, "//div[@id='info']")
	if findNode != nil {
		h1Node := htmlquery.FindOne(findNode, "./h1")
		if h1Node != nil {
			bookInfo.BookName = htmlquery.InnerText(h1Node)
		}
		pNodes := htmlquery.Find(findNode, "./p")
		if len(pNodes) >= 4 {
			bookInfo.BookAuthor = htmlquery.InnerText(pNodes[0])
			bookInfo.UpdateTime = htmlquery.InnerText(pNodes[2])
			bookInfo.LatestChapter = htmlquery.InnerText(pNodes[3])
		}
	}
	findNode = htmlquery.FindOne(root, "//div[@id='intro']")
	if findNode != nil {
		bookInfo.BookIntro = htmlquery.InnerText(findNode)
	}
	findNode = htmlquery.FindOne(root, "//div[@id='fmimg']/img")
	if findNode != nil {
		bookInfo.BookImage = htmlquery.SelectAttr(findNode, "src")
	}
	//fmt.Println(bookInfo)
	result.Data = bookInfo

	list := htmlquery.Find(root, "//div[@id='list']/dl/dd")
	for _, row := range list {
		linkNode := htmlquery.FindOne(row, "./a")
		if linkNode != nil {
			link := htmlquery.SelectAttr(linkNode, "href")
			if !strings.HasPrefix(link, "http") && !strings.HasPrefix(link, "https") {
				link = u.Scheme+"://"+u.Host + link
			}
			result.Requests[link] =parser.UrlParser{ Url:link,Parser: NewChapterDetailParser(),UrlText:htmlquery.InnerText(linkNode)}
		}
	}
	//fmt.Println(chapters)
	return result,nil
}

func (p *ChapterListParser) Serialize() (name string, args interface{}) {
	return "biquge_chapter_list_parser", nil
}
