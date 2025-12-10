package day06_part2

import (
	"fmt"

	"github.com/Baipyrus/AoC-25/internal/registry"
)

var name = "Day 06 - Part 2"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)
}
