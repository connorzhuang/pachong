package main

import (
	"fmt"
	"learngo/crawler_distributer/config"
	"learngo/crawler_distributer/rpcsupport"
	"learngo/crawler_distributer/worker"

	"testing"
	"time"
)

func TestWorkerServer(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServerRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	req := worker.Request{
		Url: "http://localhost:8080/mock/album.zhenai.com/u/8256018539338750764",
		Parser: worker.SerializedParser{
			Name: config.ProfileParser,
			Args: "寂寞成影萌宝",
		},
	}
	result := worker.ParserResult{}
	err = client.Call("CrawlService.Process", req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}

}
