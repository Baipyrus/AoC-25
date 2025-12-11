package day04_part2

import (
	"fmt"

	"github.com/Baipyrus/AoC-25/internal/day04"
	"github.com/Baipyrus/AoC-25/internal/registry"
)

var name = "Day 04 - Part 2"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	var (
		paperRolls uint
		newGrid    day04.Grid
	)

	grid := day04.ParseInput(input, true)
	gridHasChanged := true

	// Lazy solution: simply redo this loop
	// to scan the entire grid again and
	// find newly accessible paper rolls.
	for gridHasChanged {
		// Assume this will be the last iteration:
		// The grid will no longer change, the loop
		// will end after this final pass.
		gridHasChanged = false
		newGrid = grid

		for y := range grid.Height {
			for x := range grid.Width {
				cell := grid.GetCell(x, y)
				if cell.Type != '@' {
					continue
				}

				neighbors := grid.GetNeighbors(x, y, false, true)

				var neighborCount uint
				for _, n := range neighbors {
					if n.Type == '@' {
						neighborCount++
					}
				}

				if neighborCount < 4 {
					paperRolls++

					// Removing the old paper roll that
					// has been removed by the forklift
					cell.Type = '.'
					newGrid.SetCell(cell)

					// Another accessible paper roll
					// has been found! We need to redo
					// the loop and find new spots.
					gridHasChanged = true
				}
			}
		}
	}

	fmt.Printf("The forklift can access a total of %d rolls of paper.\n", paperRolls)
}
