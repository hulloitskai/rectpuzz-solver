package rectpuzz

import (
	"reflect"
	"testing"
)

var (
	PuzzNumGrid = [][]uint{
		{0, 0, 0, 0, 0, 5, 0},
		{0, 0, 0, 0, 0, 2, 2},
		{0, 3, 0, 6, 3, 2, 0},
		{4, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 4, 0, 4, 0},
		{2, 0, 6, 0, 2, 4, 0},
		{0, 0, 0, 0, 0, 0, 0}}

	SmallNumGrid = [][]uint{
		{0, 0, 3},
		{2, 0, 0},
		{0, 4, 0}}

	NeighboursNumGrid = [][]uint{
		{0, 0, 0, 2, 0, 2, 0},
		{0, 0, 0, 0, 0, 4, 0},
		{0, 3, 6, 0, 0, 0, 0},
		{4, 2, 0, 0, 0, 4, 0},
		{0, 0, 0, 2, 0, 0, 0},
		{0, 0, 0, 0, 10, 0, 7},
		{3, 0, 0, 0, 0, 0, 0}}
)

func TestGridCompletion(t *testing.T) {
	g := NewGrid(&PuzzNumGrid)
	g.ForAll(func(_ int, _ int, cPtr *Cell) {
		cPtr.IsUsed = true
	})

	if !g.IsCompletelyUsed() {
		t.Errorf("expected 'g' to be completely numbered. G is: %v",
			g)
	}
}

func TestConversionsAndPointers(t *testing.T) {
	g := NewGrid(&NeighboursNumGrid)
	gNumGrid := g.ToUintGrid()

	if !reflect.DeepEqual(gNumGrid, NeighboursNumGrid) {
		t.Errorf("expected gNumGrid to be equal to NeighboursNumGrid; "+
			"however, gNumGrid is: %v", gNumGrid)
	}

	g.CellPtrAt(3, 0).Num = 5
	gNumGrid = g.ToUintGrid()
	if gNumGrid[0][3] == NeighboursNumGrid[0][3] {
		t.Errorf("expected gNumGrid at [0][3] to be different from "+
			"NeighboursNumGrid at [0][3]; however, gNumGrid at [0][3] is: %v",
			gNumGrid[0][3])
	}
}

func TestRectAdd(t *testing.T) {
	g := NewGrid(&NeighboursNumGrid)

	r, err := NewRect(0, 0, 4, 2)
	if err != nil {
		t.Error(err)
	}

	g.AddRect(&r)
	for x := 0; x < 4; x++ {
		if g.CellPtrAt(uint(x), 0).IsUsed == false {
			t.Errorf("expected first four cells to be used: however, Cell "+
				"at (%v, 0) is unused", x)
		}
	}
}

func TestBoundedRectGeneration(t *testing.T) {
	g := NewGrid(&SmallNumGrid)
	g.CellPtrAt(2, 0).IsUsed, g.CellPtrAt(1, 1).IsUsed = true, true

	if x, y, _ := g.GetFirstUnusedCoord(); x != 0 && y != 0 {
		t.Errorf("received unexpected 'first unused' coordinates: (%v, %v)",
			x, y)
	}

	genRects, err := g.generateBoundedRects()
	if err != nil {
		t.Error(err)
	}

	expectedRects := []Rect{
		{0, 0, 1, 1},
		{0, 0, 2, 1},
		{0, 0, 1, 2},
		{0, 0, 1, 3}}
	if !reflect.DeepEqual(expectedRects, genRects) {
		t.Errorf("received unexpected Rects: %v", genRects)
	}
}

func neighboursProgressGrid() Grid {
	g := NewGrid(&NeighboursNumGrid)
	rects := [5]Rect{
		{0, 0, 1, 4}, {1, 0, 1, 3},
		{2, 0, 2, 1}, {4, 0, 2, 1},
		{6, 0, 1, 7}}

	for _, rect := range rects {
		g.AddRect(&rect)
	}
	return g
}

func TestRectValidityChecker(t *testing.T) {
	g := neighboursProgressGrid()
	r, _ := NewRect(2, 1, 3, 2)

	if !rectIsValidInGrid(&r, &g) {
		t.Errorf("expected %v to be valid in grid; but it was not", r)
	}

	r, _ = NewRect(5, 1, 2, 2)
	if rectIsValidInGrid(&r, &g) {
		goto FAILED
	}

	r, _ = NewRect(3, 2, 3, 2)
	if rectIsValidInGrid(&r, &g) {
		goto FAILED
	}

	r, _ = NewRect(2, 2, 2, 3)
	if rectIsValidInGrid(&r, &g) {
		goto FAILED
	}

	return

FAILED:
	t.Errorf("expected %v to be invalid in grid; but it was", r)
}

func TestValidRectGeneration(t *testing.T) {
	g := neighboursProgressGrid()

	generatedRects, err := g.GenerateRects()
	if err != nil {
		t.Error(err)
	}

	expectedRects := []Rect{
		{2, 1, 4, 1}, {2, 1, 3, 2},
		{2, 1, 2, 3}, {2, 1, 1, 6}}

	if !reflect.DeepEqual(generatedRects, expectedRects) {
		t.Errorf("generated Rects, %v, do not match expected Rects",
			generatedRects)
	}
}

func TestGridDuplication(t *testing.T) {
	g1 := NewGrid(&SmallNumGrid)
	g2 := g1.Copy()

	g2.CellPtrAt(0, 0).Num = 500
	if num := g1.CellPtrAt(0, 0).Num; num == 500 {
		t.Errorf("expected cell at (0, 0) for 'g1' (the original Grid) to " +
			"remain unchanged, but found a cell with a num of: %v", num)
	}
}
