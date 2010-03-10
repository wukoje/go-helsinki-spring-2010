package main


import (
	"fmt"
)


func search(s []int, e int) int {
	start, end := 0, len(s)-1
	middle := (start + end) / 2
	for start <= end {
		switch {
		case e < s[middle]:
			end = middle - 1
		case e > s[middle]:
			start = middle + 1
		default:
			return middle
		}
		middle = (start + end) / 2
	}
	return -1
}


const (
	search_start = 0
	search_end   = 128
)


func main() {
	s := make([]int, 11)
	for i, _ := range s {
		s[i] = (i + 1) * (i + 1)
	}
	fmt.Printf("Searching from %v:\n", s)
	missed := 0
	for i := search_start; i <= search_end; i++ {
		result := search(s, i)
		if result >= 0 {
			fmt.Printf("Searching %d returned %d\n", i, result)
		} else {
			missed++
		}
	}
	fmt.Printf("Searched for numbers from %d to %d, missed %d numbers\n",
		search_start, search_end, missed)
}
