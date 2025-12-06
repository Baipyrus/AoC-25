package day01

import (
	"strconv"
	"strings"
)

func ParseInput(input string) ([]Instruction, uint) {
	var (
		instructions []Instruction
		dialPosition uint = 50
	)

	lines := strings.SplitSeq(input, "\n")
	for line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		direction := line[:1]
		steps, _ := strconv.ParseUint(line[1:], 10, 64)

		var d Direction
		switch strings.ToUpper(direction) {
		case "L":
			d = LEFT
		case "R":
			d = RIGHT
		}

		instructions = append(instructions, Instruction{Dir: d, Steps: uint(steps)})
	}

	return instructions, dialPosition
}
