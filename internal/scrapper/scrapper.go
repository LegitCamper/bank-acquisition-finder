package scrapper

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func Main() {
	news_sources := []string{
		"www.npr.org",
		"www.cnbc.com",
	}
	scrape(news_sources)

}

func scrape(domains []string) {
	for _, element := range domains {
		scraper(element)
	}

	fmt.Println("Done")
}

func scraper(domain string) {
	c := colly.NewCollector(
		colly.AllowedDomains(domain),
	)

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
		fmt.Println(e)
	})

	//https://oxylabs.io/blog/golang-web-scraper

	website := "https://" + domain
	c.Visit(website)
	fmt.Println("Finished ", domain)

}
