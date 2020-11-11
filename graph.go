package main

import "fmt"

// Graph contains data that represents a bar graph which can be printed
// to standard out
type Graph struct {
	bitmap      map[[2]int]byte
	GraphHeight int
	GraphWidth  int
	MarginWidth int
	graphCount  int
}

// CalculateCatHeight calculates the height of each category within a bar
// on the bar graph.
func (g *Graph) CalculateCatHeight(totalWords int, count int) int {
	return int((g.GraphHeight / totalWords) * count)
}

// AddCharData takes a string and adds it to the bitmap at the specified
// x, y location.
func (g *Graph) AddCharData(text string, x int, y int) {
	textLen := len(text)
	var width int

	// Ensure bitmap is not nil
	if g.bitmap == nil {
		g.bitmap = make(map[[2]int]byte)
	}

	if g.GraphWidth >= textLen {
		width = g.GraphWidth
	} else {
		width = textLen
	}
	for i := 0; i < width+1; i++ {
		var coordinate [2]int = [2]int{i + 1 + x, y}
		var val byte
		if i < textLen {
			val = text[i]
		} else {
			val = ' '
		}
		g.bitmap[coordinate] = val
	}
}

// AddWordLengthGraphData writes data to the bitmap to build a bar graph of the Word Length Graph
func (g *Graph) AddWordLengthGraphData(smallWords int, medWords int, bigWords int) {
	// Ensure bitmap is not nil
	if g.bitmap == nil {
		g.bitmap = make(map[[2]int]byte)
	}

	g.graphCount++
	totalWords := smallWords + medWords + bigWords
	smallHeight := g.CalculateCatHeight(totalWords, smallWords)
	medHeight := g.CalculateCatHeight(totalWords, medWords)
	medCatLoc := smallHeight
	bigCatLoc := medCatLoc + medHeight

	for x := 0; x < g.GraphWidth+1; x++ {
		for y := 0; y < g.GraphHeight; y++ {
			coord := [2]int{x, y}
			if x == 0 || x == g.GraphWidth {
				g.bitmap[coord] = '|'
			} else if y == 0 || y == medCatLoc || y == bigCatLoc ||
				y == g.GraphHeight {
				g.bitmap[coord] = '-'
			} else if y == 1 {
				g.AddCharData("Small words", 0, 1)
			} else if y == medCatLoc+1 {
				g.AddCharData("Medium words", 0, medCatLoc+1)
			} else if y == bigCatLoc+1 {
				g.AddCharData("Big words", 0, bigCatLoc+1)
			} else {
				g.bitmap[coord] = ' '
			}
		}
	}
}

// DrawBitmap draws a text representation of bitmap.
func (g *Graph) DrawBitmap() {
	for y := 0; y < g.GraphHeight; y++ {
		for x := 0; x < g.GraphWidth*g.graphCount+
			g.MarginWidth*g.graphCount; x++ {
			coord := [2]int{x, y}
			if val, ok := g.bitmap[coord]; ok {
				fmt.Print(string(val))
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
