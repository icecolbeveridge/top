package top

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Null struct {
	Name string
	Point
	Jac   Mat3
	Evals [3]float64
	Evecs [3]Vector
}

func (n Null) String() string {
	return fmt.Sprintf("<NULL %s %v>", n.Name, n.Point)
}

func NullFromString(line string) (Null, error) {
	var err, e1 error
	fields := strings.Fields(line)
	if len(fields) < 5 {
		return Null{}, ErrInvalidString
	}
	n := Null{}
	n.Name = fields[4]
	n.X, e1 = strconv.ParseFloat(fields[1], 64)

	n.Y, err = strconv.ParseFloat(fields[2], 64)
	e1 = errors.Join(e1, err)
	n.Z, err = strconv.ParseFloat(fields[3], 64)
	e1 = errors.Join(e1, err)

	return n, e1
}
