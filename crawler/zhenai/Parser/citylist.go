package Parser

import (
	"fmt"
	"learngo/crawler/engine"
	"learngo/crawler_distributer/config"
	"regexp"
)

const cityListRe = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Requests{
			Url: string(m[1]),
			Parser: engine.NewFuncParser(
				ParseCity, config.ParseCity),
		})
		fmt.Printf("city: %s, url: %s", m[2], m[1])
		fmt.Println()
	}
	return result

}
