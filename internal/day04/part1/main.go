package day04_part1

import (
	"fmt"

	"github.com/Baipyrus/AoC-25/internal/day04"
	"github.com/Baipyrus/AoC-25/internal/registry"
)

var name = "Day 04 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	var paperRolls uint
	grid := day04.ParseInput(input, true)

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
			}
		}
	}

	fmt.Printf("The forklift can access a total of %d rolls of paper.\n", paperRolls)
}
