package day08

import "fmt"

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

func Combinations(boxes []Point) (closest []PointDistance) {
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

			closest = append(
				closest,
				PointDistance{
					A:    current,
					B:    next,
					Dist: dist})
		}
	}

	return closest
}
