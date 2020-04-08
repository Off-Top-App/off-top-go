package webScraping

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/off-top-go/webScraping/utils"
)

func AmazonScrape() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.amazon.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("div.s-result-list.s-search-results.sg-row", func(e *colly.HTMLElement) {
		e.ForEach("div.a-section.a-spacing-medium", func(_ int, e *colly.HTMLElement) {

			productName := e.ChildText("span.a-size-medium.a-color-base.a-text-normal")
			if productName == "" {
				// If we can't get any name, we return and go directly to the next element
				return
			}
			stars := e.ChildText("span.a-icon-alt")
			utils.FormatStars(&stars)

			price := e.ChildText("span.a-price > span.a-offscreen")
			utils.FormatPrice(&price)

			fmt.Printf("Product Name: %s \nStars: %s \nPrice: %s \n", productName, stars, price)
		})
	})

	c.Visit("https://www.amazon.com/s?k=pokemon+switch+games&ref=nb_sb_noss_1")
}
