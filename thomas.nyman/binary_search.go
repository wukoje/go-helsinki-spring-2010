package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	testsearch(s, 0, 10)

	s = []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}
	testsearch(s, 0, 20)

	s = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}
	testsearch(s, 0, 50)
}

// runs search in s for values in range min - max and prints results.
func testsearch(s []int, min int, max int) {

	fmt.Println("Slice is", s)
	fmt.Println("Searched slice for values in range", min, "-", max)

	for i := min; i <= max; i++ {
		if result := search(s, i); result >= 0 {
			fmt.Println("Found", i, "at index", result)
		}
	}
}

// search returns the index of the element e in the sorted
// slice s. If e does not exist in s search returns -1
func search(s []int, e int) int {
	leftbound := 0
	rightbound := len(s) - 1

	for leftbound <= rightbound {
		center := (rightbound + leftbound) / 2

		switch {
		case s[center] == e:
			return center
		case s[center] < e:
			leftbound = center + 1
		case s[center] > e:
			rightbound = center - 1
		}
	}
	return -1
}
