// Package arithmetics provides basic arithmetic operations.
// It is designed to be simple, clear, and easy to use in Go projects.

package arithmetics

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add returns the sum of two numbers.
//
// An addend is a number that is added to another number.
// The sum is the result of adding the two addends together.
// Works with integers and floating-point numbers.
// For more information on addition, see: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](left, right T) T {
	return left + right
}
