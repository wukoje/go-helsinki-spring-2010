package main

import fmt "fmt"

func search(s[] int, e int) int {
	left := 0
	right := len(s) -1
	for left <= right {
		middle := left + (right - left)/2
		if s[middle] > e {
			right = middle - 1
		}else if s[middle] < e {
			left = middle + 1
		}else if s[middle] == e {
			return middle
		}
	}
	return -1
}

func main() {

	s := []int{3,5,7,9,16,17,24,45,67}
	fmt.Println(s)
	fmt.Printf("%d\n", search(s, 5))
	fmt.Printf("%d\n", search(s, 8)) 
	fmt.Printf("%d\n", search(s, 67)) 
}


