package main

import "fmt"

// searches for the element e in the sorted slice s
// returns the index of e (some index if there are multiple e's)
// returns -1 if there is no e in s
func search(s []int, e int) int {
	left := 0
	right := len(s) - 1
	for left <= right {
		mid := left + (right - left)/2
		switch {
		case s[mid] == e:
			return mid
		case s[mid] < e:
			left = mid + 1
		case s[mid] > e:
			right = mid - 1
		}
	}
	return -1
}

func main() {
	slice := []int{1,5,17,22,40}
	fmt.Printf("slice: %v\n", slice)
	fmt.Printf("search(slice, 5): %v\n", search(slice, 5))
	fmt.Printf("search(slice, 10): %v\n", search(slice, 10))
	fmt.Printf("search([], 5): %v\n", search([]int{}, 5))
}
