package engine

import (
	"log"
	"reptile/fetcher"
)

func Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		log.Printf("fetcher url %s",r.Url)
		requests = requests[1:]
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("fetcher: err "+"fetcher url %s: %v", r.Url, err)
			continue
		}
		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}

	}

}
