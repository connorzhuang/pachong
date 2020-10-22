package main

import (
	"flag"
	"learngo/crawler/engine"
	"learngo/crawler/scheduler"
	"learngo/crawler/zhenai/Parser"
	"learngo/crawler_distributer/config"
	"learngo/crawler_distributer/persist/client"
	"learngo/crawler_distributer/rpcsupport"
	client2 "learngo/crawler_distributer/worker/client"
	"log"
	"net/rpc"
	"strings"
)

var (
	workerHosts = flag.String("worker_hosts", "", "worker hosts(comma separated)")
)

func main() {
	flag.Parse()
	itemChan, err := client.ItemSaver(":1234")
	if err != nil {
		panic(err)
	}
	pool := CreateClientPool(strings.Split(*workerHosts, ","))
	processor := client2.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Requests{
		Url:    "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(Parser.ParseCityList, config.ParseCityList),
	})

	/*e.Run(engine.Requests{
		Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun/aba",
		ParserFunc: Parser.ParseCity,
	})

	*/
}

func CreateClientPool(hosts []string) chan *rpc.Client {

	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("connecting to %s", h)
		} else {
			log.Printf("error connecting to %s", err)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, c := range clients {
				out <- c
			}
		}
	}()

	return out
}
