package main

import (
	"github.com/off-top-go/webScraping"
	"github.com/off-top-go/webScraping/utils"
)

func main() {
	slice := webScraping.WikiScrape()
	utils.CleanSliceAndWriteToFile("/Users/disjosh/go/src/github.com/off-top-go/webscrape-texts/", "wiki-sports", slice)
	// EXAMPLES BELOW:
	// webScraping.MediumScrape()
	// webScraping.AmazonScrape()
	// webScraping.TwitterScrape()
	// webScraping.StackoverflowScrape()
}
