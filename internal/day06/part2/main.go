package day06_part2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Baipyrus/AoC-25/internal/day04"
	"github.com/Baipyrus/AoC-25/internal/day06"
	"github.com/Baipyrus/AoC-25/internal/registry"
)

var name = "Day 06 - Part 2"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	grid := day04.ParseInput(input, false)

	var (
		eqNumbers []int64
		equations []*day06.Equation
	)

numberParser:
	for reverseX := uint(0); reverseX < grid.Width; reverseX++ {
		var number string

		for y := uint(0); y < grid.Height; y++ {
			cell := grid.GetCell(grid.Width-reverseX-1, y)

			if cell.Type == '+' || cell.Type == '*' {
				parsedNumber, _ := strconv.ParseInt(strings.TrimSpace(number), 10, 64)

				equations = append(equations, &day06.Equation{
					Numbers:  append(eqNumbers, parsedNumber),
					Operator: cell.Type,
				})

				eqNumbers = []int64{}
				continue numberParser
			}

			number += string(cell.Type)
		}

		if strings.TrimSpace(number) == "" {
			continue
		}

		parsedNumber, _ := strconv.ParseInt(strings.TrimSpace(number), 10, 64)
		eqNumbers = append(eqNumbers, parsedNumber)
	}

	var sumTotal int64
	for _, eq := range equations {
		value := eq.Evaluate()
		sumTotal += value
	}

	fmt.Printf("The sum total of all evaluated equations is: %d\n", sumTotal)
}
