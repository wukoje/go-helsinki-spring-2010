package main

import "fmt"

// Textbook implementation of in-place unstable quicksort, based on
// the version presented at <http://en.wikipedia.org/wiki/Quicksort>.
func sort(s []int) {
	sort2(s, 0, len(s)-1)
}

func partition(s []int, left int, right int, pivotIndex int) int {
	pivotVal := s[pivotIndex]
	s[pivotIndex], s[right] = s[right], s[pivotIndex]
	storeIndex := left
	for i := left; i < right; i++ {
		if s[i] < pivotVal {
			s[i], s[storeIndex] = s[storeIndex], s[i]
			storeIndex++
		}
	}
	s[storeIndex], s[right] = s[right], s[storeIndex]
	return storeIndex
}

func sort2(s []int, left int, right int) {
	if right > left {
		i := (left + right) / 2
		j := partition(s, left, right, i)
		sort2(s, left, j-1)
		sort2(s, j+1, right)
	}
}

func main() {
	a := []int{7, 1, 4, 2, 42, 3, -123, 0, 4, 5, 4, 5}
	fmt.Println("Before:\n", a)
	sort(a)
	fmt.Println("After:\n", a)
}
