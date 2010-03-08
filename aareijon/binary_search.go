package main

import fmt "fmt"

// search returns the index of the element e in the
// sorted slice s. If e does not exist in s search returns -1
func search(s []int, e int) int {
	low := 0
	high := len(s)-1
	
	for low <= high {
		mid := low + (high - low)/2 // avoid overflow
		if s[mid] == e {
			return mid
		} else if e < s[mid] {
			high = mid-1
		} else { // e > s[mid]
			low = mid+1
		}
	}
	return -1 
}

func main() {
	g := [...]int{0,1,2,3,4,5,6,7,8,9,10}
	fmt.Printf("11 not found: %d\n", search(&g, 11))
	fmt.Printf("-1 not found: %d\n", search(&g, -1))
	fmt.Printf("3 found at: %d\n", search(&g, 3))
	fmt.Printf("10 found at: %d\n", search(&g, 10))
	fmt.Printf("0 found at: %d\n", search(&g, 0))
	fmt.Printf("nothing found in empty table: %d\n", search(&[...]int{}, 5))
}
