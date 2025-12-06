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
		// Reverse steps if DIRECTION == LEFT
		steps := int64(inst.Steps) * int64(inst.Dir)

		// Dial rotation wrap-around
		dial = uint((int64(dial) + steps + 100) % 100)

		// Detect zero positions for use in password
		if dial == 0 {
			zeroes++
		}
	}

	fmt.Printf("The password to open the door is: '%d'", zeroes)
}
