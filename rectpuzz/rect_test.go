package rectpuzz

import (
	"fmt"
	"testing"
)

var (
	r1, err1 = NewRect(0, 0, 3, 2)
	r2, err2 = NewRect(1, 2, 1, 5)
)

func ExampleInit() {
	r, err := NewRect(2, 0, 3, 2)
	fmt.Printf("r1: %v, err: %v", r, err)

	// Output:
	// r1: Rect(x: 2, y: 0, w: 3, h: 2), err: <nil>
}

func TestErroneousInit(t *testing.T) {
	r, err := NewRect(0, 3, 0, 5)

	if zeroR := (Rect{}); r != zeroR {
		t.Errorf("expected %v, but received %v",
			r, zeroR)
	}

	if err == nil || err.Error() != "UnexpectedZeroError: 'w' must be non-zero" {
		t.Errorf("expected 'UnexpectedZeroError', but received %v", err)
	}
}
