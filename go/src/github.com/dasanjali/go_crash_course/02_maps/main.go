/* Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.  */
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		_, pass := m[word]
		if pass {
			m[word] += 1
		} else {
			m[word] = 1
		}
	}
	return m
}
func main() {
	wc.Test(WordCount)
}
