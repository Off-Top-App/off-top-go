package main

import (
	"github.com/off-top-go/webScraping"
	"github.com/off-top-go/webScraping/utils"
)

func main() {
	slice := webScraping.WikiScrape()
	outputFilePath := "/Users/disjosh/go/src/github.com/off-top-go/webscrape-texts/"
	utils.CleanSliceAndWriteToFile(outputFilePath, "wiki-sports", slice)
	// EXAMPLES BELOW:
	// webScraping.MediumScrape()
	// webScraping.AmazonScrape()
	// webScraping.TwitterScrape()
	// webScraping.StackoverflowScrape()
}
