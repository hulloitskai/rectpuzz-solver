package rectpuzz

import (
	"fmt"
	"testing"
)

var (
	cS = []Cell{{3, true}, {0, true},
		{1, true}}
)

func TestSliceAndMapping(t *testing.T) {
	result := true
	ForAllCells(&cS, func(_ int, c *Cell) {
		result = result && c.IsUsed
	})

	if !result {
		t.Errorf("expected all rectgame of 'cS' to be used; "+
			"however cS is: %v", cS)
	}
}

func TestForLoops(t *testing.T) {
	var done bool

	for k := 0; k < 5; k++ {
		for i := 0; i < 5; i++ {
			fmt.Printf("(%v, %v)\n", i, k)
			if k == 3 && i == 0 && done == false{
				k -= 1
				done = true
			}
		}
	}
}
