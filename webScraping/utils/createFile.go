package utils

import (
	"fmt"
	"os"
	"strings"
)

func WriteToFile(myDirectory string, fileName string, data []string) {
	fileDirectory := myDirectory
	exportedFileName := fmt.Sprintf("%s/%s.txt", fileDirectory, fileName)
	file, fileErr := os.Create(exportedFileName)
	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}
	stringData := strings.Join(data, " ")

	fmt.Fprintf(file, "%s", stringData)
}
