package day01_part1

import (
	"fmt"

	"github.com/Baipyrus/AoC-25/internal/day01"
	"github.com/Baipyrus/AoC-25/internal/registry"
)

var name = "Day 01 - Part 1"

func init() {
	registry.Register(name, Main)
}

// Wrap-Around Modulo:
//
// ( 0 - 1 + 100) % 100 = 99
// ( 0 + 1 + 100) % 100 =  1
//
// (99 - 1 + 100) % 100 = 98
// (99 + 1 + 100) % 100 =  0

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	var zeroes uint
	instructions, dial := day01.ParseInput(input)

	for _, inst := range instructions {
		multiplier := inst.Steps/100 + 1
		// => How many times has the dial wrapped around?
		//    (assuming that it has turned at least once)

		// Reverse steps if DIRECTION == LEFT
		steps := int64(inst.Steps) * int64(inst.Dir)

		// Dial rotation wrap-around
		dial = uint((int64(dial) + steps + 100*int64(multiplier)) % 100)

		// Detect zero positions for use in password
		if dial == 0 {
			zeroes++
		}
	}

	fmt.Printf("The password to open the door is: '%d'\n", zeroes)
}
