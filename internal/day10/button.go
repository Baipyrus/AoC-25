package day10

import (
	"fmt"
	"strconv"
	"strings"
)

func NewButton(serialized string) (out Button, _ error) {
	startEnd := string(serialized[0]) + string(serialized[len(serialized)-1])
	if startEnd != "()" {
		return out, fmt.Errorf("Unknown button format!")
	}

	sequence := serialized[1 : len(serialized)-1]
	indices := strings.Split(sequence, ",")
	for i, idx := range indices {
		uIdx, err := strconv.ParseUint(idx, 10, 32)
		if err != nil {
			return out, fmt.Errorf("Failed to parse button index at %d ('%s'): %w", i, idx, err)
		}

		out.Sequence = append(out.Sequence, uint(uIdx))
	}

	return out, nil
}

type Button struct {
	Sequence []uint
}

func (b Button) String() (out string) {
	for _, idx := range b.Sequence {
		out += fmt.Sprintf("%d,", idx)
	}

	return fmt.Sprintf("(%s)", out[:len(out)-1])
}
