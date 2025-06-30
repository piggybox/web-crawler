package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

var visitedURLs = make(map[string]bool)

func main() {
	seedurl := "https://everlane.com"

	crawl(seedurl, 0)
}

func crawl(url string, maxdepth int) {

	c := colly.NewCollector(
		colly.AllowedDomains("everlane.com", "www.everlane.com"),
		colly.MaxDepth(maxdepth),
	)

	c.OnHTML("title", func(e *colly.HTMLElement) {
		title := e.Text
		fmt.Println("Found title:", title)
	})

	// find and visit all outlinks
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if link != "" && !visitedURLs[link] {
			visitedURLs[link] = true
			fmt.Println("Found link:", link)

			err := e.Request.Visit(link)
			if err != nil {
				fmt.Println("Error visiting link:", link, "Error:", err)
			}
		}
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
