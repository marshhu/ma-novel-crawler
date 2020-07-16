package service

import "bytes"

type INovelService interface {
	GetSingleNovel(url string) (bookName string,content *bytes.Buffer,err error)
}
