package day02_part2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Baipyrus/AoC-25/internal/registry"
)

var name = "Day 02 - Part 2"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	var invalidIdSum uint64

	idRanges := strings.SplitSeq(input, ",")
	for ids := range idRanges {
		bounds := strings.Split(strings.TrimSpace(ids), "-")
		if len(bounds) != 2 {
			continue
		}

		lower, _ := strconv.ParseUint(bounds[0], 10, 64)
		upper, _ := strconv.ParseUint(bounds[1], 10, 64)

		// For lack of a better approach, simply loop
		// through the entire range of IDs to find
		// self-similar-ish looking IDs ...
		for i := lower; i <= upper; i++ {
			// A left/right symmetrical ID should
			// always be of even length
			currentId := strconv.FormatUint(i, 10)
			idLength := len(currentId)

			selfSimilar := false
		patternSize:
			for j := idLength / 2; j >= 1; j-- {
				// If idLength is not divisible into chunks of
				// length j, then length j cannot form a pattern.
				if j*(idLength/j) != idLength {
					continue
				}

				// First substring of ID to recognize pattern with
				pattern := currentId[:j]

				for k := 1; k < idLength/j; k++ {
					// Get next substring in ID
					nextSubstr := currentId[j*k : j*(k+1)]

					if nextSubstr != pattern {
						continue patternSize
					}
				}

				selfSimilar = true
				break
			}

			if selfSimilar {
				invalidIdSum += i
			}
		}
	}

	fmt.Printf("The sum total of all invalid IDs is: %d\n", invalidIdSum)
}
