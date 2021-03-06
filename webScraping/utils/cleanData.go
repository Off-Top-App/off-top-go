package utils

import "reflect"

func ItemExists(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("Invalid data-type")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func CleanSliceAndWriteToFile(myDirectory string, fileName string, words []string) {
	newSlice := cleanData(words)
	WriteToFile(myDirectory, fileName, newSlice)

}

func cleanData(words []string) []string {
	newSlice := make([]string, 5000)
	for _, word := range words {
		doesWordExist := ItemExists(newSlice, word)
		if doesWordExist == false {
			newSlice = append(newSlice, word)
		}
	}
	return newSlice
}
