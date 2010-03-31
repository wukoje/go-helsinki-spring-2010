package main

import (
	"fmt"
)

func main() {

	s := []int{2, 5, 1, 6, 3, 9, 3, 10}
	s = sort(s)
	fmt.Println(s)
	s = []int{3, 7, 8, 5, 2, 1, 9, 5, 4}
	s = sort(s)
	fmt.Println(s)

}

// sort modifies the slice s so that the integers are sorted in
// place using quicksort
func sort(s []int) (sorted []int) {

	if len(s) <= 1 {
		return s
	}

	pivotIndex := len(s) - 1
	pivot := s[pivotIndex]

	for i, e := range s {
		if i >= pivotIndex {
			break
		} else if e <= pivot {
			continue
		} else {
			for s[i] > pivot {
				s[i], s[pivotIndex-1] = s[pivotIndex-1], s[i]
				if pivot != s[pivotIndex-1] {
					s[pivotIndex], s[pivotIndex-1] = s[pivotIndex-1], s[pivotIndex]
				}
				pivotIndex--
			}
		}
	}

	var less, greater []int

	if pivotIndex == 0 {
		greater = s[pivotIndex+1 : len(s)]
	} else if pivotIndex == len(s)-1 {
		less = s[0:pivotIndex]
	} else {
		less = s[0:pivotIndex]
		greater = s[pivotIndex+1 : len(s)]
	}

	sortedLess := sort(less)
	sortedGreater := sort(greater)

	sorted = make([]int, len(s))
	copy(sorted, sortedLess)
	sorted[len(sortedLess)] = pivot
	copy(sorted[len(sortedLess)+1:cap(sorted)], sortedGreater)

	return sorted
}
