package main

import (
	"fmt"
)

// search returns the index of the element e in the
// sorted slice s. If e does not exist in s search returns -1
func search(s []int, e int) int {
	left := 0
	right := len(s)-1

	for (left <= right) {
		middle := left + (right-left)/2

		if e == s[middle] {
			return middle
		} else if e < s[middle] {
			right = middle-1
		} else {
			left = middle+1
		}
	}
	return -1
}

func main() {
	s := []int{0, 1, 2, 3, 4}
	fmt.Println(s)
	fmt.Printf("Index of 3 is: %v\n", search(s,3))
	fmt.Printf("Index of 10 is: %v\n", search(s,10))
}
