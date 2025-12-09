package day04

type Cell struct {
	Idx  uint
	X    uint
	Y    uint
	Type rune
}

type Grid struct {
	Cells  []Cell
	Width  uint
	Height uint
}

func (g *Grid) GetCell(x uint, y uint) Cell {
	return g.Cells[x+y*g.Width]
}

func (g *Grid) Rows() (rows [][]Cell) {
	var currentRow []Cell

	for i, c := range g.Cells {
		if i != 0 && uint(i)%g.Width == 0 {
			rows = append(rows, currentRow)
			currentRow = []Cell{}
		}

		currentRow = append(currentRow, c)
	}
	rows = append(rows, currentRow)

	return rows
}

func (g *Grid) Columns() (cols [][]Cell) {
	for i := uint(0); i < g.Height; i++ {
		cols = append(cols, []Cell{})
	}

	for _, c := range g.Cells {
		cols[c.X] = append(cols[c.X], c)
	}

	return cols
}
