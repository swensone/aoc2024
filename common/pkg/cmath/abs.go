package cmath

import (
	"golang.org/x/exp/constraints"
)

// Numeric is a type that describes an integer or a float number
type Numeric interface {
	constraints.Integer | constraints.Float
}

// Abs returns the absolute value of a number
func Abs[T Numeric](n T) T {
	if n < 0 {
		return -n
	}
	return n
}
