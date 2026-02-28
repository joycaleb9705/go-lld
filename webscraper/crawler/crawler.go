package crawler

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/html"
)

type CrawlResult struct {
	URL      string
	Title    string
	Duration time.Duration
	Error    error
}

var ErrTitleNotFound error = fmt.Errorf("Title not found")

var client = &http.Client{
	Timeout: 10 * time.Second,
}

func CrawlAll(urls []string) <-chan CrawlResult {
	var wg sync.WaitGroup
	c := make(chan CrawlResult, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go Crawl(url, &wg, c)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	return c
}

func Crawl(url string, wg *sync.WaitGroup, results chan<- CrawlResult) {
	defer wg.Done()
	start := time.Now()

	// make the GET request
	resp, err := client.Get(url)
	if err != nil {
		results <- CrawlResult{URL: url, Error: err}
		return
	}
	// defer closing the TCP connection
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		results <- CrawlResult{URL: url, Error: fmt.Errorf("status code: %d", resp.StatusCode)}
	}

	// parse the response body
	doc, err := html.Parse(resp.Body)
	if err != nil {
		results <- CrawlResult{URL: url, Error: err}
		return
	}

	// fetch title
	if title, ok := fetchTitle(doc); ok {
		results <- CrawlResult{
			URL:      url,
			Title:    title,
			Duration: time.Since(start),
			Error:    nil,
		}
	} else {
		results <- CrawlResult{URL: url, Error: ErrTitleNotFound}
	}
}

// travers the HTML tree to fetch the title
func fetchTitle(doc *html.Node) (string, bool) {
	if isTitleElement(doc) {
		if doc.FirstChild != nil {
			return doc.FirstChild.Data, true
		}
		return "", false
	}

	for child := doc.FirstChild; child != nil; child = child.NextSibling {
		title, ok := fetchTitle(child)
		if ok {
			return title, true
		}
	}

	return "", false
}

func isTitleElement(doc *html.Node) bool {
	return doc.Type == html.ElementNode && doc.Data == "title"
}
