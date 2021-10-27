/* Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program,
it will display your picture, interpreting the integers as grayscale (well, bluescale) values. */

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for y := range image {
		for x := 0; x < dx; x++ {
			image[y] = append(image[y], uint8(x*y))
		}
	}
	return image
}
func main() {
	pic.Show(Pic)
}
