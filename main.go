package main

import "fmt"

// checkError checks for an error and panics (stops execution)
// if one is found.
func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fd := &FileData{}
	fd.Populate("file.txt")
	fmt.Println(fd.WordCounts)

	word, count, err := fd.GetMostFrequentSmallWord()
	checkError(err)
	fmt.Printf("Small word %v, count %d\n", word, count)

	word, count, err = fd.GetMostFrequentMediumWord()
	checkError(err)
	fmt.Printf("Medium word %v, count %d\n", word, count)

	word, count, err = fd.GetMostFrequentBigWord()
	checkError(err)
	fmt.Printf("Big word %v, count %d\n", word, count)
}
