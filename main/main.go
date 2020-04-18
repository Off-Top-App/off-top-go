package main

import (
	"github.com/off-top-go/webScraping"
	"github.com/off-top-go/webScraping/utils"
)

func main() {
	slice := webScraping.WikiScrape()
	utils.CleanSlice(slice)
	// EXAMPLES BELOW:
	// webScraping.MediumScrape()
	// webScraping.AmazonScrape()
	// webScraping.TwitterScrape()
	// webScraping.StackoverflowScrape()
}
