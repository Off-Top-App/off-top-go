package main

import (
	"fmt"

	"github.com/off-top-go/webScraping"
	"github.com/off-top-go/webScraping/utils"
)

func main() {
	fmt.Println("Running go service!")
	fileName, slice := webScraping.WikiScrape()
	systemUserName := "/Users/disjosh/"
	outputFilePath := systemUserName + "go/src/github.com/off-top-go/webscrape-texts/"
	utils.CleanSliceAndWriteToFile(outputFilePath, fileName, slice)
// EXAMPLES BELOW:
	// slice := webScraping.WhiskeyLemonadeScrape()
	// webScraping.MediumScrape()
	// webScraping.AmazonScrape()
	// webScraping.TwitterScrape()
	// webScraping.StackoverflowScrape()
}
