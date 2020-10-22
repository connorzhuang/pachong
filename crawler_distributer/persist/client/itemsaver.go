package client

import (
	"context"
	"errors"
	"learngo/crawler/engine"
	"learngo/crawler_distributer/rpcsupport"
	"log"

	"github.com/olivere/elastic/v7"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver got item"+"#%d: %v", itemCount, item)
			itemCount++
			result := ""
			client.Call("ItemSaverService.Save", item, &result)
			if err != nil {
				log.Printf("Item Saver :error"+
					"saving item %v:%v", item, err)
			}
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, index string, item engine.Item) error {
	if item.Type == "" {
		return errors.New("must supply Type")
	}
	indexServer := client.Index(). //将数据输入到elasticsearch
					Index(index).
					Type(item.Type).
					Id(item.Id).
					BodyJson(item)
	if item.Id != "" {
		indexServer.Id(item.Id)
	}
	_, err := indexServer.Do(context.Background())

	if err != nil {
		return err
	}
	return nil
}
