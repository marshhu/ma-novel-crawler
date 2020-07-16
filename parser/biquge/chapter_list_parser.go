package biquge

import (
	"github.com/antchfx/htmlquery"
	"ma-novel-crawler/parser"
	"ma-novel-crawler/parser/model"
	"strings"
)

type ChapterListParser struct {
}

func NewChapterListParser() *ChapterListParser {
	return &ChapterListParser{}
}

func (p *ChapterListParser) Parse(contents []byte, url string) parser.ParseResult {
	result := parser.ParseResult{}
	root, _ := htmlquery.Parse(strings.NewReader(string(contents)))
	//获取小说信息
	bookInfo := model.BookInfo{}
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
	var chapters []model.BookChapter
	index := 0
	for _, row := range list {
		index++
		var link, text string
		linkNode := htmlquery.FindOne(row, "./a")
		if linkNode != nil {
			link = htmlquery.SelectAttr(linkNode, "href")
			if !strings.HasPrefix(link, "http") && !strings.HasPrefix(link, "https") {
				link = url + link
			}
			text = htmlquery.InnerText(linkNode)
		}

		request := parser.Request{Url: link, Parser: NewChapterDetailParser()}
		result.Requests = append(result.Requests, request)
		chapter := model.BookChapter{Index: index, Name: text}
		chapters = append(chapters, chapter)
		linkInfo := parser.LinkInfo{Url: link, Info: chapter}
		result.LinkInfos = append(result.LinkInfos, linkInfo)
	}
	//fmt.Println(chapters)
	return result
}

func (p *ChapterListParser) Serialize() (name string, args interface{}) {
	return "biquge_chapter_list_parser", nil
}
