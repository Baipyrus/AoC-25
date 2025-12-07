package day01_part2

import (
	"fmt"

	"github.com/Baipyrus/AoC-25/internal/day01"
	"github.com/Baipyrus/AoC-25/internal/registry"
)

var name = "Day 01 - Part 2"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	var zeroes uint
	instructions, dial := day01.ParseInput(input)

	for _, inst := range instructions {
		// Calculate least required steps to reach 0:
		// If already on 0, another 100 steps to either
		// side are required to reach another 0 again.
		var required uint = 100

		// If the dial is moving up ("positive steps"),
		// then 100 - dial equals the required steps.
		if inst.Dir == day01.RIGHT {
			required -= dial
		} else if inst.Dir == day01.LEFT && dial != 0 {
			// Lastly, simply step #dial times to the left.
			required = dial
		}

		// If steps taken is more than least required,
		// at least one zero is counted. Number of zeroes
		// can be calculated by deviding the difference.
		if inst.Steps >= required {
			zeroes += 1 + (inst.Steps-required)/100
		}

		multiplier := inst.Steps/100 + 1
		// => How many times has the dial wrapped around?
		//    (assuming that it has turned at least once)

		// Reverse steps if DIRECTION == LEFT
		steps := int64(inst.Steps) * int64(inst.Dir)

		// Dial rotation wrap-around
		dial = uint((int64(dial) + steps + 100*int64(multiplier)) % 100)
	}

	fmt.Printf("The password to open the door using the 0x434C49434B method is: '%d'", zeroes)
}
