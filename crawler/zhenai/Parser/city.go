package Parser

import (
	"learngo/crawler/engine"
	"learngo/crawler_distributer/config"
	"regexp"
)

var profileRE = regexp.MustCompile(`<a href="(http://localhost:8080/mock/album.zhenai.com/u/[\d]+)">([^<]+)</a>`)
var profileUrlRe = regexp.MustCompile(`href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[^"]+)"`)

func ParseCity(contents []byte, _ string) engine.ParserResult {
	match := profileRE.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, m := range match {
		result.Requests = append(result.Requests, engine.Requests{
			Url:    string(m[1]),
			Parser: NewProfileParser(string(m[2])),
		})
	}

	matches := profileUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Requests{
			Url: string(m[1]),
			Parser: engine.NewFuncParser(
				ParseCity, config.ParseCity),
		})
	}
	return result

}
