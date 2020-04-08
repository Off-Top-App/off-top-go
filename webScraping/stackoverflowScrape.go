package webScraping

import (
	"fmt"

	"github.com/gocolly/colly"
)

func StackoverflowScrape() {

	c := colly.NewCollector(
		colly.AllowedDomains("stackoverflow.com"),
	)
	c.IgnoreRobotsTxt = false
	c.Limit(&colly.LimitRule{
		DomainGlob: "stackoverflow.com",
	})
	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		if e.Text != "" || e.Text != "	\n" {
			fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		}
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://stackoverflow.com/questions/25282572/multiple-assignment-by-if-statement")
}
