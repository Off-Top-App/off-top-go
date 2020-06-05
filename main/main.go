package main

import (
	"fmt"

	"github.com/off-top-go/webScraping"
	"github.com/off-top-go/webScraping/utils"
)

func main() {
	fmt.Println("Running go service!")
	slice := webScraping.WikiScrape()
	outputFilePath := "/Users/disjosh/go/src/github.com/off-top-go/webscrape-texts/"
	utils.CleanSliceAndWriteToFile(outputFilePath, "wiki-sports", slice)
	// EXAMPLES BELOW:
	// webScraping.MediumScrape()
	// webScraping.AmazonScrape()
	// webScraping.TwitterScrape()
	// webScraping.StackoverflowScrape()
}
