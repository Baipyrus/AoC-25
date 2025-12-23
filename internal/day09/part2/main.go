package day09_part2

import (
	"fmt"
	"slices"

	"github.com/Baipyrus/AoC-25/internal/day08"
	"github.com/Baipyrus/AoC-25/internal/registry"
)

var name = "Day 09 - Part 2"

func init() {
	registry.Register(name, Main)
}

func Main(input string) {
	fmt.Printf("Executing: %s\n", name)

	// Parse input vertices
	tiles := day08.ParseInput(input)

	var walls []Wall
	for idx, current := range tiles {
		next := tiles[(idx+1)%len(tiles)]

		walls = append(
			walls,
			Wall{
				Start: current,
				End:   next})
	}

	// Calculate all the different tile-combinations
	// that would form rectangles on the grid:
	combinations := day08.Combinations(tiles)

	// Sort all combinations such that we can start with
	// the two points that are furthest from each other!
	slices.SortFunc(
		combinations,
		func(a, b day08.Rectangle) int {
			return int(b.Area - a.Area)
		})

	var furthest day08.Rectangle
	for _, rect := range combinations {
		xRange := day08.NewKeyPair(rect.A.X, rect.B.X)
		yRange := day08.NewKeyPair(rect.A.Y, rect.B.Y)

		topLeft := day08.Point{X: xRange.Min, Y: yRange.Min}
		topRight := day08.Point{X: xRange.Max, Y: yRange.Min}
		bottomLeft := day08.Point{X: xRange.Min, Y: yRange.Max}
		bottomRight := day08.Point{X: xRange.Max, Y: yRange.Max}

		// Check if rectangle corners are inside shape
		tlInside := isInsideShape(walls, topLeft)
		trInside := isInsideShape(walls, topRight)
		blInside := isInsideShape(walls, bottomLeft)
		brInside := isInsideShape(walls, bottomRight)

		top := Wall{Start: topLeft, End: topRight}
		left := Wall{Start: topLeft, End: bottomLeft}
		right := Wall{Start: topRight, End: bottomRight}
		bottom := Wall{Start: bottomLeft, End: bottomRight}

		// Rectangle is NOT within shape if shape perimeter
		// intersects the rectangle edges at any point!
		tIntersect := intersectsShape(walls, top)
		lIntersect := intersectsShape(walls, left)
		rIntersect := intersectsShape(walls, right)
		bIntersect := intersectsShape(walls, bottom)

		isInside := tlInside && trInside && blInside && brInside
		notIntersected := !(tIntersect || lIntersect || rIntersect || bIntersect)
		if isInside && notIntersected {
			furthest = rect
			break
		}
	}

	fmt.Printf("The largest area of any of the available rectangles is: %d\n", furthest.Area)
}

func intersectsShape(s []Wall, w Wall) bool {
	for _, o := range s {
		if o.Intersect(w) {
			return true
		}
	}
	return false
}

func isInsideShape(s []Wall, p day08.Point) bool {
	var intersections uint

	for _, w := range s {
		// If the point is exactly on a wall, it counts as "inside".
		if w.Contains(p) {
			return true
		}

		// Only consider walls that span the Y-level of our point.
		// Horizontal walls can be completely ignored because they
		// don't change the outcome of the horizontal ray's collision.
		// Additonally, we use a "half-open" interval [Min, Max) to
		// handle the wall vertices correctly for collisions with
		// vertical wall segments.
		yRange := day08.NewKeyPair(w.Start.Y, w.End.Y)
		if p.Y >= yRange.Min && p.Y < yRange.Max {
			// Since we are casting a ray to the left (towards X=0),
			// we only care about walls to the left of our point.
			if w.Start.X == w.End.X && w.Start.X < p.X {
				intersections++
			}
		}
	}

	return intersections%2 == 1
}

type Wall struct {
	Start day08.Point
	End   day08.Point
}

func (w *Wall) Contains(p day08.Point) bool {
	xRange := day08.NewKeyPair(w.Start.X, w.End.X)
	yRange := day08.NewKeyPair(w.Start.Y, w.End.Y)

	// If wall is vertical, detect if p.Y is in range
	if w.Start.X == w.End.X && p.X == w.Start.X {
		return yRange.Min <= p.Y && p.Y <= yRange.Max
	}

	// If wall is horizontal, detect if p.X is in range
	if w.Start.Y == w.End.Y && p.Y == w.Start.Y {
		return xRange.Min <= p.X && p.X <= xRange.Max
	}

	return false
}

func (self *Wall) Intersect(other Wall) bool {
	// If not horizontal, it must be vertical
	selfHorizontal := self.Start.Y == self.End.Y
	otherHorizontal := other.Start.Y == other.End.Y

	// Perpendicular line intersections:
	// NOTE: We *specifically* only care about open
	//       interval intersections, meaning that the
	//       start and end coordinates do not count!
	if selfHorizontal && !otherHorizontal {
		return h_intersect_v(*self, other)
	}
	if !selfHorizontal && otherHorizontal {
		return h_intersect_v(other, *self)
	}

	// NOTE: We want to *specifically* ignore parallel
	//       line intersections (overlap) !!!
	//
	// if selfHorizontal && otherHorizontal && self.Start.Y == other.Start.Y {
	// 	return h_intersect_h(*self, other)
	// }
	// if !selfHorizontal && !otherHorizontal && self.Start.X == other.Start.X {
	// 	return v_intersect_v(*self, other)
	// }

	return false
}

// func v_intersect_v(a, b Wall) bool {
// 	// The start and end points have to be ordered for this to work
// 	y_1 := day08.NewKeyPair(a.Start.Y, a.End.Y)
// 	y_2 := day08.NewKeyPair(b.Start.Y, b.End.Y)
//
// 	// If the largest of the starting positions is larger or equal to
// 	// the smallest of the end positions, there must be some overlap.
// 	return max(y_1.Min, y_2.Min) <= min(y_1.Max, y_2.Max)
// }
//
// func h_intersect_h(a, b Wall) bool {
// 	// The start and end points have to be ordered for this to work
// 	x_1 := day08.NewKeyPair(a.Start.X, a.End.X)
// 	x_2 := day08.NewKeyPair(b.Start.X, b.End.X)
//
// 	// If the largest of the starting positions is larger or equal to
// 	// the smallest of the end positions, there must be some overlap.
// 	return max(x_1.Min, x_2.Min) <= min(x_1.Max, x_2.Max)
// }

func h_intersect_v(h, v Wall) bool {
	// NOTE: v.Start.X == v.End.X && h.Start.Y == h.End.Y
	// horizontal_x_start < vertical_x < horizontal_x_end
	//                          AND
	// vertical_y_start < horizontal_y < vertical_y_end

	hx := day08.NewKeyPair(h.Start.X, h.End.X)
	vy := day08.NewKeyPair(v.Start.Y, v.End.Y)

	return hx.Min < v.Start.X && v.End.X < hx.Max && vy.Min < h.Start.Y && h.End.Y < vy.Max
}
