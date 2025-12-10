package day05_part1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Baipyrus/AoC-25/internal/registry"
)

var name = "Day 05 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	segments := strings.Split(input, "\n\n")
	freshRanges := strings.Split(strings.TrimSpace(segments[0]), "\n")
	ingredients := strings.Split(strings.TrimSpace(segments[1]), "\n")

	var freshIngredients uint64
	for _, ingredient := range ingredients {
		ingredientId, _ := strconv.ParseUint(ingredient, 10, 64)

		for _, freshRange := range freshRanges {
			bounds := strings.Split(freshRange, "-")
			lower, _ := strconv.ParseUint(bounds[0], 10, 64)
			upper, _ := strconv.ParseUint(bounds[1], 10, 64)

			if ingredientId >= lower && ingredientId <= upper {
				freshIngredients++
				break
			}
		}
	}

	fmt.Printf("There are a total of %d fresh ingredients left.\n", freshIngredients)
}
