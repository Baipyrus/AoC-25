package day08

import "fmt"

// This type is meant to represent a pair of two points
// all the while providing a cached valued of their
// distance between each other. Because of the similarities,
// we interpret this as a rectangle made up of these points.
type Rectangle struct {
	A        Point
	B        Point
	Diagonal float64
	Area     uint64
}

func Area(a, b Point) uint64 {
	maxX, maxY := max(a.X, b.X), max(a.Y, b.Y)
	minX, minY := min(a.X, b.X), min(a.Y, b.Y)
	return (maxX - minX + 1) * (maxY - minY + 1)
}

func (rect Rectangle) String() string {
	return fmt.Sprintf("Point A: %s\n   -   = Dist.: %.2f\nPoint B: %s", rect.A, rect.Diagonal, rect.B)
}

func Combinations(boxes []Point) (closest []Rectangle) {
	// Scan all *unique* combinations of boxes to create
	// a list of two-point-distance combinations.
	scannedBoxes := make(map[KeyPair]bool)

	for i := range boxes {
		current := boxes[i]

		for j := range boxes {
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
			area := Area(current, next)

			closest = append(
				closest,
				Rectangle{
					A:        current,
					B:        next,
					Diagonal: dist,
					Area:     area})
		}
	}

	return closest
}
