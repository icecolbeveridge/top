package top_test

import (
	"math"
	"strconv"
	"testing"
	"top/top"

	"github.com/stretchr/testify/assert"
)

func TestNull(t *testing.T) {
	assert := assert.New(t)
	n, err := top.NullFromString("0 1 2 3 A0")
	assert.NoError(err)

	assert.Equal(top.Point{1, 2, 3}, n.Point)
	assert.Equal("<NULL A0 (  1.00,   2.00,   3.00)>", n.String())

	n, err = top.NullFromString("bingybongyboo")
	assert.ErrorIs(top.ErrInvalidString, err)

	n, err = top.NullFromString("x y z a b c")
	assert.ErrorIs(err, strconv.ErrSyntax)

}

func TestPole(t *testing.T) {
	assert := assert.New(t)
	p, err := top.PoleFromString("P1 1 2 0 3")
	assert.NoError(err)
	assert.Equal(top.Point{1, 2, 0}, p.Point)
	assert.Equal(3., p.Strength)

	B := p.Field(top.Origin)
	t.Log(B)
	v := top.Vector{-1, -2, 0}
	v = v.ScalarMult(3 * math.Pow(5, -1.5))
	assert.Equal(v, B)

}
