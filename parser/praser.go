package parser

type Parser interface {
	Parse(crawlerUrl string,contents []byte) (*ParseResult,error)
	Serialize() (name string, args interface{})
}

type ParseResult struct {
	Data      interface{}
	Requests  map[string]UrlParser
}

type UrlParser struct {
	Url string
	UrlText string
	Parser
}

type NilParser struct{}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}
