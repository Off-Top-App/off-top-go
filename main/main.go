package main

import (
	"github.com/off-top-go/webScraping"
	"github.com/off-top-go/webScraping/utils"
)

func main() {
	slice := webScraping.NBAWikiScrape()
	utils.CleanSliceAndWriteToFile("C:/go-work/src/github.com/off-top-go/webscrape-texts", "hel-lo", slice)
	// EXAMPLES BELOW:
	// webScraping.MediumScrape()
	// webScraping.AmazonScrape()
	// webScraping.TwitterScrape()
	// webScraping.StackoverflowScrape()
}
