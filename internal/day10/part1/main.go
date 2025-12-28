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
	// NOTE: This approach should be identical to running breadth-first search
	//       on a tree that represents all the different mutations that a
	//       machine's buttons can produce on an initial state.
	//       (This also means there are more optimizations to be made. For example,
	//       you could simply try using a different algorithm like Dijkstra's!)

	// Save goal in readable variable, initialize empty starting state
	var start day10.MachineState
	goal := m.State

	// Remember which states have already been calculated, because
	// it will always result in the same follow-up action!
	// NOTE: Although the actual boolean value of this map
	//       (true => solution) is never *actually* used
	visited := make(map[day10.MachineState]bool)

	// Initialize queue with empty state as its first entry
	// NOTE: This is a very mild optimization compared to the
	//       dynamic array from before this change. I have also
	//       tried using a 3rd party FIFO queue, but it only
	//       managed to improve another ~3ms on large inputs.
	//       The real problem with this large static array is
	//       that small inputs will always carry this unnecessary
	//       mass with them, which slows them down...

	var queue [16384]EnqueuedState // <- len: 2^14
	queue[0] = EnqueuedState{State: start, Steps: 0}
	firstPos := 0
	lastPos := 0

	// FIFO Queue: (FIRST IN, FIRST OUT).
	// Take the first element in the queue, find what's wrong with it,
	// find a way to correct the state one step closer to the goal.
	// Repeat until current state equals goal state.
	for firstPos <= lastPos {
		current := queue[firstPos]

		// "Delete" and forget about queue entry
		queue[firstPos] = EnqueuedState{}
		firstPos++

		// If they are equal, this machine is solved!
		if current.State == goal {
			visited[current.State] = true
			return current.Steps
		}

		// Find first indicator light that is unequal to its goal position
		var firstUnequal uint
		for i := range day10.BIT_PACKING_SIZE {
			if goal.At(i) == current.State.At(i) {
				continue
			}

			firstUnequal = i
			break
		}

		// Find all buttons that change at least this one indicator
		uFirst := uint(firstUnequal)
		for _, b := range m.Buttons {
			if !slices.Contains(b.Sequence, uFirst) {
				continue
			}

			// With this button, mutate state and enqueue it as the
			// next child item of this tree node!
			cpyState := current.State
			cpyState.Mutate(b)

			// Ignore mutated state if it's already been computed before.
			// This prevents loops and similarities in sub-branches!
			if _, ok := visited[cpyState]; ok {
				continue
			}

			// Increment last queue position index first because
			// it is now the new index of the next queue entry!
			lastPos++

			// Enqueue newly mutated state as next step
			queue[lastPos] = EnqueuedState{State: cpyState, Steps: current.Steps + 1}
		}

		visited[current.State] = false
	}

	// No solution to reach the goal could be found!
	return math.MaxUint
}
