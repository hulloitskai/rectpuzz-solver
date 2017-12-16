package main

import (
	"bufio"
	"fmt"
	"github.com/steven-xie/rectpuzz-solver/rectpuzz"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	showIntro()
	cols, rows := readGridDimensions()
	grid := generateGrid(cols, rows)
	readGridSetup(&grid)
	displayResults(&grid)
}

func displayResults(gridPtr *[][]uint) {
	clearScreen()
	println("<< RESULTS >>")

	rects, err := rectpuzz.SearchGraph(*gridPtr)
	if err != nil {
		println("Failed to find a possible solution to the puzzle.")
		println("\nInputed grid for reference:")
		showGrid(gridPtr)
		return
	} else {
		println("Found the following solution:")
		drawGridWithRects(gridPtr, rects)
	}
}

func showIntro() {
	clearScreen()
	println("-------------------------")
	println("|| RECTPUZZ SOLVER 1.0 ||")
	println("-------------------------")
}

func readGridDimensions() (cols uint, rows uint) {
	var cols64, rows64 uint64
	var err error
	for {
		colStr := readlnWithPrompt("How many columns does the grid have? ")
		cols64, err = strconv.ParseUint(strings.TrimSpace(colStr), 10, 64)
		if err == nil && cols64 != 0 {
			break
		} else {
			println("That's not a valid number of columns! Try again.")
		}
	}
	for {
		rowStr := readlnWithPrompt("How many rows does this grid have? ")
		rows64, err = strconv.ParseUint(strings.TrimSpace(rowStr), 10, 64)
		if err == nil || rows64 == 0 {
			break
		} else {
			println("That's not a valid number of rows! Try again.")
		}
	}
	return uint(cols64), uint(rows64)
}

func generateGrid(cols uint, rows uint) (grid [][]uint) {
	grid = make([][]uint, rows)
	for i := range grid {
		grid[i] = make([]uint, cols)
	}
	return
}

func clearScreen() {
	var command *exec.Cmd
	if runtime.GOOS == "windows" {
		command = exec.Command("cls")
	} else {
		command = exec.Command("clear")
	}

	command.Stdout = os.Stdout
	command.Run()
}

func showGrid(gridPtr *[][]uint) {
	showGridWithMarker(gridPtr, -1, -1)
}

func showGridWithMarker(gridPtr *[][]uint, currentX int, currentY int) {
	cols := len((*gridPtr)[0])

	printFullRow := func() {
		println("|" + strings.Repeat("----", cols) + " |")
	}
	printEmptyRow := func() {
		println("|" + strings.Repeat("    ", cols) + " |")
	}

	printFullRow()
	for y, row := range *gridPtr {
		rowStr := "|"
		for x, num := range row {
			if x == currentX && y == currentY {
				rowStr += "  _ "
			} else {
				if num == 0 {
					rowStr += "  x "
				} else {
					numStr := strconv.FormatUint(uint64(num), 10)
					if len(numStr) > 1 {
						rowStr += " " + numStr + " "
					} else {
						rowStr += "  " + numStr + " "
					}
				}
			}
		}
		println(rowStr + " |")
		if y != len(*gridPtr)-1 {
			printEmptyRow()
		}
	}
	printFullRow()
}

func hasVerticalBorderLeftOfHere(x, y uint, rectsPtr *[]rectpuzz.Rect) bool {
	for _, rect := range *rectsPtr {
		if x == rect.X {
			if rect.Y <= y && y < rect.Y+rect.H() {
				return true
			}
		}
	}
	return false
}

func hasHorizontalBorderAboveHere(x, y uint, rectsPtr *[]rectpuzz.Rect) bool {
	for _, rect := range *rectsPtr {
		if y == rect.Y {
			if rect.X <= x && x < rect.X+rect.W() {
				return true
			}
		}
	}
	return false
}

func drawGridWithRects(gridPtr *[][]uint, rects []rectpuzz.Rect) {
	cols := len((*gridPtr)[0])
	println("|" + strings.Repeat("----", cols) + "-|")

	for y, row := range *gridPtr {
		rowStr, borderStr := "| ", "|"

		for x, num := range row {
			if x > 0 {
				if hasVerticalBorderLeftOfHere(uint(x), uint(y), &rects) {
					rowStr += "|"
					if hasHorizontalBorderAboveHere(uint(x), uint(y), &rects) {
						if x == cols - 1 {
							if borderStr[len(borderStr)-1] == '-' {
								borderStr += "-----|"
							} else {
								borderStr += "|----|"
							}
						} else {
							if borderStr[len(borderStr)-1] == '-' {
								borderStr += "----"
							} else {
								borderStr += "|---"
							}
						}
					} else {
						if x == cols - 1 {
							borderStr += "|    |"
						} else {
							borderStr += "|   "
						}
					}
				} else {
					rowStr += " "
					if hasHorizontalBorderAboveHere(uint(x), uint(y), &rects) {
						if x == cols-1 {
							borderStr += "-----|"
						} else {
							borderStr += "----"
						}
					} else {
						if x == cols-1 {
							borderStr += "     |"
						} else {
							borderStr += "    "
						}
					}
				}
			} else {
				if hasHorizontalBorderAboveHere(uint(x), uint(y), &rects) {
					borderStr += "----"
				} else {
					borderStr += "    "
				}
			}

			if num == 0 {
				rowStr += " x "
			} else {
				numStr := strconv.FormatUint(uint64(num), 10)
				if len(numStr) > 1 {
					rowStr += numStr + " "
				} else {
					rowStr += " " + numStr + " "
				}
			}
		}

		if y > 0 {
			println(borderStr)
		}
		println(rowStr + " |")
	}
	println("|" + strings.Repeat("----", cols) + "-|")
}

func readGridSetup(gridPtr *[][]uint) {
	for y := 0; y < len(*gridPtr); y++ {
		for x := 0; x < len((*gridPtr)[y]); x++ {
			clearScreen()
			println("<< GRID SETUP >>")
			showGridWithMarker(gridPtr, x, y)

			println("\nEnter a number to put in the grid at the specified " +
				"position.")
			println("If there should be no number in the specified position, " +
				"just press enter.\n")
			println("If you make a mistake, enter 'b' or 'back' to go back.\n")

			inputNum, goBack := readGridNumWithPrompt("Number: ")
			if goBack {
				if x == 0 {
					if y != 0 {
						y -= 1
						x = len((*gridPtr)[0]) - 2
					} else {
						x -= 1
					}
				} else {
					x -= 2
				}
			} else {
				(*gridPtr)[y][x] = inputNum
			}
		}
	}
}

func readGridNumWithPrompt(fmtPrompt string) (uint, bool) {
	for {
		input := strings.TrimSpace(readlnWithPrompt(fmtPrompt))
		if input == "" {
			return 0, false
		} else if lStr := strings.ToLower(input); lStr == "b" || lStr == "back" {
			return 0, true
		}

		input64, err := strconv.ParseUint(input, 10, 64)
		if err == nil || input64 > 99 {
			return uint(input64), false
		}
		println("That's not a valid natural number beween 0 and 99! Try again.")
	}
}

func readlnWithPrompt(fmtPrompt string) string {
	fmt.Printf(fmtPrompt)
	return readln()
}

func readln() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}
