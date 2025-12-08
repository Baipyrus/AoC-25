package day03_part2

import (
	"fmt"
	"strconv"

	"github.com/Baipyrus/AoC-25/internal/day03"
	"github.com/Baipyrus/AoC-25/internal/registry"
)

var name = "Day 03 - Part 2"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	var joltageSum uint64

	// WARNING: Logically speaking, this value should only ever
	//          go as low as the length of 2.
	// NOTE: If you *do* set the valaue to 2, you should expect
	//       the same result as from day03_part1!
	combinationSize := 12

	batteryBanks := day03.ParseInput(input)

	for _, batteryBank := range batteryBanks {
		batteries := batteryBank.Batteries

		var maxJoltage uint64
		// We want a combination of length 12. Meaning the last
		// possible position for the first number is index at
		// length len - (12 - 1) because we want 11 more numbers.
		for i := 0; i < len(batteries)-(combinationSize-1); i++ {
			combinedBatteries := fmt.Sprint(batteries[i].Joltage)
			lastIndex := i

			// We want to choose 11 more digits. We need to subtract
			// 1 because it's an index, not a length, and another 1
			// because we already have one digit ready.
			for j := combinationSize - 2; j >= 0; j-- {
				var largestBattery uint64

				// The last possible digit would be closest to
				// the end but offset by j positions to leave
				// enough room for the other digits
				for k := lastIndex + 1; k < len(batteries)-j; k++ {
					currentBattery := batteries[k].Joltage

					if currentBattery > largestBattery {
						largestBattery = currentBattery
						lastIndex = k
					}
				}

				combinedBatteries += fmt.Sprint(largestBattery)
			}

			// Lastly, convert the newly found combined jotage and compare it
			currentJoltage, _ := strconv.ParseUint(combinedBatteries, 10, 64)
			if currentJoltage > maxJoltage {
				maxJoltage = currentJoltage
			}
		}

		joltageSum += maxJoltage
	}

	fmt.Printf("The sum total of all battery's max joltages is: %d\n", joltageSum)
}
