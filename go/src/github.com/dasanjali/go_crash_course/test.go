package main

import "fmt"

func main() {
	country_array := []string{"Australia", "France", "Germany", "India", "Ireland"}

	fmt.Println("Country list:", country_array)

	country_slicing := country_array[2:5]
	fmt.Println("After slicing:", country_slicing)
}
