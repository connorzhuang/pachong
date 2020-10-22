package engine

import "learngo/crawler_distributer/config"

type Requests struct {
	Url    string
	Parser Parser
}
type ParserFunc func([]byte, string) ParserResult

type Parser interface {
	Parser([]byte, string) ParserResult
	SerializedParser() (name string, args interface{}) //序列化和反序列化
}

type ParserResult struct {
	Requests []Requests
	Items    []Item
}
type Item struct {
	Url      string
	Type     string
	Id       string
	PayRound interface{}
}

type NilParser struct{}

func (n *NilParser) Parser(_ []byte, _ string) ParserResult {
	return ParserResult{}
}

func (n *NilParser) SerializedParser() (name string, args interface{}) {
	return config.NilParser, nil
}

type FuncParser struct {
	parse ParserFunc
	name  string
}

func (f *FuncParser) Parser(contents []byte, url string) ParserResult {
	return f.parse(contents, url)
}

func (f *FuncParser) SerializedParser() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parse: p,
		name:  name,
	}

}
