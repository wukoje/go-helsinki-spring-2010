package main

import (
	"fmt"
	"rand"
	"time"
	"os"
	sortpkg "sort"
)

func main() {
	const slicesize = 100000

	s := make([]int, slicesize)

	// Initialize s with random values.
	for i := 0; i < len(s); i++ {
		s[i] = rand.Int()
	}

	// Verify disorder.
	if sortpkg.IntsAreSorted(s) {
		fmt.Println("Elements already in order, aborting.")
		os.Exit(1)
	}

	// Perform the sort and time it.
	startTime := time.Nanoseconds()
	sort(s)
	endTime := time.Nanoseconds()

	// Check that s is in fact sorted.
	if sortpkg.IntsAreSorted(s) {
		fmt.Println("Sorted", slicesize, "elements succesfully in", (endTime-startTime)/1000000, "ms.")
	} else {
		fmt.Println("Sort failed.")
	}
}

// sort modifies the slice s so that the integers are sorted in
// place using quicksort.
func sort(s []int) { quicksort(s, 0, len(s)-1) }

// quicksorts the section of s between the left bounding index and
// the right bounding index.
func quicksort(s []int, leftbound int, rightbound int) {
	if leftbound < rightbound {
		pivot := partition(s, leftbound, rightbound)
		quicksort(s, leftbound, pivot-1)
		quicksort(s, pivot+1, rightbound)
	}
}

// partitions the section of s between the left bounding index and
// the right bounding index.
func partition(s []int, leftbound int, rightbound int) int {
	pivotIndex := leftbound + ((rightbound - leftbound) / 2)
	pivotValue := s[pivotIndex]
	s[pivotIndex], s[rightbound] = s[rightbound], s[pivotIndex]
	tmp := leftbound

	for i := leftbound; i < rightbound; i++ {
		if s[i] < pivotValue {
			s[i], s[tmp] = s[tmp], s[i]
			tmp++
		}
	}
	s[tmp], s[rightbound] = s[rightbound], s[tmp]
	return tmp
}
