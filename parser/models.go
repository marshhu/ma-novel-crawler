package parser

type BookInfo struct {
	BookName      string `json:"bookName"`
	BookAuthor    string `json:"bookAuthor"`
	BookImage     string `json:"bookImage"`
	BookIntro     string `json:"bookIntro"`
	LatestChapter string `json:"latestChapter"`
	UpdateTime    string `json:"update_time"`
}
