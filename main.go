package main

import "fmt"

// checkError checks for an error and panics (stops execution)
// if one is found.
func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func printHeader(fd FileData) {
	smallWord, smallWordCount, err := fd.GetMostFrequentSmallWord()
	checkError(err)
	medWord, medWordCount, err := fd.GetMostFrequentMediumWord()
	checkError(err)
	bigWord, bigWordCount, err := fd.GetMostFrequentBigWord()
	checkError(err)
	fmt.Println(fd.FileName)
	fmt.Printf("Total Unique Words: %d\n", len(fd.WordCounts))
	fmt.Printf(
		"Most used words (s/m/b): "+
			"%v (%dx) %v (%dx) %v (%dx)\n",
		smallWord,
		smallWordCount,
		medWord,
		medWordCount,
		bigWord,
		bigWordCount)
}

func main() {
	fd := &FileData{}
	fd.Populate("file.txt")
	fmt.Println(fd.WordCounts)
	printHeader(*fd)
	g := &Graph{
		GraphHeight: 20,
		GraphWidth:  10,
		MarginWidth: 10,
	}
	g.AddWordLengthGraphData(7, 3, 1)
	g.DrawBitmap()
}
