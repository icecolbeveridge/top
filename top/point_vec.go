package top

import "fmt"

// A Vector is a 3D vector.
type Vector struct {
	X, Y, Z float64
}

// Zero is the Vector (0,0,0)
var Zero = Vector{0, 0, 0}

// String prints the details of a Vector
func (v Vector) String() string {
	return fmt.Sprintf("[%6.2f, %6.2f, %6.2f]", v.X, v.Y, v.Z)
}

// Mag2 returns the square of the magnitude of v.
func (v Vector) Mag2() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// ScalarMult multiplies v by a scalar.
func (v Vector) ScalarMult(t float64) Vector {
	return Vector{X: v.X * t, Y: v.Y * t, Z: v.Z * t}
}

// Add adds two Vectors together, giving a Vector
func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

// A Point is a point in 3D space
type Point struct {
	X, Y, Z float64
}

// Origin is the point (0,0,0)
var Origin = Point{0, 0, 0}

// String prints the details of a Point
func (p Point) String() string {
	return fmt.Sprintf("(%6.2f, %6.2f, %6.2f)", p.X, p.Y, p.Z)
}

// Sub subtracts p2 from the given point, returning a Vector.
func (p1 Point) Sub(p2 Point) Vector {
	return Vector{X: p1.X - p2.X, Y: p1.Y - p2.Y, Z: p1.Z - p2.Z}
}

// Add adds a Vector to a Point, giving a Point
func (p1 Point) Add(v2 Vector) Point {
	return Point{p1.X + v2.X, p1.Y + v2.Y, p1.Z + v2.Z}
}
