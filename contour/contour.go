package contour

import "top/top"

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
