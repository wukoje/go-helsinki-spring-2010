package main

import (
	"fmt"
)

// search returns the index of element e in a sorted slice s. 
// If s doesn't contain e, search returns -1.
func search(s []int, e int) int {

	low := 0
	high := len(s) - 1
	var mid int

	for low <= high {
		mid = (low + high) / 2
		if s[mid] == e {
			return mid
		} else if s[mid] > e {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func main() {
	s := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println(s)
	fmt.Printf("1: %v, 4: %v, 6: %v, -1: %v, 12: %v\n", search(&s, 1), search(&s, 4), search(&s, 6), search(&s, -1), search(&s, 12))
}
