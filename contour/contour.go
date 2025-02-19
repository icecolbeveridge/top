package contour

import (
	"math"
	"top/top"
)

// Autogrid gives the potential on a 21x21 grid picked in a
// fairly arbitrary way. TODO: make it customisable
func Autogrid(t *top.Topology) [][]float64 {
	const (
		NX = 21
		NY = 21
	)
	x0 := 1.e30
	y0 := 1.e30
	x1 := -1.e30
	y1 := -1.e30

	for _, p := range t.Poles {
		if p.X < x0 {
			x0 = p.X
		} else if p.X > x1 {
			x1 = p.X
		}
		if p.Y < y0 {
			y0 = p.Y
		} else if p.Y > y1 {
			y1 = p.Y
		}
	}
	dx := (x1 - x0) / (NX - 1)
	dy := (y1 - y0) / (NY - 1)

	out := make([][]float64, NX)
	y := y0
	for i := range NY {
		y += dy
		x := x0
		out[i] = make([]float64, NY)
		for j := range NX {
			x += dx
			p := top.Point{X: x, Y: y, Z: 0.}
			out[i][j] = t.Potential(p)
		}
	}
	return out
}

// A GridPoint is a point in grid coordinates -- TODO: a conversion between
// grid coords and "real" coords
type GridPoint struct {
	X, Y      float64
	Potential float64
}

// A Curve is a list of GridPoints that can be closed or not.
type Curve struct {
	Points []GridPoint
	Closed bool
}

// A BrokenEdge is a pair of gridpoints on either side of a contour line
// and an estimate of where the contour ought to be.
type BrokenEdge struct {
	First    GridPoint
	Second   GridPoint
	EstBreak GridPoint
}

// Let's think this through. I certainly need to figure out which
// edges are broken, and then decide which pairs are adjacent.
//
// What's the output? Must be a list of curves.
func Contour(grid [][]float64, level float64) []Curve {
	NX := len(grid)
	NY := len(grid[0])
	out := make([]Curve, 0)
	bhEdges := make([]BrokenEdge, 0)
	// bvEdges := make([]BrokenEdge, 0)
	// bdEdges := make([]BrokenEdge, 0)
	for x := range NX {
		for y := range NY {
			if x < NX-1 { // check horizontal
				if (grid[x][y]-level)*(grid[x+1][y]-level) < 0. {
					fx := float64(x)
					fy := float64(y)
					// if p1 = -a and p2 = b, then the zero estimate
					// is (|b|x2 + |a|x1)/(|a| + |b|)
					p1 := math.Abs(grid[x][y])
					p2 := math.Abs(grid[x+1][y])
					ex := (p1*(fx+1) + p2*fx) / (p1 + p2)
					bEdge := BrokenEdge{
						First:    GridPoint{X: fx, Y: fy, Potential: grid[x][y]},
						Second:   GridPoint{X: fx + 1, Y: fy, Potential: grid[x+1][y]},
						EstBreak: GridPoint{X: ex, Y: fy},
					}
					bhEdges = append(bhEdges, bEdge)
				}
			}
		}
	}
	return out
}
