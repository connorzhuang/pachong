package Parser

import (
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"learngo/crawler_distributer/config"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(
	`<td><span class="label">年龄：</span>(\d+)岁</td>`)
var heightRe = regexp.MustCompile(
	`<td><span class="label">身高：</span>(\d+)CM</td>`)
var incomeRe = regexp.MustCompile(
	`<td><span class="label">月收入：</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(
	`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
var genderRe = regexp.MustCompile(
	`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var xinzuoRe = regexp.MustCompile(
	`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var marriageRe = regexp.MustCompile(
	`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(
	`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(
	`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
var hokouRe = regexp.MustCompile(
	`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(
	`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(
	`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var guessRe = regexp.MustCompile(` <a class="exp-user-name" [^>]*href="(http://localhost:8080/mock/album.zhenai.com/u/[\d]+)">([^<]+)</a>`)
var idUrlRe = regexp.MustCompile(`http://localhost:8080/mock/album.zhenai.com/u/([\d]+)`)

func parseProfile(
	contents []byte, url string, name string) engine.ParserResult {
	profile := model.Profile{}
	profile.Name = name
	age, err := strconv.Atoi(
		extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(
		extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(
		extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}
	profile.Car = extractString(
		contents, carRe)
	profile.Income = extractString(
		contents, incomeRe)
	profile.Gender = extractString(
		contents, genderRe)
	profile.Education = extractString(
		contents, educationRe)
	profile.Hokou = extractString(
		contents, hokouRe)
	profile.House = extractString(
		contents, houseRe)
	profile.Marriage = extractString(
		contents, marriageRe)
	profile.Occupation = extractString(
		contents, occupationRe)
	profile.Xinzuo = extractString(
		contents, xinzuoRe)
	result := engine.ParserResult{
		Items: []engine.Item{
			{
				Url:      url,
				Type:     "zhenai",
				Id:       extractString([]byte(url), idUrlRe),
				PayRound: profile,
			},
		},
	}
	match := guessRe.FindAllSubmatch(contents, -1)
	for _, m := range match {
		// name := string(m[2])  等于拷贝
		result.Requests = append(result.Requests, engine.Requests{
			Url:    string(m[1]),
			Parser: NewProfileParser(string(m[2])),
		})
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}

}

type ProfileParser struct {
	username string
}

func (p *ProfileParser) Parser(contents []byte, url string) engine.ParserResult {
	return parseProfile(contents, url, p.username)
}

func (p *ProfileParser) SerializedParser() (name string, args interface{}) {
	return config.ProfileParser, p.username
}

func NewProfileParser(username string) *ProfileParser {
	return &ProfileParser{
		username: username,
	}

}
