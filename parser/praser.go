package parser

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type Request struct {
	BaseUrl string
	Url     string
	Parser  Parser
}

type ParseResult struct {
	Requests  []Request
	LinkInfos []LinkInfo
	Data      interface{}
}

type LinkInfo struct {
	Url  string
	Info interface{}
}

type NilParser struct{}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}
