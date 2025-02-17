package top

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

// A Pole is a magnetic monopole (don't @ me) with a name, position and strength.
type Pole struct {
	Name string
	Point
	Strength float64
}

// PoleFromString creates a Pole from a line of a .pns file, which must be in
// the form [label, as string with no space] [x, as float] [y, as float] [z, as float], [phi, as float]
// Anything after phi is ignored, although it should strictly fail if it's not a comment.
func PoleFromString(line string) (Pole, error) {
	// always label, x, y, z, phi
	var err, e1 error
	fields := strings.Fields(line)
	p := Pole{}
	p.Name = fields[0]
	p.X, e1 = strconv.ParseFloat(fields[1], 64)

	p.Y, err = strconv.ParseFloat(fields[2], 64)
	e1 = errors.Join(e1, err)
	p.Z, err = strconv.ParseFloat(fields[3], 64)
	e1 = errors.Join(e1, err)
	p.Strength, err = strconv.ParseFloat(fields[4], 64)
	e1 = errors.Join(e1, err)
	return p, e1
}

// Field returns the magnetic field at a point due to a point source, given as
// phi (pt - p)/ |pt - p|^3
func (p Pole) Field(pt Point) Vector {
	dx := pt.Sub(p.Point)
	m := p.Strength * math.Pow(dx.Mag2(), -1.5)
	return dx.ScalarMult(m)
}

// Potential returns the magnetic potential at a point due to a point source, given as
// phi / |pt - p|
func (p Pole) Potential(pt Point) float64 {
	dx := pt.Sub(p.Point)
	return p.Strength * math.Pow(dx.Mag2(), -0.5)
}
