package day03_part1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Baipyrus/AoC-25/internal/registry"
)

var name = "Day 03 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	var joltageSum uint64

	lines := strings.SplitSeq(input, "\n")
	for line := range lines {
		batteryBank := strings.TrimSpace(line)
		if batteryBank == "" {
			continue
		}

		var maxJoltage uint64
		for i := 0; i < len(batteryBank)-1; i++ {
			leftBattery := string(batteryBank[i])

			for j := i + 1; j < len(batteryBank); j++ {
				rightBattery := string(batteryBank[j])

				// Combine the left and right batteries to get joltage
				joltage, _ := strconv.ParseUint(string(leftBattery+rightBattery), 10, 64)

				// Compare with last maximum joltage
				if joltage > maxJoltage {
					maxJoltage = joltage
				}
			}
		}

		joltageSum += maxJoltage
	}

	fmt.Printf("The sum total of all battery's max joltages is: %d\n", joltageSum)
}
