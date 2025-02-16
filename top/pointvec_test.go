package top_test

import (
	"testing"
	"top/top"

	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
	assert := assert.New(t)

	P := top.Point{X: 1, Y: 2, Z: 3}
	assert.Equal("(  1.00,   2.00,   3.00)", P.String())

	Z := top.Vector{X: 0, Y: 0, Z: 0}
	assert.Equal(P.Sub(P), Z)

	Q := top.Point{X: 2, Y: 3, Z: 4}
	d := Q.Sub(P)
	u := top.Vector{X: 1, Y: 1, Z: 1}
	assert.Equal(u, d)
	assert.Equal(3., d.Mag2())

	v := u.ScalarMult(2)
	assert.Equal(top.Vector{2, 2, 2}, v)
}
