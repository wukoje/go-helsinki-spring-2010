package main

import (
	"fmt"
)

// In place, sorts elements smaller than pivot to be befor it
// and elements greater than pivot to be after it.
// Returns the index of the pivot value.
func partition(s []int, pivotIndex int) int {
	pivot := s[pivotIndex]
	right := len(s) - 1
	s[right], s[pivotIndex] = s[pivotIndex], s[right]
	tmpIndex := 0
	for i := 0; i < right; i++ {
		if s[i] <= pivot {
			s[i], s[tmpIndex] = s[tmpIndex], s[i]
			tmpIndex++
		}
	}
	s[tmpIndex], s[right] = s[right], s[tmpIndex]
	return tmpIndex
}

// Partitions the slice around the element at pivotIndex,
// and then quicksorts the parts on the left and right hand side
// of the newPivotIndex.
func quicksort(s []int) {
	if len(s) > 1 {
		pivotIndex := int(len(s) / 2)
		newPivotIndex := partition(s, pivotIndex)
		quicksort(s[0:newPivotIndex])
		quicksort(s[newPivotIndex:len(s)])
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(s)
	quicksort(s)
	fmt.Println(s)
	s = []int{4, 2, 3, 8, 5, 13, 7, 2, 98, 10}
	fmt.Println(s)
	quicksort(s)
	fmt.Println(s)
}
