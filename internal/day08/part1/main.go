package day08_part1

import (
	"fmt"
	"slices"

	"github.com/Baipyrus/AoC-25/internal/day08"
	"github.com/Baipyrus/AoC-25/internal/registry"
)

var name = "Day 08 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	var connections int = 1000

	boxes := day08.ParseInput(input)
	closest := day08.Combinations(boxes)

	// Sort all combinations such that we can start with
	// the two points that are closest to each other!
	slices.SortFunc(
		closest,
		func(a, b day08.Rectangle) int {
			return int(a.Diagonal - b.Diagonal)
		})

	// NOTE: IMPORTANT! We only want to make n connections!
	circuits := connectPoints(closest[:connections])

	// For our final solution, we want to multiply the three
	// largest circuit's lengths with each other:
	slices.SortFunc(
		circuits,
		func(a, b *day08.Circuit) int {
			return len(b.Boxes) - len(a.Boxes)
		})

	product := len(circuits[0].Boxes) * len(circuits[1].Boxes) * len(circuits[2].Boxes)
	fmt.Printf("The product of multiplying the length of the three largest circuits is: %d\n", product)
}

func mergeCircuits(cache *map[int]*day08.Circuit, a, b *day08.Circuit) {
	a.Boxes = append(a.Boxes, b.Boxes...)

	// Update cache references
	for _, box := range b.Boxes {
		(*cache)[box.Idx] = a
	}

	// "Delete" b by resetting its boxes
	b.Boxes = []day08.Point{}
}

func connectPoints(points []day08.Rectangle) (circuits []*day08.Circuit) {
	// We want to remember which box is part
	// of which circuit without having to search!
	connected := make(map[int]*day08.Circuit)

	for _, pd := range points {
		aCache := connected[pd.A.Idx]
		bCache := connected[pd.B.Idx]

		// If none of the points have been connected yet,
		// create a new circuit and cache their connection!
		if aCache == nil && bCache == nil {
			circuitIdx := len(circuits)
			newCircuit := &day08.Circuit{Boxes: []day08.Point{pd.A, pd.B}, Idx: circuitIdx}

			connected[pd.A.Idx] = newCircuit
			connected[pd.B.Idx] = newCircuit

			circuits = append(circuits, newCircuit)
			continue
		} else if aCache != nil && bCache != nil {
			if aCache != bCache {
				mergeCircuits(&connected, aCache, bCache)
				continue
			}

			// If both points are already in the circuit
			// we do not need to do anything anymore.
			continue
		}

		// First, assume point A has already been cached.
		// This means that we can add point B to the circuit
		// that point A is also in!
		oldCircuit := aCache
		newPoint := pd.B

		// If the assumtion is wrong, then this must mean
		// that point B has been cached, so we need to
		// add point A to the circuit of point B!
		if aCache == nil {
			oldCircuit = bCache
			newPoint = pd.A
		}

		connected[newPoint.Idx] = oldCircuit
		oldCircuit.Boxes = append(oldCircuit.Boxes, newPoint)

	}

	return circuits
}
