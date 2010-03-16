/**

1 slice_reversal.go

// reverse reverses the contents of s in place

func reverse(s []int)


Output some slices before and after reversal.

*/

package main

import (
	"fmt"
)

func reverse(s []int) {
	for i := 0; i < len(s)/2; i++ {
		a := s[i]
		s[i] = s[(len(s)-1)-i]
		s[(len(s)-1)-i] = a
	}
}

func main() {

	// two lists for testing
	t1 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	t2 := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}

	fmt.Printf("t1 before: %v\n", t1)
	reverse(t1)
	fmt.Printf("t1 after:  %v\n", t1)

	fmt.Printf("t2 before: %v\n", t2)
	reverse(t2)
	fmt.Printf("t2 after:  %v\n", t2)
}
