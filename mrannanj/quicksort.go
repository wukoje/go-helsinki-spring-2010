package main

import (
	"fmt"
)

// Sort slice s from left to right
func quicksort(s []int, left, right int) {
	if left < right {
		q := partition(s, left, right)
		quicksort(s, left, q-1)
		quicksort(s, q+1, right)
	}
}

func partition(s []int, left, right int) int {
	pivot := s[right]
	i := left
	for j := left; j < right; j++ {
		if s[j] < pivot {
			s[i], s[j] = s[j], s[i]
			i++
		}
	}
	s[i], s[right] = s[right], s[i]
	return i
}

func sort(s []int) { quicksort(s, 0, len(s)-1) }

func main() {
	s := []int{22, 5, 324, 23, 42, 34, 234, 24, 2, 10, 4, 33, 6, 7, 22, 22}
	fmt.Println(s)
	sort(s)
	fmt.Println(s)
}
