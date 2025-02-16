package top

import (
	"fmt"
)

var ErrInvalidString error

type Topology struct {
	Poles      map[string]Pole
	Nulls      map[string]Null
	Separators map[string]Separator
}

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

func NewTopology() *Topology {
	t := new(Topology)
	t.Poles = make(map[string]Pole)
	t.Nulls = make(map[string]Null)
	t.Separators = make(map[string]Separator)
	return t
}

type Mat3 [3][3]float64

type Separator struct {
	PNull, NNull Null
	Points       []Point
}

func (T *Topology) Field(pt Point) Vector {
	out := Zero
	for _, p := range T.Poles {
		out = out.Add(p.Field(pt))
	}
	return out
}
