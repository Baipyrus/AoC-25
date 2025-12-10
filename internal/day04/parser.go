package day04

import (
	"strings"
)

func ParseInput(input string, trimSpaces bool) (grid Grid) {
	lines := strings.SplitSeq(input, "\n")

	var (
		rows   []string
		width  uint
		height uint
	)

	for line := range lines {
		row := strings.TrimSpace(line)

		if width == 0 {
			if trimSpaces {
				width = uint(len(row))
			} else {
				width = uint(len(line))
			}
		}

		// Technically, if the current width does not match
		// the previously recorded width, the program should
		// error anyways because it could be an incomplete input
		// or similar ... So we just ignore it here.
		if row == "" {
			continue
		}

		if trimSpaces {
			rows = append(rows, row)
		} else {
			rows = append(rows, line)
		}
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
