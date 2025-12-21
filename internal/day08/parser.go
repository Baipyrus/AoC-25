package day08

import (
	"strconv"
	"strings"
)

func ParseInput(input string) (boxes []Point) {
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		bline := strings.TrimSpace(line)
		if bline == "" {
			continue
		}

		coordinates := strings.Split(bline, ",")
		x, _ := strconv.ParseUint(coordinates[0], 10, 64)
		y, _ := strconv.ParseUint(coordinates[1], 10, 64)
		z, _ := strconv.ParseUint(coordinates[2], 10, 64)

		next := Point{X: x, Y: y, Z: z, Idx: i}
		next.Mag = next.Magnitude()
		boxes = append(boxes, next)
	}

	return boxes
}
