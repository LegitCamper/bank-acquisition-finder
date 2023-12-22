package scrapper

import (
	"fmt"
	"sync"

	"github.com/gocolly/colly/v2"
)

type Site string

const (
	CNBC Site = "www.cnbc.com"
	NPR       = "www.npr.org"
)

func Main() {
	news_sources := []Site{NPR} // ,CNBC
	scrape(news_sources)
}

func scrape(domains []Site) {
	var wg sync.WaitGroup

	for _, element := range domains {
		wg.Add(1)
		scraper(&wg, element)
	}
	// Wait for all sites to complete search
	wg.Wait()

	fmt.Println("Done")
}

func scraper(wg *sync.WaitGroup, domain Site) {
	defer wg.Done()
	c := colly.NewCollector(
		colly.AllowedDomains(string(domain)),
	)

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	//https://oxylabs.io/blog/golang-web-scraper
	switch domain {
	case CNBC:
		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			fmt.Println(e.Text)
		})
	case NPR:
		c.OnHTML("p", func(e *colly.HTMLElement) {
			fmt.Println(e.Text)
		})
	}

	website := "https://" + string(domain)
	c.Visit(website)

	fmt.Println("Finished ", domain)
}
