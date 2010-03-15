package main

import (
	"fmt"
)

// sort quicksorts s
func sort(s []int) {
	if len(s) <= 1 {
		return
	}

	pivot := len(s) / 2
	pivot = partition(s, pivot)
	
	sort(s[0:pivot])
	sort(s[pivot+1 : len(s)])
}

// partition is the in-place partition function for quicksort
func partition(s []int, pivot int) int {

	pivotValue := s[pivot]
	s[len(s)-1], s[pivot] = s[pivot], s[len(s)-1]
	storeIndex := 0
	
	for i := 0; i < len(s)-1; i++ {
		if s[i] <= pivotValue {
			s[i], s[storeIndex] = s[storeIndex], s[i]
			storeIndex = storeIndex + 1
		}
	}
	
	s[len(s)-1], s[storeIndex] = s[storeIndex], s[len(s)-1]
	
	return storeIndex
}

func main() {
	s := [...]int{3, 7, 3, 9, 16, 5, 14, 235, 1, 32, 1, 6, 7, 3, 9, 4, 235}

	fmt.Println(s)
	sort(&s)
	fmt.Println(s)
}
