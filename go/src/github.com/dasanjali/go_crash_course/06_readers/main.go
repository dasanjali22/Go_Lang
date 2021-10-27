/* Implement a Reader type that emits an infinite stream of the ASCII character 'A'. */

package main

import "golang.org/x/tour/reader"

type MyReader struct{}
type MyReaderError bool

func (MyReaderError) Error() string {
	return "Capacity is insufficient"
}

func (MyReader) Read(b []byte) (int, error) {
	if cap(b) < 1 {
		return 0, MyReaderError(true)
	}
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
