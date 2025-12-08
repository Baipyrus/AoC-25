package day03

import "fmt"

type Battery struct {
	Joltage uint64
	Index   int
}

type BatteryBank struct {
	Batteries  []Battery
	MaxJoltage uint64
}

func (bank BatteryBank) String() (out string) {
	for _, battery := range bank.Batteries {
		out += fmt.Sprint(battery.Joltage)
	}

	return out
}
