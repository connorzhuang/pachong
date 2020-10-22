package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"learngo/crawler_distributer/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {

	go ServerRpc(":1234", "test1")

	time.Sleep(time.Second)
	newClient, err := rpcsupport.NewClient(":1234")
	if err != nil {
		panic(err)
	}
	item := engine.Item{
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
	var result = ""
	err = newClient.Call("ItemSaverService.Save", item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result:%s,err :%s", result, err)
	}
}
