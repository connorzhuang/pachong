package engine

import (
	"log"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seed ...Requests) {
	var requests []Requests
	for _, r := range seed {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parserResult, err := Worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parserResult.Requests...)
		for _, item := range parserResult.Items {
			log.Printf("Got Items:%v", item)
		}
	}
}
