package utils

import (
	"fmt"
	"os"
	"strings"
)

func WriteToFile(fileName string, data []string) {
	fileDirectory := "/Users/disjosh/go/src/github.com/off-top-go/webscrape-texts/"
	exportedFileName := fmt.Sprintf("%s/%s.txt", fileDirectory, fileName)
	file, fileErr := os.Create(exportedFileName)
	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}
	stringData := strings.Join(data, " ")

	fmt.Fprintf(file, "%s", stringData)
}
