package day08_part1

import (
	"fmt"

	"github.com/Baipyrus/AoC-25/internal/registry"
)

var name = "Day 08 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)
}
