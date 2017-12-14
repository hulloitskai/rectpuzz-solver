package rectpuzz

import "fmt"

type State struct {
	g Grid
	rects []Rect
}

func (s State) String() string {
	return fmt.Sprintf("State(grid: %v, rects: %v)", s.g, s.rects)
}

func (sPtr *State) G() Grid {
	return sPtr.g
}

func (sPtr *State) Rects() []Rect {
	return sPtr.rects
}

func (sPtr *State) AddRect(r Rect) {
	sPtr.rects = append(sPtr.rects, r)
	sPtr.g.AddRect(&r)
}

func NewState(numGridPtr *[][]uint) State {
	return State{NewGrid(numGridPtr), make([]Rect, 0)}
}

func (sPtr *State) IsSolved() bool {
	return sPtr.g.IsCompletelyUsed()
}

func (sPtr *State) GPtr() *Grid {
	return &sPtr.g
}

func (sPtr *State) Copy() State {
	return State{sPtr.GPtr().Copy(),append([]Rect(nil), sPtr.Rects()...)}
}

func (sPtr *State) Neighbours() ([]State, error) {
	rects, err := sPtr.GPtr().GenerateRects()
	if err != nil {
		return nil, err
	}

	states := make([]State, len(rects))
	for i, r := range rects {
		clone := sPtr.Copy()
		clone.AddRect(r)
		states[i] = clone
	}

	return states, nil
}

func StateEquals(sPtr1, sPtr2 *State) bool {
	for i, r := range sPtr1.rects {
		if r != sPtr2.rects[i] {
			return false
		}
	}
	return GridEquals(sPtr1.GPtr(), sPtr2.GPtr())
}

// Reallocates the underlying arrays for a State's rects, resulting in improved
// memory usage at the expense of processing time.
func (sPtr *State) OptimizeMemoryUsage() {
	sPtr.rects = append([]Rect(nil), sPtr.rects...)
}

/*
func (sPtr *State) optimizeMemoryAlternate() {
	rects := make([]Rect, len(sPtr.rects))
	copy(rects, sPtr.rects)
	sPtr.rects = rects
}
*/

