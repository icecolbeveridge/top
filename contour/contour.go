package contour

import (
	"io"
	"math"
	"top/top"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// Autogrid gives the potential on a 21x21 grid picked in a
// fairly arbitrary way. TODO: make it customisable
func Autogrid(t *top.Topology) [][]float64 {
	const (
		NX = 401
		NY = 401
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
	ax := (x1 - x0) * 0.5
	x0 = x0 - ax
	x1 = x1 + ax

	ay := (y1 - y0) * 0.5
	y0 = y0 - ay
	y1 = y1 + ay

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
	Edges []BrokenEdge
	Start MIDPOINT
	End   MIDPOINT
}

func Adjacent(m1, m2 MIDPOINT) bool {
	dx := (m1[0] - m2[0])
	dy := (m1[1] - m2[1])
	switch {
	case dx == 0 && dy*dy == 1:
		return true
	case dy == 0 && dx*dx == 1:
		return true
	case dx*dy == 1:
		return true
	default:
		return false
	}
}

// A BrokenEdge is a pair of gridpoints on either side of a contour line
// and an estimate of where the contour ought to be.
type BrokenEdge struct {
	First    GridPoint
	Second   GridPoint
	EstBreak GridPoint
}

// BreakEdge determines whether the ends of an edge are on different
// sides of a contour, and returns an estimate of the crossing point
// (or (0, false) if it's not broken)
func BreakEdge(p1, p2, level float64) (float64, bool) {
	if p1 > level && p2 > level {
		return 0, false
	}
	if p1 < level && p2 < level {
		return 0, false
	}
	L1 := math.Abs(p1 - level)
	L2 := math.Abs(p2 - level)
	return L1 / (L1 + L2), true
}

// MIDPOINT is (2x, 2y) for the midpoint of an edge -- using the
// doubles keeps everything nice and integer.
type MIDPOINT = [2]int

// Contours returns the curves for a list of levels
func Contours(grid [][]float64, levels ...float64) []Curve {
	out := make([]Curve, 0)
	for _, l := range levels {
		out = append(out, Contour(grid, l)...)
	}
	return out
}

// Contour returns contours for a single level
func Contour(grid [][]float64, level float64) []Curve {
	NX := len(grid)
	NY := len(grid[0])
	brokenEdges := make(map[MIDPOINT]BrokenEdge)
	var l float64
	var isBroken bool
	for x := range NX {
		for y := range NY {
			if x < NX-1 { // check horizontal
				if l, isBroken = BreakEdge(grid[x][y], grid[x+1][y], level); isBroken {
					fx := float64(x)
					fy := float64(y)
					bEdge := BrokenEdge{
						First:    GridPoint{X: fx, Y: fy, Potential: grid[x][y]},
						Second:   GridPoint{X: fx + 1, Y: fy, Potential: grid[x+1][y]},
						EstBreak: GridPoint{X: fx + l, Y: fy},
					}
					a := [...]int{2*x + 1, 2 * y}
					brokenEdges[a] = bEdge
				}
			}
			if x < NX-1 && y < NY-1 { // check diagonal
				if l, isBroken = BreakEdge(grid[x][y], grid[x+1][y+1], level); isBroken {
					fx := float64(x)
					fy := float64(y)
					bEdge := BrokenEdge{
						First:    GridPoint{X: fx, Y: fy, Potential: grid[x][y]},
						Second:   GridPoint{X: fx + 1, Y: fy + 1, Potential: grid[x+1][y+1]},
						EstBreak: GridPoint{X: fx + l, Y: fy + l},
					}
					a := [...]int{2*x + 1, 2*y + 1}
					brokenEdges[a] = bEdge
				}
			}
			if y < NY-1 { // check vertical
				if l, isBroken = BreakEdge(grid[x][y], grid[x][y+1], level); isBroken {
					fx := float64(x)
					fy := float64(y)
					bEdge := BrokenEdge{
						First:    GridPoint{X: fx, Y: fy, Potential: grid[x][y]},
						Second:   GridPoint{X: fx, Y: fy + 1, Potential: grid[x][y+1]},
						EstBreak: GridPoint{X: fx, Y: fy + l},
					}
					a := [...]int{2 * x, 2*y + 1}
					brokenEdges[a] = bEdge
				}
			}
		}
	}
	return CombineBrokenEdges(brokenEdges)
}

// CombineBrokenEdges takes a map of BrokenEdges and combines them together into disjoint Curves.
func CombineBrokenEdges(m map[MIDPOINT]BrokenEdge) []Curve {
	out := make([]Curve, 0)
	for {
		c := Curve{
			Edges: []BrokenEdge{},
			Start: [2]int{},
			End:   [2]int{},
		}
		for {
			tr := make([]MIDPOINT, 0)
			for mp, be := range m {
				switch {
				case len(c.Edges) == 0:
					c.Start = mp
					c.End = mp
					c.Edges = []BrokenEdge{be}
					tr = append(tr, mp)
				case Adjacent(mp, c.Start):
					c.Start = mp
					c.Edges = append([]BrokenEdge{be}, c.Edges...)
					tr = append(tr, mp)
				case Adjacent(mp, c.End):
					c.End = mp
					c.Edges = append(c.Edges, be)
					tr = append(tr, mp)
				default:
				}
			}
			for _, t := range tr {
				delete(m, t)
			}
			if len(tr) == 0 {
				break
			}
		}
		out = append(out, c)
		if len(m) == 0 {
			break
		}
	}
	return out
}

// Plot takes a set of curves and writes the output to a Writer. TODO: make customisable
func Plot(curves []Curve, w io.Writer) {
	p := plot.New()
	for _, c := range curves {
		pp := plotter.XYs{}
		for _, cp := range c.Edges {
			pp = append(pp, plotter.XY{X: cp.EstBreak.X, Y: cp.EstBreak.Y})
		}
		if Adjacent(c.Start, c.End) {
			pp = append(pp, plotter.XY{X: c.Edges[0].EstBreak.X, Y: c.Edges[0].EstBreak.Y})
		}
		line, err := plotter.NewLine(pp)
		if err != nil {
			panic(err)
		}
		p.Add(line)
	}
	wt, err := p.WriterTo(vg.Length(800), vg.Length(600), "png")
	if err != nil {
		panic(err)
	}
	_, err = wt.WriteTo(w)
	if err != nil {
		panic(err)
	}
}
