package service

import (
	"bytes"
	"errors"
	"ma-novel-crawler/fetcher"
	"ma-novel-crawler/parser"
	"ma-novel-crawler/parser/biquge"
	"net/http"
	"net/url"
	"sync"
)

type NovelService struct {

}

func(s *NovelService)GetSingleNovel(novelUrl string) (bookName string,content *bytes.Buffer,err error){
	_,err = url.Parse(novelUrl)
	if err !=nil{
		return "",nil,err
	}
	status, contents, err := fetcher.Fetcher(novelUrl, "", 5)
	if err != nil {
		return "",nil,errors.New("访问站点失败")
	}
	if status != http.StatusOK {
		return "",nil,errors.New("访问站点失败")
	}
	chapterParser := biquge.NewChapterListParser()
	parseResult,err:= chapterParser.Parse(novelUrl,contents)
	if err!=nil{
		return "",nil,err
	}
	bookInfo,ok := parseResult.Data.(parser.BookInfo)
	if !ok{
		return "",nil,errors.New("解析小说信息失败")
	}
	bookName = bookInfo.BookName

	inputChan := make(chan parser.UrlParser,10)
	outputChan := make(chan NovelChapter,10)
	buffer := new(bytes.Buffer)
	var mutex sync.Mutex
	go createJob(parseResult.Requests,inputChan)
	go handleOutPut(outputChan,&mutex,buffer)
	var wg sync.WaitGroup
	workerNum := 20
    for i:=0;i<workerNum;i++{
    	wg.Add(1)
        go worker(inputChan,outputChan,&wg)
	}
	wg.Wait()
    close(outputChan)
	return bookName,buffer,nil
}

func createJob(requests map[string]parser.UrlParser,inputChan chan parser.UrlParser){
	for _,request := range requests{
		inputChan <- request
	}
	close(inputChan)
}

func worker(inputChan <-chan parser.UrlParser,outputChan chan NovelChapter,wg *sync.WaitGroup){
	defer wg.Done()
    for request := range inputChan{
		status, contents, err := fetcher.Fetcher(request.Url, "", 5)
		if err != nil {
			return
		}
		if status != http.StatusOK {
			return
		}
		chapterParseResult,err := request.Parse(request.Url,contents)
		if err!=nil{
			continue
		}
		chapterContent := chapterParseResult.Data.(string)
		outputChan <- NovelChapter{Url: request.Url,Content: chapterContent,ChapterName: request.UrlText}
	}
}

func handleOutPut(outputChan chan NovelChapter,mutex *sync.Mutex,buffer *bytes.Buffer){
	for parseResult := range outputChan{
		mutex.Lock()
		buffer.WriteString(parseResult.ChapterName+"\n\n")
		buffer.WriteString(parseResult.Content+"\n")
		mutex.Unlock()
	}
}