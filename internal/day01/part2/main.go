package day01_part2

import (
	"fmt"

	"github.com/Baipyrus/AoC-25/internal/registry"
)

var name = "Day 01 - Part 2"

func init() {
	registry.Register(name, Main)
}

var (
	dial   uint
	zeroes uint
)

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)
}
