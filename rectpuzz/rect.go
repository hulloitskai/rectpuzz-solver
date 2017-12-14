package rectpuzz

import "fmt"

type Rect struct {
	X, Y, w, h uint
}

// Getter for W
func (rPtr *Rect) W() uint {
	return rPtr.w
}

// Setter for W
func (rPtr *Rect) SetW(width uint) {
	rPtr.w = width
}

// Getter for H
func (rPtr *Rect) H() uint {
	return rPtr.h
}

// Setter for H
func (rPtr *Rect) SetH(height uint) {
	rPtr.h = height
}

// Safely creates a new Rect
func NewRect(x, y, w, h uint) (Rect, error){
	if w == 0 {
		return Rect{}, UnexpectedZeroError{"w"}
	} else if h == 0 {
		return Rect{}, UnexpectedZeroError{"h"}
	}

	return Rect{x, y, w, h}, nil
}

func (r Rect) String() string {
	return fmt.Sprintf("Rect(x: %v, y: %v, w: %v, h: %v)", r.X, r.Y, r.w, r.h)
}

func FilterRectSlice(rSlicePtr *[]Rect, pred func(rPtr *Rect) bool) {
	for _, rect := range *rSlicePtr {
		if !pred(&rect) {

		}
	}
}

func (rPtr *Rect) Area() uint {
	return rPtr.w * rPtr.h
}
