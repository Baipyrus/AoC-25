package day03_part1

import (
	"fmt"
	"strconv"

	"github.com/Baipyrus/AoC-25/internal/day03"
	"github.com/Baipyrus/AoC-25/internal/registry"
)

var name = "Day 03 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	var joltageSum uint64
	batteryBanks := day03.ParseInput(input)

	for _, batteryBank := range batteryBanks {
		batteries := batteryBank.Batteries

		var maxJoltage uint64
		for i := 0; i < len(batteries)-1; i++ {
			leftBattery := fmt.Sprint(batteries[i].Joltage)

			for j := i + 1; j < len(batteries); j++ {
				rightBattery := fmt.Sprint(batteries[j].Joltage)

				// Combine the left and right batteries to get joltage
				joltage, _ := strconv.ParseUint(leftBattery+rightBattery, 10, 64)

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
