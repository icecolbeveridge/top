package top

import "fmt"

type Point struct {
	X, Y, Z float64
}

var Origin = Point{0, 0, 0}
var Zero = Vector{0, 0, 0}

func (p Point) String() string {
	return fmt.Sprintf("(%6.2f, %6.2f, %6.2f)", p.X, p.Y, p.Z)
}

func (p1 Point) Sub(p2 Point) Vector {
	return Vector{X: p1.X - p2.X, Y: p1.Y - p2.Y, Z: p1.Z - p2.Z}
}

func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

type Vector struct {
	X, Y, Z float64
}

func (v Vector) Mag2() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vector) ScalarMult(t float64) Vector {
	return Vector{X: v.X * t, Y: v.Y * t, Z: v.Z * t}
}
