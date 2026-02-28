package main

import (
	"fmt"
	"os"

	"github.com/joycaleb9705/go-lld/webscraper/crawler"
)

func main() {
	urls := os.Args[1:]
	if len(urls) == 0 {
		fmt.Println("Usage: go run . <url1> <url2> ...")
		return
	}

	fmt.Printf("Fetching %d URLS:\n", len(urls))

	for result := range crawler.CrawlAll(urls) {
		if result.Error != nil {
			fmt.Printf("%s: %v\n", result.URL, result.Error.Error())
		} else {
			fmt.Printf("%s: %s (%f)\n", result.URL, result.Title, result.Duration.Seconds())
		}
	}
}
