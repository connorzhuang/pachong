package main

import (
	"learngo/crawler_distributer/persist"
	"learngo/crawler_distributer/rpcsupport"

	"github.com/olivere/elastic/v7"
)

func main() {
	err := ServerRpc(":1234", "database")
	if err != nil {
		panic(err)
	}

}

func ServerRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return rpcsupport.ServerRpc(host, persist.ItemSaverService{
		client,
		index,
	})
}
