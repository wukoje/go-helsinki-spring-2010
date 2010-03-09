/*
 * binary_search.go: Binary search for slices containing integers.
 *
 * For Introduction to Go, Spring 2010
 * Kimmo Kulovesi <kkuloves@cs.helsinki.fi>
 */
package main

import (
	"fmt"
)

// Returns the index of the (or an) element with value e
// in the sorted slice s, or -1 if no element has value e.
func search(s []int, e int) int {
	left, right := 0, len(s)

	for left < right {
		// Avoid overflowing the integer in case s is _very_ large
		middle := left + ((right - left) / 2)

		if s[middle] < e {
			// e, if it exists, is to the right of center
			left = middle + 1
		} else {
			// e, if it exists, is in the middle now or to the left of  it
			right = middle
		}
	}

	// To save comparisons in the loop, the test for equality is deferred here
	if left < len(s) && s[left] == e {
		return left
	}
	return -1
}

// A function to calculate the correct index for a given element in
// our generated test data where the data[i] = i * 2. Returns -1
// if the e doesn't exist in the test data at an index lower than
// sliceLength.
func indexInTestData(e int, sliceLength int) int {
	if e%2 != 0 {
		return -1
	}
	e /= 2
	if e >= 0 && e < sliceLength {
		return e
	}
	return -1
}

// The size of the test array
const testDataLength = 5

// Test the binary search implementation
func main() {
	data := make([]int, testDataLength)
	for i := 0; i < len(data); i++ {
		// Make the test data such that data[i] = i * 2
		data[i] = i * 2
	}

	// Test searching in slices of the data
	fmt.Printf("Test data: %v\n", data)
	errorCount := 0
	for length := 0; length <= len(data); length++ {
		s := data[0:length]
		fmt.Printf("\nTesting slice of length %d\n", length)
		for i := length * 2; i >= -1; i-- {
			foundAt := search(s, i)
			if foundAt == indexInTestData(i, len(s)) {
				if foundAt >= 0 {
					fmt.Printf("%d is in index %d\n", i, foundAt)
				} else {
					fmt.Printf("%d is not present\n", i)
				}
			} else {
				fmt.Printf("ERROR: Got %d for element %d, should have %d\n",
					foundAt, i, indexInTestData(i, len(s)))
				errorCount++
			}
		}
	}

	fmt.Printf("\nThere were %d errors.\n", errorCount)
}
