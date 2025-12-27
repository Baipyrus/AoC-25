package day10

import (
	"fmt"
	"slices"
	"strings"
)

func NewMachineState(serialized string) (out MachineState, _ error) {
	startEnd := string(serialized[0]) + string(serialized[len(serialized)-1])
	if startEnd != "[]" {
		return nil, fmt.Errorf("Unknown machine state format!")
	}

	state := serialized[1 : len(serialized)-1]
	toggles := strings.Split(state, "")
	for i, t := range toggles {
		switch t {
		case ".":
			out = append(out, false)
		case "#":
			out = append(out, true)
		default:
			return out, fmt.Errorf("Failed to parse machine state at %d ('%s')!", i, t)
		}
	}

	return out, nil
}

type MachineState []bool

func (ms *MachineState) Copy() MachineState {
	cpyState := make(MachineState, len(*ms))
	copy(cpyState, *ms)

	return cpyState
}

// Flip all indicators in a state according to a button
func (ms *MachineState) Mutate(b Button) {
	for i := range *ms {
		if slices.Contains(b.Sequence, uint(i)) {
			(*ms)[i] = !(*ms)[i]
		}
	}
}

func (ms MachineState) String() (out string) {
	for _, s := range ms {
		switch s {
		case true:
			out += "#"
		case false:
			out += "."
		}
	}

	return fmt.Sprintf("[%s]", out)
}
