package rectpuzz

import (
	"testing"
)

func TestStateInit(t *testing.T) {
	s := NewState(&NeighboursNumGrid)
	s.AddRect(Rect{0, 0, 1, 4})

	for y := 0; y < 4; y++ {
		if !s.GPtr().CellPtrAt(0, uint(y)).IsUsed {
			t.Errorf("expected cells associated rect to be 'used'; however,"+
				" cell at (0, %v) is not", y)
			return
		}
	}
}

func TestStateDuplication(t *testing.T) {
	s1 := NewState(&SmallNumGrid)
	s2 := s1.Copy()
	s2.AddRect(Rect{0, 0, 2, 1})

	if len(s1.Rects()) != 0 {
		t.Errorf("expected modification to 's2' to not affect its original, "+
			"'s1', however 's1' is: %v", s1)
	}
}

func neighboursProgressState() State {
	s := NewState(&NeighboursNumGrid)
	rects := [5]Rect{
		{0, 0, 1, 4}, {1, 0, 1, 3},
		{2, 0, 2, 1}, {4, 0, 2, 1},
		{6, 0, 1, 7}}

	for _, r := range rects {
		s.AddRect(r)
	}

	return s
}

func TestStateNeighbours(t *testing.T) {
	s := neighboursProgressState()

	// Neighbour State generation
	neighbours, err := s.Neighbours()
	if err != nil {
		t.Error(err)
	}

	if n := len(neighbours); n != 4 {
		t.Errorf("expected '4' neighbours; instead, found %v", n)
	}

	// Memory management testing
	first := &neighbours[0]
	firstRects := &first.rects
	if ln, cp := len(*firstRects), cap(*firstRects); ln != 6 || cp != 10 {
		t.Errorf("expected first neighbour State's rects to have a " +
			"(len, cap) of (6, 10); instead, got: (%v, %v)", ln, cp)
	}

	first.OptimizeMemoryUsage()
	if ln, cp := len(*firstRects), cap(*firstRects); ln != 6 || cp != 6 {
		t.Errorf("expected first neighbour State's rects to have a " +
			"(len, cap) of (6, 6) after optimization; instead, got: (%v, %v)",
				ln, cp)
	}

}
