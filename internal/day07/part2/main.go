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
	splittersHit := shootTachyonBeam(&grid, manifold)

	fmt.Printf("The tachyon beam was split a total of %d times.\n", splittersHit)
}

func shootTachyonBeam(grid *day04.Grid, manifold day04.Cell) (splittersHit uint64) {
	// WARNING: Manifold must not be empty space or splitter!
	if !(manifold.Type == 'S' || manifold.Type == '|') {
		return 0
	}
	currentCell := manifold

reachBottomLoop:
	for true {
		if currentCell.Y+1 == grid.Height {
			break reachBottomLoop
		}

		nextCell := grid.GetCell(currentCell.X, currentCell.Y+1)
		switch nextCell.Type {
		case '.':
			nextCell.Type = '|'
			grid.SetCell(nextCell)

			currentCell = nextCell
			continue reachBottomLoop
		case '|':
			break reachBottomLoop
		case 'S':
			// WARNING: This should never be possible, as beams only travel
			//          downwards and there can only be one manifold at which
			//          the process should be started on!
			return 0
		}

		// NOTE: We should only ever get here if the next cell is a splitter!
		splittersHit++

		// Split beam to the left of this splitter
		if nextCell.X > 0 {
			splittersHit += setTachyonBeam(grid, nextCell.X-1, nextCell.Y)
		}

		// Split beam to the right of this splitter
		if nextCell.X < grid.Width {
			splittersHit += setTachyonBeam(grid, nextCell.X+1, nextCell.Y)
		}

		// End the current beam (successfully hit splitter)
		break reachBottomLoop
	}

	return splittersHit
}

func setTachyonBeam(grid *day04.Grid, x, y uint) uint64 {
	cell := grid.GetCell(x, y)

	if cell.Type == '.' {
		cell.Type = '|'
		grid.SetCell(cell)

		return shootTachyonBeam(grid, cell)
	}

	return 0
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
