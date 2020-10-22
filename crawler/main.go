package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/persist"
	"learngo/crawler/scheduler"
	"learngo/crawler/zhenai/Parser"
)

func main() {
	itemChan, err := persist.ItemSaver("database")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    itemChan,
	}
	e.Run(engine.Requests{
		Url: "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(
			Parser.ParseCityList, "ParseCityList"),
	})

	/*e.Run(engine.Requests{
		Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun/aba",
		ParserFunc: Parser.ParseCity,
	})

	*/
}
