package day04

import (
	"strings"
)

func ParseInput(input string) (grid Grid) {
	lines := strings.SplitSeq(input, "\n")

	var (
		rows   []string
		width  uint
		height uint
	)

	for line := range lines {
		row := strings.TrimSpace(line)
		currentWidth := uint(len(row))

		if width == 0 {
			width = currentWidth
		}

		if row == "" || currentWidth != width {
			// Technically, if the width does not match,
			// this should error because it could be an
			// incomplete input or similar ...
			continue
		}

		rows = append(rows, row)
	}
	height = uint(len(rows))

	grid = Grid{Width: width, Height: height, Cells: &[]Cell{}}
	for y, row := range rows {
		cells := strings.Split(row, "")

		for x, c := range cells {
			cell := Cell{
				Idx:  uint(x) + uint(y)*width,
				X:    uint(x),
				Y:    uint(y),
				Type: rune(c[0]),
			}

			*grid.Cells = append(*grid.Cells, cell)
		}
	}

	return grid
}
