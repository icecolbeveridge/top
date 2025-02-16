package top

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

type Pole struct {
	Name string
	Point
	Strength float64
}

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

func (p Pole) Field(pt Point) Vector {
	dx := pt.Sub(p.Point)
	m := p.Strength * math.Pow(dx.Mag2(), -1.5)
	return dx.ScalarMult(m)
}
