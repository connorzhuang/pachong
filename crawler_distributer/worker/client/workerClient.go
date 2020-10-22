package client

import (
	"learngo/crawler/engine"
	"learngo/crawler_distributer/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	return func(requests engine.Requests) (engine.ParserResult, error) {
		sReq := worker.SerializeRequest(requests)
		var result worker.ParserResult
		client := <-clientChan
		client.Call("CrawlService.Process", sReq, &result)

		return worker.DeserializeParserResult(result), nil

	}
}
