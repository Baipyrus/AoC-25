package day09_part2

import (
	"fmt"
	"slices"

	"github.com/Baipyrus/AoC-25/internal/day08"
	"github.com/Baipyrus/AoC-25/internal/registry"
)

var name = "Day 09 - Part 2"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	boxes := day08.ParseInput(input)
	furthest := day08.Combinations(boxes)

	// Sort all combinations such that we can start with
	// the two points that are furthest from each other!
	slices.SortFunc(
		furthest,
		func(a, b day08.Rectangle) int {
			return int(b.Area - a.Area)
		})

	fmt.Printf("The largest area of any of the available rectangles is: %d\n", furthest[0].Area)
}
