package main

import (
	"fmt"
	"rand"
	"time"
)

const (
	SLICE_SIZE int = 31
	MAX_NUMBER = 99
)

func sort(s []int) {
	quicksort(s, 0, len(s)-1)
}

func quicksort(s []int, leftIndex int, rightIndex int) {

	if rightIndex>leftIndex {
		// choose pivot point from the middle
		pivotIndex := partition(s, leftIndex, rightIndex, (leftIndex+rightIndex)/2)
		quicksort(s, leftIndex, pivotIndex-1)
		quicksort(s, pivotIndex+1, rightIndex)
	}
}

func partition(s []int, leftIndex int, rightIndex int, pivotIndex int) int {
	pivotValue := s[pivotIndex]
	s[pivotIndex], s[rightIndex] = s[rightIndex], s[pivotIndex]
	storeIndex := leftIndex
	for i := leftIndex; i< rightIndex; i++ {
		if s[i] <= pivotValue {
			s[i], s[storeIndex] = s[storeIndex], s[i]
			storeIndex++
		}
	}
	s[storeIndex], s[rightIndex] = s[rightIndex], s[storeIndex]
	return storeIndex
}

func main() {
	slice := make([]int, SLICE_SIZE)
	rand.Seed(time.Nanoseconds())

	for i := 0; i < SLICE_SIZE; i++ {
		slice[i] = rand.Intn(MAX_NUMBER)
	}
	
	fmt.Printf("slice before sort=%v\n", slice)
	sort(slice)
	fmt.Printf("slice after sort=%v\n", slice)
}
