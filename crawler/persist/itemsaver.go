package persist

import (
	"context"
	"errors"
	"learngo/crawler/engine"
	"log"

	"github.com/olivere/elastic/v7"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false)) //连接elasticSearch
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver got item"+"#%d: %v", itemCount, item)
			itemCount++
			err := Save(client, index, item)
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
