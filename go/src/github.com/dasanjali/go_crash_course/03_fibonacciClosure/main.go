/* Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...). */
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var first, second, x int

	return func() int {
		var temp int
		var fib_series int
		x += 1
		if x == 1 {
			fib_series = 0
			first = fib_series
		} else if x == 2 {
			fib_series = 1

		} else {
			fib_series = first + second

		}
		temp = first
		first = fib_series
		second = temp

		return fib_series
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 15; i++ {
		fmt.Println(f())
	}
}
