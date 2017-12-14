package rectpuzz

import (
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
