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
			if len(currentId)%2 != 0 {
				continue
			}

			// Split the ID into two
			midway := len(currentId) / 2
			left := currentId[:midway]
			right := currentId[midway:]

			// The halves are equal looking
			if left == right {
				invalidIdSum += i
			}
		}
	}

	fmt.Printf("The sum total of all invalid IDs is: %d\n", invalidIdSum)
}
