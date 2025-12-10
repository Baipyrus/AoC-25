package day05_part2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Baipyrus/AoC-25/internal/registry"
)

var name = "Day 05 - Part 2"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	// Get fresh id ranges from input
	segments := strings.Split(input, "\n\n")
	inputRanges := strings.Split(strings.TrimSpace(segments[0]), "\n")

	// Parse input into range list
	var freshRanges []*Range
	for _, freshRange := range inputRanges {
		bounds := strings.Split(freshRange, "-")
		currentLower, _ := strconv.ParseUint(bounds[0], 10, 64)
		currentUpper, _ := strconv.ParseUint(bounds[1], 10, 64)

		freshRanges = append(freshRanges, &Range{Lower: currentLower, Upper: currentUpper})
	}

	// Calculate all unique fresh ID ranges
	uniqueRanges := calculateUniqueRanges(freshRanges)

	// Count all fresh IDs within all unique ranges
	var uniqueFreshIds uint64
	for _, freshRange := range uniqueRanges {
		uniqueFreshIds += freshRange.Upper - freshRange.Lower + 1
	}

	fmt.Printf("The total number of fresh ingredient IDs considered to be fresh is: %d\n", uniqueFreshIds)
}

func calculateUniqueRanges(rawRanges []*Range) (uniqueRanges []*Range) {
	// Loop through every range, find all uniques
	for i, self := range rawRanges {
		// If not unique, try merging with another
		isNewUnique := true
		for _, other := range uniqueRanges {
			if other.TryMergingWith(self) {
				isNewUnique = false
				break
			}
		}

		if isNewUnique {
			uniqueRanges = append(uniqueRanges, self)
			continue
		}

		// If we got here, that means that 'other'
		// has been merged with 'self', meaning that
		// its bounds have changed and thusly it may
		// cause overlap with future or previous ranges.
		// Since we have to reevaluate anyways, we just
		// cut the process in half and start over!
		return calculateUniqueRanges(append(uniqueRanges, rawRanges[i+1:]...))
	}

	return uniqueRanges
}

type Range struct {
	Lower uint64
	Upper uint64
}

func (self *Range) TryMergingWith(other *Range) bool {
	// *Some* overlap must exist because either of 'self's boundaries
	// must be overlapping with either of the 'other's boundaries.
	if self.Lower <= other.Upper && self.Upper >= other.Lower {
		// Account for LOWER boundary overlap
		if other.Lower < self.Lower {
			self.Lower = other.Lower
		}
		// Account for UPPER boundary overlap
		if other.Upper > self.Upper {
			self.Upper = other.Upper
		}

		// TOTAL overlap ('other' is inside 'self')
		return true
	}

	return false
}
