package day10_part1

import (
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/Baipyrus/AoC-25/internal/registry"
)

var name = "Day 10 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	machines := ParseInput(input)

	var sum uint
	for _, m := range machines {
		sum += m.Solve()
	}

	fmt.Printf("The sum of the fewest button presses to solve each machine is: %d\n", sum)
}

type EnqueuedState struct {
	State MachineState
	Steps uint
}

// Solve this machine from an empty state all the way to
// the goal state in the fewest possible button presses:
func (m *Machine) Solve() uint {
	// Save goal in readable variable, initialize empty starting state
	goal := m.State
	start := make(MachineState, len(goal))

	// Initialize queue with empty state as its first entry
	// NOTE: This could be optimized with an actual FIFO queue
	queue := []EnqueuedState{{State: start, Steps: 0}}

	// FIFO Queue: (FIRST IN, FIRST OUT).
	// Take the first element in the queue, find what's wrong with it,
	// find a way to correct the state one step closer to the goal.
	// Repeat until current state equals goal state.
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Find first indicator light that is unequal to its goal position
		firstUnequal := -1
		for i, s := range goal {
			if s == current.State[i] {
				continue
			}

			firstUnequal = i
			break
		}

		// If all are equal, this machine is solved!
		if firstUnequal == -1 {
			return current.Steps
		}

		// Find all buttons that change at least this one indicator
		uFirst := uint(firstUnequal)
		for _, b := range m.Buttons {
			if !slices.Contains(b.Sequence, uFirst) {
				continue
			}

			// With this button, mutate state and enqueue it as the
			// next child item of this tree node!
			cpyState := current.State.Copy()
			cpyState.Mutate(b)

			queue = append(queue, EnqueuedState{State: cpyState, Steps: current.Steps + 1})
		}
	}

	// No solution to reach the goal could be found!
	return math.MaxUint
}

func ParseInput(input string) (machines []*Machine) {
	lines := strings.SplitSeq(input, "\n")

	for line := range lines {
		bline := strings.TrimSpace(line)
		if bline == "" {
			continue
		}

		var machine Machine

		// NOTE: I know, I know ... I don't usually do error
		//       handling in AoC entries like this, but I was
		//       already writing all this "clean" looking
		//       boilerplate and so I just had to keep passing
		//       the error upward! :D
		err := machine.Deserialize(bline)
		if err != nil {
			log.Fatal(err)
		}

		machines = append(machines, &machine)
	}

	return machines
}

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

type Machine struct {
	State    MachineState
	Buttons  []Button
	Joltages string
}

func (m *Machine) Deserialize(serialized string) error {
	parts := strings.Split(serialized, " ")

	for i, p := range parts {
		startEnd := string(p[0]) + string(p[len(p)-1])
		switch startEnd {
		case "[]":
			// NOTE: Machine state *should* only appear once
			//       per serialized machine string!!!
			machineStateGoal, err := NewMachineState(p)
			if err != nil {
				return fmt.Errorf("Failed to parse state for machine at %d ('%s'): %w", i, p, err)
			}

			m.State = machineStateGoal
		case "()":
			b, err := NewButton(p)
			if err != nil {
				return fmt.Errorf("Failed to parse button for machine at %d ('%s'): %w", i, p, err)
			}

			m.Buttons = append(m.Buttons, b)
		case "{}":
			// NOTE: Machine joltages *should* only appear once
			//       per serialized machine string!!!
			m.Joltages = p
		default:
			return fmt.Errorf("Unknown machine syntax!")
		}
	}

	// NOTE: Technically speaking, you'd have to check if all the
	//       buttons actually reference existing indices for state
	//       indicator lights here ... D:

	return nil
}

func (m Machine) String() string {
	var buttons string
	for _, b := range m.Buttons {
		buttons += fmt.Sprintf("%s ", b)
	}
	buttons = strings.TrimSpace(buttons)

	return fmt.Sprintf("%s %s %s", m.State, buttons, m.Joltages)
}
