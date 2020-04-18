package webScraping

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func WikiScrape() []string {
	slice := make([]string, 3000)
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("div.mw-body-content", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, e *colly.HTMLElement) {
			text := e.Text
	
			if !strings.Contains(text, ".") {
				if strings.Contains(text, "CS1 maint") == false {
					if strings.Contains(text, "Jump to") == false {
						text = strings.TrimSpace(text)
						slice = append(slice, text)
					}
				}
			}
		})
	})

	c.Visit("https://en.wikipedia.org/wiki/List_of_sports")
	return slice
}
