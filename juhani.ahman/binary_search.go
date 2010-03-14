package main

import (
	"fmt"
)

func search(s []int, e int) int {
	low := 0
	high := len(s) - 1

	for low < high {
		mid := low + ((high - low) / 2)
		if s[mid] < e {
			low = mid + 1
		} else {
			high = mid
		}
	}

	if (low < len(s)) && (s[low] == e) {
		return low // found
	}

	return -1 // not found
}

func main() {
	s := make([]int, 11)
	for i, _ := range s {
		s[i] = i * 2
	}
	fmt.Printf("Slice: %v:\n", s)
	for i := 0; i <= 20; i++ {
		result := search(s, i)
		fmt.Printf("Searching for %d. Got %d.\n", i, result)
	}
}
