package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ErrorDataNotPopulated is thrown when an operation requires a member of
// FileData to be populated and that member is not populated.
type ErrorDataNotPopulated struct {
	Item string
}

func (err ErrorDataNotPopulated) Error() string {
	return fmt.Sprintf("The operation requires the following to be populated: %v", err.Item)
}

// FileData stores statistical information about a file.
type FileData struct {
	Lines      []string
	WordCounts map[string]int
	FileName   string
}

func (fd *FileData) getLines() error {
	lines := make([]string, 0)
	f, err := os.Open("file.txt")
	checkError(err)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fd.Lines = lines
	return scanner.Err()
}

func (fd *FileData) getWords() error {
	result := make(map[string]int)
	lines := fd.Lines
	if lines == nil {
		return &ErrorDataNotPopulated{"lines"}
	}
	for _, line := range lines {
		for _, word := range strings.Split(line, " ") {
			if _, ok := result[word]; ok {
				result[word] = result[word] + 1
			} else {
				result[word] = 1
			}
		}
	}
	fd.WordCounts = result
	return nil
}

// GetMostFrequentSmallWord returns the word that is at under 5 characters that is
// present the most in the file.
func (fd *FileData) GetMostFrequentSmallWord() (string, int, error) {
	var smallWord string
	var smallWordCount int

	if len(fd.WordCounts) == 0 {
		return "", 0, &ErrorDataNotPopulated{"WordCount"}
	}

	for word, count := range fd.WordCounts {
		if len(word) > 0 && len(word) <= 4 {
			if count > smallWordCount {
				smallWord = word
				smallWordCount = count
			}
		}
	}

	return smallWord, smallWordCount, nil
}

// GetMostFrequentMediumWord returns the word tha tis between 5 and 7 characters long
// and is present the most in the file.
func (fd *FileData) GetMostFrequentMediumWord() (string, int, error) {
	var medWord string
	var medWordCount int

	if len(fd.WordCounts) == 0 {
		return "", 0, &ErrorDataNotPopulated{"WOrdCount"}
	}

	for word, count := range fd.WordCounts {
		if len(word) >= 5 && len(word) <= 7 {
			if count > medWordCount {
				medWord = word
				medWordCount = count
			}
		}
	}

	return medWord, medWordCount, nil
}

// GetMostFrequentBigWord returns the word tha tis greater than 6 characters long
// and is present the most in the file.
func (fd *FileData) GetMostFrequentBigWord() (string, int, error) {
	var bigWord string
	var bigWordCount int

	if len(fd.WordCounts) == 0 {
		return "", 0, &ErrorDataNotPopulated{"WOrdCount"}
	}

	for word, count := range fd.WordCounts {
		if len(word) > 7 {
			if count > bigWordCount {
				bigWord = word
				bigWordCount = count
			}
		}
	}

	return bigWord, bigWordCount, nil
}

// Populate populates the members of FileData object instance.
func (fd *FileData) Populate(fileName string) (err error) {
	if len(fileName) == 0 {
		return fmt.Errorf("fileName cannot be empty")
	}
	fd.FileName = fileName
	if err = fd.getLines(); err != nil {
		return
	}

	if err = fd.getWords(); err != nil {
		return
	}
	return nil
}
