package day08_part1

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/Baipyrus/AoC-25/internal/registry"
)

var name = "Day 08 - Part 1"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	var (
		boxes       []Point
		closest     []PointDistance
		connections int = 1000
	)

	// Parse raw input into list of junction boxes
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		bline := strings.TrimSpace(line)
		if bline == "" {
			continue
		}

		coordinates := strings.Split(bline, ",")
		x, _ := strconv.ParseUint(coordinates[0], 10, 64)
		y, _ := strconv.ParseUint(coordinates[1], 10, 64)
		z, _ := strconv.ParseUint(coordinates[2], 10, 64)

		next := Point{X: x, Y: y, Z: z, Idx: i}
		next.Mag = next.Magnitude()
		boxes = append(boxes, next)
	}

	// Scan all *unique* combinations of boxes to create
	// a list of two-point-distance combinations.
	scannedBoxes := make(map[KeyPair]bool)
	for i := 0; i < len(boxes); i++ {
		current := boxes[i]

		for j := 0; j < len(boxes); j++ {
			// We want to ignore previously scanned
			// pairs where (i, j) = (j, i).
			keyPair := NewKeyPair(i, j)
			if i == j || scannedBoxes[keyPair] {
				continue
			}
			scannedBoxes[keyPair] = true

			// Calculate and save closest point pair
			next := boxes[j]
			dist := Distance(current, next)

			closest = append(
				closest,
				PointDistance{
					A:    current,
					B:    next,
					Dist: dist})
		}
	}

	// Sort all combinations such that we can start with
	// the two points that are closest to each other!
	slices.SortFunc(
		closest,
		func(a, b PointDistance) int {
			return int(a.Dist - b.Dist)
		})

	// NOTE: IMPORTANT! We only want to make n connections!
	circuits := connectPoints(closest[:connections])

	// For our final solution, we want to multiply the three
	// largest circuit's lengths with each other:
	slices.SortFunc(
		circuits,
		func(a, b *Circuit) int {
			return len(b.Boxes) - len(a.Boxes)
		})

	product := len(circuits[0].Boxes) * len(circuits[1].Boxes) * len(circuits[2].Boxes)
	fmt.Printf("The product of multiplying the length of the three largest circuits is: %d\n", product)
}

func mergeCircuits(cache *map[int]*Circuit, a, b *Circuit) {
	a.Boxes = append(a.Boxes, b.Boxes...)

	// Update cache references
	for _, box := range b.Boxes {
		(*cache)[box.Idx] = a
	}

	// "Delete" b by resetting its boxes
	b.Boxes = []Point{}
}

func connectPoints(points []PointDistance) (circuits []*Circuit) {
	// We want to remember which box is part
	// of which circuit without having to search!
	connected := make(map[int]*Circuit)

	for _, pd := range points {
		aCache := connected[pd.A.Idx]
		bCache := connected[pd.B.Idx]

		// If none of the points have been connected yet,
		// create a new circuit and cache their connection!
		if aCache == nil && bCache == nil {
			circuitIdx := len(circuits)
			newCircuit := &Circuit{Boxes: []Point{pd.A, pd.B}, Idx: circuitIdx}

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

type Circuit struct {
	Boxes []Point
	Idx   int
}

// This type is meant to be a composite key for a
// hashmap where the sub-keys are meant to be sorted
// such that the order of their combination does not
// matter. Because of this, it is recommended to NEVER
// create a custom instance of this, but to use the
// provided `NewKeyPair` function instead.
type KeyPair struct {
	Max int
	Min int
}

// Takes any integers (a, b) and creates a new
// KeyPair instance, regardless of the order of
// either of the two integer values.
func NewKeyPair(a, b int) KeyPair {
	// Assume the keys are in order
	if a > b {
		return KeyPair{Max: a, Min: b}
	}

	// And if not, switch them
	return KeyPair{Max: b, Min: a}
}

// This type is meant to represent a pair of two points
// all the while providing a cached valued of their
// distance between each other. Additionally, we allow
// its previous index in any array to be stored inside it.
type PointDistance struct {
	A    Point
	B    Point
	Dist float64
}

func (pd PointDistance) String() string {
	return fmt.Sprintf("Point A: %s\n   -   = Dist.: %.2f\nPoint B: %s", pd.A, pd.Dist, pd.B)
}

// This type represents a single point in 3D space.
// Additionally, it also provides a cached value of
// it's magnitude if viewed as a vector and we allow
// the index of this point in some outside array to
// be stored as part of the Point for easy access.
type Point struct {
	X   uint64
	Y   uint64
	Z   uint64
	Mag float64
	Idx int
}

// Calculate magnitude simply as distance from origin
func (p *Point) Magnitude() float64 {
	return Distance(*p, Point{X: 0, Y: 0, Z: 0})
}

func (p Point) String() string {
	return fmt.Sprintf("XYZ: %d,%d,%d (Mag.: %.2f)", p.X, p.Y, p.Z, p.Mag)
}

// Simple 3D Euclidian distance function
func Distance(a, b Point) float64 {
	dx, dy, dz := a.X-b.X, a.Y-b.Y, a.Z-b.Z
	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}
