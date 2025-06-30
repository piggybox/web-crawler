package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	seedurl := "https://everlane.com"

	crawl(seedurl)
}

func crawl(url string) {
	c := colly.NewCollector(
		colly.AllowedDomains("everlane.com", "www.everlane.com"),
		colly.MaxDepth(2),
	)

	c.OnHTML("title", func(e *colly.HTMLElement) {
		title := e.Text
		fmt.Println("Found title:", title)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error visiting:", r.Request.URL, "Error:", err)
	})

	err := c.Visit(url)
	if err != nil {
		fmt.Println("Error starting crawl:", err)
	}
}