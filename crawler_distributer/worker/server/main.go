package main

import (
	"flag"
	"fmt"
	"learngo/crawler_distributer/rpcsupport"
	"learngo/crawler_distributer/worker"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		log.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServerRpc(
		fmt.Sprintf(":%d", *port),
		worker.CrawlService{}))

}
