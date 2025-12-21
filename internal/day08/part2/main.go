package day08_part2

import (
	"fmt"
	"slices"

	"github.com/Baipyrus/AoC-25/internal/day08"
	"github.com/Baipyrus/AoC-25/internal/registry"
)

var name = "Day 08 - Part 2"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	boxes := day08.ParseInput(input)
	closest := day08.Combinations(boxes)

	// Sort all combinations such that we can start with
	// the two points that are closest to each other!
	slices.SortFunc(
		closest,
		func(a, b day08.PointDistance) int {
			return int(a.Dist - b.Dist)
		})

	lastPair := connectPoints(closest)

	product := lastPair.A.X * lastPair.B.X
	fmt.Printf("The product of multiplying the x-coordinates of the last connected pair is: %d\n", product)
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

func connectPoints(points []day08.PointDistance) (previousPair day08.PointDistance) {
	var circuits []*day08.Circuit

	// We want to remember which box is part
	// of which circuit without having to search!
	connected := make(map[int]*day08.Circuit)

	for _, pd := range points {
		aCache := connected[pd.A.Idx]
		bCache := connected[pd.B.Idx]

		// If both points are already in the circuit
		// we do not need to do anything anymore.
		if aCache != nil && aCache == bCache {
			continue
		}

		// Save the last used pair of points for output
		previousPair = pd

		// If none of the points have been connected yet,
		// create a new circuit and cache their connection!
		if aCache == nil && bCache == nil {
			circuitIdx := len(circuits)
			newCircuit := &day08.Circuit{Boxes: []day08.Point{pd.A, pd.B}, Idx: circuitIdx}

			connected[pd.A.Idx] = newCircuit
			connected[pd.B.Idx] = newCircuit

			circuits = append(circuits, newCircuit)
			continue
		} else if aCache != nil && bCache != nil && aCache != bCache {
			mergeCircuits(&connected, aCache, bCache)
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

	return previousPair
}
