package main

import "fmt"

func main() {
	var l = [...]int {1,2,3,4,5,6,7,8,9,10}
	
	// Search 6 in the array
	n := search(&l, 6)
	fmt.Printf("Pos: %v\n", n)
	
	// Search 8 in the array
	n = search(&l, 8)
	fmt.Printf("Pos: %v\n", n)
	
	// Search 99 in the array
	n = search(&l, 99)
	fmt.Printf("Pos: %v\n", n)

}

func search(s []int, e int) int {
	n := -1
	left := 0
	right := len(s) - 1
	
	for left < right {
		mid := left + (right - left)/2
		switch {
		case s[mid] < e:
			left = mid + 1
		case s[mid] == e:
			left, right, n = mid, mid, mid
		case s[mid] > e:
			right = mid - 1
		}
	}
	
	return n
}

