package main

import (
	"fmt"
	"testing"
)

func TestSandbox(t *testing.T) {
	rowCount := 5
	colCount := 5

	numGrid := make([][]uint, rowCount)
	strs := make([]string, rowCount)

	for i := range numGrid {
		numGrid[i] = make([]uint, colCount)

		runes := make([]rune, 3*colCount)
		for k := 0; k < 3*colCount; k += 3 {
			runes[k], runes[k+1], runes[k+2] = ' ', 'x', ' '
		}
		strs[i] = string(runes)
	}

	for _, str := range strs {
		fmt.Println(str)
	}
}


