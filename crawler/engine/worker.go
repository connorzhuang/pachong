package engine

import (
	"learngo/crawler/Fetcher"
	"log"
)

func Worker(r Requests) (ParserResult, error) {
	// log.Printf("Fetching %s", r.Url)
	body, err := Fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher : error "+
			"fetching url %s: %v", r.Url, err)

	}
	return r.Parser.Parser(body, r.Url), nil

}
