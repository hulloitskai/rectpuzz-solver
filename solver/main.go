package main

import (
	"bufio"
	"fmt"
	tm "github.com/buger/goterm"
	"os"
	"strconv"
)

func main() {
	tm.Clear()
	for {
		introSeq()
		colCountStr := readlnPrompt("Number of columns? ")
		rowCountStr := readlnPrompt("\nNumber of rows? ")

		colCount, _ := strconv.Atoi(colCountStr)
		rowCount, _ := strconv.Atoi(rowCountStr)
		setupNumGrid(colCount, rowCount)
		colCountStr = readln()
	}
}

func setupNumGrid(colCount int, rowCount int) [][]uint {
	numGrid := make([][]uint, rowCount)

	strs := make([]string, rowCount)
	strLength := 3 * colCount

	for i := range numGrid {
		numGrid[i] = make([]uint, colCount)

		runes := make([]rune, strLength)
		for i := 0; i < strLength; i += 3 {
			runes[i], runes[i+1], runes[i+2] = ' ', 'x', ' '
		}
		strs[i] = string(runes)
	}

	printBoxSeq(strs, 50, 15)
	return numGrid
}

func readlnPrompt(fmtPrompt string) string {
	fmt.Printf(fmtPrompt)
	return readln()
}

func readln() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}

func introSeq() {
	strs := []string{"", "          RECT-PUZZ SOLVER v0.1"}
	printBoxSeq(strs, 45, 5)
}

func printBoxSeq(fmtStrs []string, width int, height int) {
	// Reset for new sequence
	tm.Clear()

	// Create new box, fill it with content
	box := tm.NewBox(width, height, 0)
	for _, fmtStr := range fmtStrs {
		fmt.Fprintln(box, fmtStr)
	}

	// Move box into place, and print!
	tm.Print(tm.MoveTo(box.String(), 0, 1))
	tm.Flush()
}
