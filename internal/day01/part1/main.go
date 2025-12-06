package day01_part1

import (
	"fmt"
	"strconv"
	"strings"

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

var (
	dial   uint = 50
	zeroes uint
)

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	// Parse input lines
	lines := strings.SplitSeq(input, "\n")
	for line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		// Each lines contains: "[Direction (Left / Right)][Steps (0,1,2,...)]"
		direction := line[:1]
		steps, _ := strconv.ParseInt(line[1:], 10, 64)

		// Calculate new dial rotation with wrap-around
		if strings.ToUpper(direction) == "L" {
			steps *= -1
		}
		dial = uint((int64(dial) + steps + 100) % 100)

		// Detect zero position for use in password
		if dial == 0 {
			zeroes++
		}
	}

	fmt.Printf("The password to open the door is: '%d'", zeroes)
}
