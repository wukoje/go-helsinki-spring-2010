package main

import "fmt"

func main() {
	empty := []int{}
	fmt.Printf("Looking for something in empty: %v\n\n", search(empty, 0))

	evens := []int{0, 2, 4, 6, 8, 10}
	for i := 0; i <= 10; i++ {
		fmt.Printf("Looking for %v in evens: %v\n", i, search(evens, i))
	}
	fmt.Printf("\n")
	
	odds := []int{1, 3, 5, 7, 9}
	for i := 0; i <= 10; i++ {
		fmt.Printf("Looking for %v in odds: %v\n", i, search(odds, i))
	}
}

func search(s []int, e int) int {
	if len(s) > 0 {
		left := 0
		right := len(s) - 1
		for left < right {
			mid := (left + right) / 2
			if e < s[mid] {
				right = mid - 1
			} else if e > s[mid] {
				left = mid + 1
			} else {
				left, right = mid, mid
			}
		}
		if s[left] == e {
			return left
		}
	}
	return -1
}
