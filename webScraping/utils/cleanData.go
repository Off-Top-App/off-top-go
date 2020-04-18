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

func CleanSlice(words []string) {
	newSlice := make([]string, 3000)
	for _, word := range words {
		doesWordExist := ItemExists(newSlice, word)
		if doesWordExist == false {
			newSlice = append(newSlice, word)
		}
	}
	WriteToFile("wiki-sports", newSlice)

}
