package day07_part2

import (
	"fmt"

	"github.com/Baipyrus/AoC-25/internal/day04"
	"github.com/Baipyrus/AoC-25/internal/registry"
)

var name = "Day 07 - Part 2"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	grid := day04.ParseInput(input, true)
	manifold := findTachyonManifold(grid)

	timelineMemory := make(map[uint]uint64)
	timelineCount := shootTachyonBeam(&grid, manifold, &timelineMemory)

	fmt.Printf("The tachyon beam splitting has created a total of %d timelines.\n", timelineCount)
}

func shootTachyonBeam(grid *day04.Grid, manifold day04.Cell, memory *map[uint]uint64) (timelineCount uint64) {
	startIdx := manifold.Idx
	if count := (*memory)[startIdx]; count > 0 {
		return count
	}
	currentCell := manifold

reachBottomLoop:
	for true {
		if currentCell.Y+1 == grid.Height {
			timelineCount++
			break reachBottomLoop
		}

		nextCell := grid.GetCell(currentCell.X, currentCell.Y+1)
		switch nextCell.Type {
		case '.':
			currentCell = nextCell
			continue reachBottomLoop
		case 'S':
			// WARNING: This should never be possible, as beams only travel
			//          downwards and there can only be one manifold at which
			//          the process should be started on!
			return 0
		}
		// NOTE: We should only ever get here if the next cell is a splitter!

		// Split beam to the left of this splitter
		if nextCell.X > 0 {
			leftSplit := grid.GetCell(nextCell.X-1, nextCell.Y)
			timelineCount += shootTachyonBeam(grid, leftSplit, memory)
		}

		// Split beam to the right of this splitter
		if nextCell.X < grid.Width-1 {
			rightSplit := grid.GetCell(nextCell.X+1, nextCell.Y)
			timelineCount += shootTachyonBeam(grid, rightSplit, memory)
		}

		// End the current beam (successfully hit splitter)
		break reachBottomLoop
	}

	(*memory)[startIdx] = timelineCount
	return timelineCount
}

func findTachyonManifold(grid day04.Grid) (manifold day04.Cell) {
manifoldSearch:
	for y := range grid.Height {
		for x := range grid.Width {
			cell := grid.GetCell(x, y)
			if cell.Type != 'S' {
				continue
			}

			manifold = cell
			break manifoldSearch
		}
	}

	return manifold
}
