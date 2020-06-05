package webScraping

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func WhiskeyLemonadeScrape() []string {
	slice := make([]string, 3000)
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("div.post.singlepost", func(e *colly.HTMLElement) {
		e.ForEach("p", func(_ int, e *colly.HTMLElement) {
			text := e.Text

			text = strings.TrimSpace(text)
			slice = append(slice, text)

		})
		fmt.Println(slice)
	})

	c.Visit("https://www.gimmesomeoven.com/whiskey-lemonade/")
	return slice
}
