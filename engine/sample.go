package engine

import (
	"log"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {

	var requests []Request

	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	for len(requests) > 0 {
		firstRequest := requests[0]
		requests = requests[1:]

		ParseResult, err := Worker(firstRequest)

		if err != nil {
			continue
		}

		requests = append(requests, ParseResult.Requests...)
		for _, item := range ParseResult.Items {
			log.Printf("Got item %+v", item)
		}
	}
}
