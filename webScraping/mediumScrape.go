package webScraping

import (
	"fmt"

	"github.com/gocolly/colly"
)

func MediumScrape() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiiting", r.URL)
	})

	c.OnHTML("div.z.ab.ac.ae.af.do.ah.ai", func(e *colly.HTMLElement) {
		fmt.Print(e)
	})

	c.Visit("https://medium.com/analytics-vidhya/learn-how-to-code-and-deploy-machine-learning-models-on-structured-streaming-868b4081d242")
}