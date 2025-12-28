package main

import (
	"fmt"
	"github.com/Shuddown/Advent-of-Code/utils"
	"io"
	"strings"
)

const (
	SPLITTER = '^'
)

type splitter struct {
	row  int
	used bool
}

func rowCol(index, numCols int) (int, int) {
	return index / (numCols + 1), index % (numCols + 1)
}

var totalCount = 0

var totalIndex [][]int
var splitters [][]*splitter

func printNumSplitsPartOne(row, col int) {
	for _, s := range splitters[col] {
		if row < s.row {
			if s.used {
				return
			}
			totalCount += 1
			s.used = true
			printNumSplitsPartOne(s.row, col+1)
			printNumSplitsPartOne(s.row, col-1)
			return
		}
	}
}

func printNumSplitsPartTwo(row, col int) int {
	for _, s := range splitters[col] {
		if row < s.row {
			if totalIndex[s.row][col] != 0 {
				return totalIndex[s.row][col]
			}
			totalIndex[s.row][col] += printNumSplitsPartTwo(s.row, col+1)
			totalIndex[s.row][col] += printNumSplitsPartTwo(s.row, col-1)
			return totalIndex[s.row][col]
		}
	}
	return 1
}

func main() {
	inputFile, err := utils.GetInput(7)
	utils.HandleError(err)
	bytes, err := io.ReadAll(inputFile)
	utils.HandleError(err)
	manifoldDiagram := string(bytes)
	width := strings.Index(manifoldDiagram, "\n")
	height := (len(manifoldDiagram) + 1) / (width + 1)
	totalIndex = make([][]int, width)
	for i := range totalIndex {
		totalIndex[i] = make([]int, height)
	}
	splitters = make([][]*splitter, width)
	for i, r := range manifoldDiagram {
		if r == SPLITTER {
			row, col := rowCol(i, width)
			splitters[col] = append(splitters[col], &splitter{row, false})
		}
	}
	sourceCol := width / 2
	printNumSplitsPartOne(0, sourceCol)
	fmt.Println(totalCount)
	count := printNumSplitsPartTwo(0, sourceCol)
	fmt.Println(count)
}
