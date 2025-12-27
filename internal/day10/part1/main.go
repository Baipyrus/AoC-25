package day10_part1

import (
	"fmt"
	"math"
	"slices"

	"github.com/Baipyrus/AoC-25/internal/day10"
	"github.com/Baipyrus/AoC-25/internal/registry"
)

var name = "Day 10 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	machines := day10.ParseInput(input)

	var sum uint
	for _, m := range machines {
		sum += SolveMachine(m)
	}

	fmt.Printf("The sum of the fewest button presses to solve each machine is: %d\n", sum)
}

type EnqueuedState struct {
	State day10.MachineState
	Steps uint
}

// Solve this machine from an empty state all the way to
// the goal state in the fewest possible button presses:
func SolveMachine(m *day10.Machine) uint {
	// Save goal in readable variable, initialize empty starting state
	goal := m.State
	start := make(day10.MachineState, len(goal))

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
