/*
 * slice_reversal.go: In-place reversal of a slice.
 *
 * For Introduction to Go, Spring 2010
 * Kimmo Kulovesi <kkuloves@cs.helsinki.fi>
 */
package main

import (
	"fmt"
)

// Reverse the contents of the slice s in place.
func reverse(s []int) {
	for head, tail := 0, len(s)-1; head < tail; head, tail = head+1, tail-1 {
		s[tail], s[head] = s[head], s[tail]
	}
}

// Test reversing by repeatedly slicing and reversing parts of an array
func main() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for length := 0; length <= len(data); length++ {
		for start := 0; start < len(data)-length; start++ {
			end := start + length
			s := data[start:end]

			fmt.Printf("Slice [%d:%d]:\t%v\nReversed:\t", start, end, s)
			reverse(s)
			fmt.Println(s)

			// One test case suffices for length < 2
			if length < 2 {
				break
			}
		}
	}
}
