package day10

import (
	"fmt"
	"slices"
	"strings"
)

// WARNING: DO NOT edit this size! It is directly linked to
//          the usage as 'MachineState' as type uint64!

const BIT_PACKING_SIZE = uint(16)

func NewMachineState(serialized string) (out MachineState, _ error) {
	startEnd := string(serialized[0]) + string(serialized[len(serialized)-1])
	if startEnd != "[]" {
		return out, fmt.Errorf("Unknown machine state format!")
	}

	state := serialized[1 : len(serialized)-1]
	total := len(state)

	if uint(total) > BIT_PACKING_SIZE {
		return out, fmt.Errorf("Unable to pack any more bits into type uint64 (want: %d)!", total)
	}

	toggles := strings.Split(state, "")
	for i, t := range toggles {
		switch t {
		case ".":
			// Since we're bit-packing the state in a 64-bit uint,
			// we do not need to concern ourselves with "false" state
			// since the uint is initialized as zero anyways!
		case "#":
			out ^= 1 << i
		default:
			return out, fmt.Errorf("Failed to parse machine state at %d ('%s')!", i, t)
		}
	}

	return out, nil
}

type MachineState uint16

// Flip all indicators in a state according to a button
func (ms *MachineState) Mutate(b Button) {
	for i := range BIT_PACKING_SIZE {
		if slices.Contains(b.Sequence, i) {
			*ms ^= 1 << i
		}
	}
}

func (ms *MachineState) At(i uint) bool {
	return (*ms & (1 << i)) > 0
}

func (ms MachineState) String() (out string) {
	for i := range BIT_PACKING_SIZE {
		state := ms.At(i)

		switch state {
		case true:
			out += "#"
		case false:
			out += "."
		}
	}

	return fmt.Sprintf("[%s]", out)
}
