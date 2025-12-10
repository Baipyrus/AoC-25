package day06_part1

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/Baipyrus/AoC-25/internal/day06"
	"github.com/Baipyrus/AoC-25/internal/registry"
)

var name = "Day 06 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	whitespace := regexp.MustCompile(`\s+`)
	var equations []*day06.Equation

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for i, line := range lines {
		isOperatorLine := len(lines)-1 == i
		columns := slices.DeleteFunc(
			whitespace.Split(line, -1),
			func(s string) bool { return s == "" })

		for j, entry := range columns {
			if len(equations) <= j {
				equations = append(equations, &day06.Equation{})
			}

			eq := equations[j]
			if isOperatorLine {
				eq.Operator = entry
				// This *could* be a `break` since it's
				// always last by definition, but it
				// makes for little to no difference.
				continue
			}

			num, _ := strconv.ParseInt(entry, 10, 64)
			eq.Numbers = append(eq.Numbers, num)
		}
	}

	var sumTotal int64
	for _, eq := range equations {
		value := eq.Evaluate()
		sumTotal += value
	}

	fmt.Printf("The sum total of all evaluated equations is: %d\n", sumTotal)
}
