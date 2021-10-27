/* Copy your Sqrt function from the earlier exercise and modify it to return an error value.

Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers. */

package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (err ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Sorry! This is a negative number: %f", float64(err))
}

func Sqrt(num float64) (float64, error) {
	if num > 0 {
		return math.Sqrt(num), nil
	} else {
		return 0, ErrNegativeSqrt(num)
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
