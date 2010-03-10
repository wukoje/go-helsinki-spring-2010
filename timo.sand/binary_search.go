package main

import (
	"fmt"
)

func main() {

	s1 := []int{2, 3, 4, 5, 6}
	s2 := []int{2, 3, 5, 7, 11, 13, 7, 19, 23}
	s3 := []int{0, 10, 20, 30, 40, 50}
	s4 := []int{3, 4, 34, 45, 56, 67}

	fmt.Println("Looking for", 5, "in:", s1)
	fmt.Printf("%d\n", search(s1, 5))
	fmt.Println("Looking for", 5, "in:", s2)
	fmt.Printf("%d\n", search(s2, 5))
	fmt.Println("Looking for", 5, "in:", s3)
	fmt.Printf("%d\n", search(s3, 5))
	fmt.Println("Looking for", 5, "in:", s4)
	fmt.Printf("%d\n", search(s4, 5))
}

// search returns the index of the element e in the
// sorted slice s. If e does not exist in s search returns -1
func search(s []int, e int) int {

	left := 0
	right := len(s) - 1
	mid := int(left + right/2)

	for {
		if e < s[left] || e > s[right] {
			return -1
		} else if e > s[mid] {
			left = mid + 1
			mid = int((left + right) / 2)
		} else if e < s[mid] {
			right = mid - 1
			mid = int((left + right) / 2)
		} else if e == s[mid] {
			return mid
		} else {
			return -1
		}
	}

	return -1
}
