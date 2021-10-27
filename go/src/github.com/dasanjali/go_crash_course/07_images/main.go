/* Remember the picture generator you wrote earlier? Let's write another one, but this time it will return an implementation of image.Image instead of a slice of data  */
package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	a int
	b int
}

func (img Image) At(a, b int) color.Color {
	v := uint8(a * b)
	return color.RGBA{v, v, 91, 255}
}
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.a, img.b)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}
func main() {
	picture := Image{91, 255}
	pic.ShowImage(picture)
}
