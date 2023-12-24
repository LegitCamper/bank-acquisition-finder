package scrapper

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Domain string
type Filter struct {
	allowed    string
	disallowed string
}
type Site struct {
	domain Domain
	filter Filter
}

const (
	CNBC Domain = "www.cnbc.com"
	NPR         = "www.npr.org"
)

func Main() {
	// filter ex. ".*" + string(NPR) + "/" + strconv.Itoa(time.Now().Year()) + ".*"
	var (
		// CNBC = Site{CNBC, Filter{"", ""}}
		NPR = Site{NPR, Filter{"", "/podcasts/|/series/|/programs/|/people/|/sections/|/local/|/tags/|/templates/|/corrections/|/about-npr/"}}
	)
	fmt.Println(NPR.filter.allowed)

	news_sources := []Site{NPR} // CNBC,
	scrape(news_sources)
}

func scrape(domains []Site) {
	var wg sync.WaitGroup

	for _, element := range domains {
		wg.Add(1)
		go scraper(&wg, element)
	}
	// Wait for all sites to complete search
	wg.Wait()

	fmt.Println("Done")
}

func scraper(wg *sync.WaitGroup, site Site) {
	defer wg.Done()
	c := colly.NewCollector(
		colly.AllowedDomains(string(site.domain)),
		colly.URLFilters(regexp.MustCompile(site.filter.allowed)),
		colly.DisallowedURLFilters(regexp.MustCompile(site.filter.disallowed)),
	)

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})
	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	// Match only the correct elements per site
	switch site.domain {
	case CNBC:
		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			fmt.Println(e.Text)
		})
	case NPR:
		c.OnHTML("article.story", func(e *colly.HTMLElement) {
			e.ForEach("p", func(i int, h *colly.HTMLElement) {
				url := e.Request.URL.String()
				if strings.Contains(url, strconv.Itoa(time.Now().Year())) || strings.Contains(url, "transcripts") {
					text := strings.Join(strings.Fields(strings.TrimSpace(e.Text)), " ")
					fmt.Println(text)
				}
			})
		})
	}

	website := "https://" + string(site.domain)
	c.Visit(website)

	fmt.Println("Finished ", site.domain)
}

func bank_acquisition_check(s string) {
	banks := []string{
		"CVB Financial",
		"First Financial Bankshares",
		"Enterprise Financial Services",
		"Columbia",
		"OZK",
		"Prosperity Bancshares",
		"Home BancShares",
		"East West Bancorp",
		"Pacific Premier Bancorp",
		"Cathay General Bancorp",
		"Independent",
		"Banner",
		"Western Alliance Bancorp",
		"Capital One",
		"Servisfirst Bancshares",
		"Cullen/Frost Bankers",
		"Hancock Holding",
		"Glacier Bancorp",
		"Popular",
		"SVB Financial Group",
		"United Bankshares",
		"PacWest Bancorp",
		"Mechanics",
		"State Street",
		"Community System",
		"First Busey",
		"Pinnacle Financial Partners",
		"Sandy Spring Bancorp",
		"WSFS",
		"United Community",
		"International Bancshares",
		"Central Bancompany",
		"Customers Bancorp",
		"South State",
		"First Merchants",
		"WaFd",
		"Dime Community Bancshares",
		"WesBanco",
		"First BanCorp",
		"TowneBank",
		"Hope Bancorp",
		"Ameris Bancorp",
		"First Republic Bank",
		"Eastern Bankshares	",
		"Cadence Bank",
		"Umpqua Holdings",
		"Axos Financial",
		"First Hawaiian",
		"First Foundation",
		"New York Community Bancorp",
		"Bank of Hawaii",
		"Atlantic Union Bankshares",
		"Webster Financial",
		"Oceanfirst Financial",
		"UMB Financial",
		"Simmons",
		"Heartland Financial",
		"Valley National Bancorp",
		"Texas Capital Bancshares",
		"Provident Financial Svcs",
		"Comerica",
		"Hilltop Holdings	",
		"Commerce Bancshares",
		"Huntington",
		"Synovus",
		"Bank of New York Mellon",
		"First National Of Nebraska",
		"First Financial Bancorp.",
		"First Financial Bancorp.",
		"JPMorgan Chase",
		"Independent Bank Group",
		"Northern Trust",
		"Signature",
		"FB Financial",
		"F.N.B.",
		"TFS Financial",
		"Fulton Bank",
		"Renasant",
		"Northwest Bancshares",
		"Regions Financial",
		"Citigroup",
		"KeyCorp",
		"Old National",
		"M&T",
		"Wintrust Financial",
		"Bank of America",
		"First Citizens Bank (NC)",
		"US Bancorp",
		"Zions Bancorp",
		"Trustmark",
		"First Interstate Bank",
		"First Horizon",
		"Truist Financial",
		"Bankunited",
		"Fifth Third",
		"BOK Financial",
		"Associated Banc-Corp",
		"PNC Financial Services",
		"Citizens Financial Group",
		"Wells Fargo",
	}
	buys := []string{
		"acquire",
		"get",
		"market",
		"obtain",
		"procure",
		"purchase",
		"take",
	}
	fmt.Println(banks, buys)
}
