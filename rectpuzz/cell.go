package rectpuzz

type Cell struct {
	Num    uint
	IsUsed bool
}

func NewCell(num uint) Cell {
	return Cell{num, false}
}

func ForAllCells(slice *[]Cell, action func(i int, cPtr *Cell)) {
	for i, cell := range *slice {
		action(i, &cell)
	}
}