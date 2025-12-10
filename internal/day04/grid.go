package day04

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Cell struct {
	Idx  uint
	X    uint
	Y    uint
	Type rune
}

type Grid struct {
	Cells  *[]Cell
	Width  uint
	Height uint
}

func (g *Grid) GetCell(x uint, y uint) Cell {
	return (*g.Cells)[x+y*g.Width]
}

func (g *Grid) SetCell(x uint, y uint, c Cell) {
	(*g.Cells)[x+y*g.Width] = c
}

func (g *Grid) GetNeighbors(x uint, y uint, wrap bool, diagonal bool) (neighbors []Cell) {
	iWidth := int(g.Width)
	iHeight := int(g.Height)
	iX := int(x)
	iY := int(y)

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			absSum := abs(i) + abs(j)
			if absSum == 0 || !diagonal && absSum == 2 {
				continue
			}

			nx := iX + i
			if nx < 0 || nx >= iWidth {
				if wrap {
					nx = (nx + iWidth) % iWidth
				} else {
					continue
				}
			}

			ny := iY + j
			if ny < 0 || ny >= iHeight {
				if wrap {
					ny = (ny + iHeight) % iHeight
				} else {
					continue
				}
			}

			cell := g.GetCell(uint(nx), uint(ny))
			neighbors = append(neighbors, cell)
		}
	}

	return neighbors
}

func (g *Grid) Rows() (rows [][]Cell) {
	var currentRow []Cell

	for i, c := range *g.Cells {
		// If first of row, push new row
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
	for i := uint(0); i < g.Width; i++ {
		cols = append(cols, []Cell{})
	}

	for _, c := range *g.Cells {
		cols[c.X] = append(cols[c.X], c)
	}

	return cols
}
