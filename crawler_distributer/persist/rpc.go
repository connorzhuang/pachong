package persist

import (
	"learngo/crawler/engine"
	"learngo/crawler/persist"
	"log"

	"github.com/olivere/elastic/v7"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	log.Printf("item %v saved", item)
	if err == nil {
		*result = "ok"
	}
	return err
}
