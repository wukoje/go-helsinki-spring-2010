package main

import (
	"fmt"
)

func main() {
	s := make([]int, 11)
	for i, _ := range s {
		s[i] = i + 1
	}
	fmt.Println(s)
	fmt.Println("0 is too small:", search(s, 0))
	fmt.Println("12 is too big:", search(s, 0))
	fmt.Println("2 at index:", search(s, 2))
	fmt.Println("6 at index:", search(s, 6))
	fmt.Println("11 at index:", search(s, 11))

	s = make([]int, 0)
	fmt.Println()
	fmt.Println(s)
	fmt.Println("1 is not in in empty slice:", search(s, 1))
	fmt.Println("and neither is 4:", search(s, 4))
}

func search(s []int, e int) int {
	low := 0
	high := len(s) - 1

	for low <= high {
		mid := (low + high) / 2
		switch {
		case e == s[mid]:
			return mid
		case e < s[mid]:
			high = mid - 1
		case e > s[mid]:
			low = mid + 1
		}
	}
	return -1
}
