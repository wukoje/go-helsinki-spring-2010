package main

import "fmt"


func search(s []int, e int) int {

	low := 0
	high := len(s) - 1
	for low < high {
		mid := low + ((high - low) / 2)
		if s[mid] < e {
			low = mid + 1

		} else {
			high = mid
		}
	}
	if low < len(s)-1 && s[low] == e {
		return low // found
	}
	return -1 // not found
}


func main() {

	s := []int{1, 2, 3, 4, 8, 9, 10, 12}
	e := 9

	result := search(s, e)
	fmt.Printf("Search for %v in %v", e, s)
	fmt.Printf("\nFound at index: %v\n", result)

	s1 := []int{5, 7, 8, 23, 65, 55}
	e = 50

	result = search(s1, e)
	fmt.Printf("Search for %v in %v", e, s1)
	fmt.Printf("\nNot found: %v\n", result)
}

