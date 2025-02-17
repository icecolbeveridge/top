package top

import (
	"fmt"
)

var ErrInvalidString error

// A Topology is a list of Poles, Nulls and Separators, usually representing a region of the
// solar photosphere. It is not a great physical model since (a) the solar magnetic field is not
// potential, (b) magnetic monopoles don't exist and (c) the Sun is not flat (Beveridge, 2003).
type Topology struct {
	Poles      map[string]Pole
	Nulls      map[string]Null
	Separators map[string]Separator
}

// String represents the topology as a string, listing the poles and nulls. (Once I have a good idea how to show the separators,
// I'll also do that. TODO)
func (t Topology) String() string {
	out := ""
	if len(t.Poles) > 0 {
		out += "POLES\n"
		for k, v := range t.Poles {
			out += fmt.Sprintf("%s: %v\n", k, v)
		}
	}
	if len(t.Nulls) > 0 {
		out += "NULLS\n"
		for k, v := range t.Nulls {
			out += fmt.Sprintf("%s: %v\n", k, v)
		}
	}
	return out
}

// NewTopology initialises the fields of a Topology.
func NewTopology() *Topology {
	t := new(Topology)
	t.Poles = make(map[string]Pole)
	t.Nulls = make(map[string]Null)
	t.Separators = make(map[string]Separator)
	return t
}

// Mat3 is a 3x3 matrix (and doesn't belong here TODO)
type Mat3 [3][3]float64

// A Separator is a list of points on a field line connecting two Nulls.
type Separator struct {
	PNull, NNull Null
	Points       []Point
}

// Field returns the magnetic field at a point, simply summing the field from each Pole.
func (T *Topology) Field(pt Point) Vector {
	out := Zero
	for _, p := range T.Poles {
		out = out.Add(p.Field(pt))
	}
	return out
}

// Potential returns the magnetic potential at a point, summing the potential field due to each Pole.
func (T *Topology) Potential(pt Point) float64 {
	out := 0.
	for _, p := range T.Poles {
		out += p.Potential(pt)
	}
	return out
}
