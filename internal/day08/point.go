package day08

import (
	"fmt"
	"math"
)

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
