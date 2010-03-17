package main

import (
	"fmt"
)


func main() {
	s := []int{10,15,20,25,30,35,44,56,90,91,100,101,156,157,190,250}

	fmt.Printf("Test slice: ")
	for _, v := range s {
		fmt.Printf("%d ", v)
	} 
	
	fmt.Printf("\nFind 57, result: %d", search(s,57))
        fmt.Printf("\nFind 20, result: %d", search(s,20))
        fmt.Printf("\nFind 101, result: %d", search(s, 101))
        fmt.Printf("\nFind 249, result: %d", search(s,249))
        fmt.Printf("\nFind 250, result: %d", search(s, 250))
}


func search (s []int, e int) int {
	low := 0
	high := len(s)-1
	middle := 0

	for low <= high {
		middle = low + (high-low)/2
		switch {
			case s[middle] == e: 
				return middle
			case s[middle] < e:
				low = middle+1
			default:
				high=middle-1
		}

	}	
	return -1
}
