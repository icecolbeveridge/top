package top_test

import (
	"testing"
	"top/top"

	"github.com/stretchr/testify/assert"
)

func TestTop(t *testing.T) {
	assert := assert.New(t)

	T := top.NewTopology()
	assert.NotNil(T)
	assert.Equal("", T.String())

	T.Poles["P1"] = top.Pole{
		Name:     "P1",
		Point:    top.Point{1, 2, 3},
		Strength: -3,
	}
	T.Nulls["A0"] = top.Null{
		Name:  "A0",
		Point: top.Point{0, 0, 0},
		Jac:   [3][3]float64{},
		Evals: [3]float64{},
		Evecs: [3]top.Vector{},
	}

	s := "POLES\nP1: (  1.00,   2.00,   3.00)\n"
	s += "NULLS\nA0: <NULL A0 (  0.00,   0.00,   0.00)>\n"
	assert.Equal(s, T.String())
}
