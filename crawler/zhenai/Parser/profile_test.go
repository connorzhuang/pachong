package Parser

import (
	"io/ioutil"
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"testing"
)

func TestParseProfile(t *testing.T) {
	content, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := parseProfile(content, "http://localhost:8080/mock/album.zhenai.com/u/8256018539338750764", "寂寞成影萌宝")
	if len(result.Items) != 1 {
		t.Errorf("expected %d ,but was %d", 1, len(result.Items))
	}
	act := result.Items[0]

	expected := engine.Item{
		Url:  "http://localhost:8080/mock/album.zhenai.com/u/8256018539338750764",
		Type: "zhenai",
		Id:   "8256018539338750764",
		PayRound: model.Profile{
			Name:       "寂寞成影萌宝",
			Gender:     "女",
			Age:        83,
			Height:     105,
			Weight:     137,
			Income:     "财务自由",
			Car:        "无车",
			Marriage:   "离异",
			Education:  "初中",
			Occupation: "金融",
			Hokou:      "南京市",
			Xinzuo:     "狮子座",
			House:      "无房",
		},
	}

	if expected != act {
		t.Errorf("expected %v,but was %v", expected, act)
	}

}
