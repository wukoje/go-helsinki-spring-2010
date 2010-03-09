/*
 * quicksort.go: Quicksort for ints.
 *
 * For Introduction to Go, Spring 2010
 * Kimmo Kulovesi <kkuloves@cs.helsinki.fi>
 */
package main

import (
	"fmt"
	"rand"
	srt "sort"
)

// Sort the slice s in place using quicksort.
func sort(s []int) {
	if len(s) < 2 {
		// Slices of length 0 and 1 are already sorted
		return
	}

	// Pick the middle element as the pivot and move it to the end
	// temporarily so that it's out of the way while we sort
	pivotIndex := len(s) / 2
	pivot := s[pivotIndex]
	lastIndex := len(s) - 1
	s[pivotIndex], s[lastIndex] = s[lastIndex], s[pivotIndex]

	// Move any element smaller than the pivot to the beginning
	pivotIndex = 0
	for i := 0; i < lastIndex; i++ {
		if s[i] < pivot {
			s[i], s[pivotIndex] = s[pivotIndex], s[i]
			pivotIndex++
		}
	}

	// Move the pivot element to its final place
	s[pivotIndex], s[lastIndex] = s[lastIndex], s[pivotIndex]

	// Sort the elements smaller than the pivot
	sort(s[0:pivotIndex])
	// Sort the elements greater than the pivot
	sort(s[pivotIndex+1 : len(s)])
}

// Maximum length of array to test the sorting of
const maxTestDataLength = 2000

// Maximum length of array to actually print
const maxPrintableLength = 10

// Test sorting on random slices of length 0..maxTestDataLength
func main() {
	errorCount := 0
	for length := 0; length <= maxTestDataLength; length++ {
		// Make a test slice and populate it with random numbers
		data := make([]int, length)
		for i := len(data) - 1; i >= 0; i-- {
			data[i] = rand.Intn(maxTestDataLength)
		}

		// Sort the slice and confirm the result
		if length <= maxPrintableLength {
			fmt.Printf("Unsorted: %v\n", data)
		}
		sort(data)
		if length <= maxPrintableLength {
			fmt.Printf("Sorted: %v\n", data)
		}
		if !srt.IntsAreSorted(data) {
			fmt.Printf("Error: Sort failed, length %d!\n", len(data))
			errorCount++
		}
	}
	fmt.Printf("\nThere were %d errors up to array length %d.\n",
		errorCount, maxTestDataLength)
}
