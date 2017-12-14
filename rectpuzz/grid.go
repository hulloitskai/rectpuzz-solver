package rectpuzz

import (
	"errors"
)

type Grid [][]Cell

func NewGrid(numGridPtr *[][]uint) Grid {
	cellGrid := make([][]Cell, len(*numGridPtr))
	for y, numRow := range *numGridPtr {
		cellRow := make([]Cell, len(numRow))
		for x, num := range numRow {
			cellRow[x] = NewCell(num)
		}
		cellGrid[y] = cellRow
	}
	return cellGrid
}

func (gPtr *Grid) ForAll(action func(x int, y int, cPtr *Cell)) {
	for y, row := range *gPtr {
		for x, c := range row {
			action(x, y, &c)
			(*gPtr)[y][x] = c
		}
	}
}

func (gPtr *Grid) IsCompletelyUsed() bool {
	for _, row := range *gPtr {
		for _, cell := range row {
			if !cell.IsUsed {
				return false
			}
		}
	}
	return true
}

func (gPtr *Grid) GetFirstUnusedCoord() (x int, y int, success bool) {
	for y, row := range *gPtr {
		for x, c := range row {
			if !c.IsUsed {
				return x, y, true
			}
		}
	}
	return x, y, false
}

func (gPtr *Grid) generateBoundedRects() ([]Rect, error) {
	rects := make([]Rect, 0)

	xBase, yBase, success := gPtr.GetFirstUnusedCoord()
	if (!success) {
		return nil, errors.New("unable to locate coordinates of first unused" +
			" cell; could not generate bounded Rects")
	}

	var maxW int
	for h := 1; h <= len(*gPtr)-yBase; h++ {
		for w := 1; w <= len((*gPtr)[0])-xBase; w++ {
			if maxW != 0 && w == maxW {
				break
			} else if (*gPtr)[yBase+h-1][xBase+w-1].IsUsed {
				maxW = w
				break
			}

			r, err := NewRect(uint(xBase), uint(yBase), uint(w), uint(h))
			if err != nil {
				return nil, err
			}
			rects = append(rects, r)
		}
	}

	return rects, nil
}

func (gPtr *Grid) GenerateRects() ([]Rect, error) {
	rects, err := gPtr.generateBoundedRects()
	if err != nil {
		return rects, err
	}

	filteredRects := rects[:0]
	for _, rect := range rects {
		if rectIsValidInGrid(&rect, gPtr) {
			filteredRects = append(filteredRects, rect)
		}
	}

	return filteredRects, nil
}

func rectIsValidInGrid(rPtr *Rect, gPtr *Grid) bool {
	var num uint

	for y := rPtr.Y; y < rPtr.Y+rPtr.h; y++ {
		for x := rPtr.X; x < rPtr.X+rPtr.w; x++ {
			cell := &(*gPtr)[y][x]
			if cell.IsUsed {
				return false
			} else if n := cell.Num; n != 0 {
				if num != 0 {
					return false
				}
				num = n
			}
		}
	}

	return num == rPtr.Area()
}

func (gPtr *Grid) AddRect(rPtr *Rect) {
	maxX, maxY := rPtr.X+rPtr.w, rPtr.Y+rPtr.h
	for y := rPtr.Y; y < maxY; y++ {
		for x := rPtr.X; x < maxX; x++ {
			(*gPtr)[y][x].IsUsed = true
		}
	}
}

func (gPtr *Grid) ToUintGrid() [][]uint {
	uintGrid := make([][]uint, len(*gPtr))
	for i, cellRow := range *gPtr {
		uintRow := make([]uint, len(cellRow))
		for k, cell := range cellRow {
			uintRow[k] = cell.Num
		}
		uintGrid[i] = uintRow
	}
	return uintGrid
}

func (gPtr *Grid) CellPtrAt(x uint, y uint) *Cell {
	return &((*gPtr)[y][x])
}

func (gPtr *Grid) CellAt(x uint, y uint) Cell {
	return (*gPtr)[y][x]
}

func (gPtr *Grid) Copy() Grid {
	newG := make(Grid, len(*gPtr))

	for y, row := range *gPtr {
		newG[y] = append([]Cell(nil), row...)
	}

	return newG
}

func GridEquals(gPtr1, gPtr2 *Grid) bool {
	for y, row := range *gPtr1 {
		for x, cell := range row {
			if cell != *gPtr1.CellPtrAt(uint(x), uint(y)) {
				return false
			}
		}
	}
	return true
}