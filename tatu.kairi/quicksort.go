/*
	 1 quicksort.go

// sort modifies the slice s so that the integers are sorted in
// place using quicksort
func sort(s []int)
 */

package main

import (
	"fmt"
	"rand"
)

// sort modifies the slice s so that the integers are sorted in
// place using quicksort
func sort(s []int){
	if len(s) <= 1 {
		return
	}

	less, pivotList, more := make([]int, len(s)), make([]int, len(s)), make([]int, len(s)) // FIXME: how to optimize the size of the list?

	lessIndex, pivotListIndex, moreIndex := 0, 0, 0

	pivot := s[0]

	for _, e := range s {

		if e < pivot {
			less[lessIndex] = e
			lessIndex += 1
		} else if e > pivot {
			more[moreIndex] = e
			moreIndex += 1
		} else {
			pivotList[pivotListIndex] = e
			pivotListIndex += 1
		}
	}

	// trim lists
	less = less[0:lessIndex]
	pivotList = pivotList[0:pivotListIndex]
	more = more[0:moreIndex]

	sort(less)
	sort(more)

	// copy values to result list
	i := 0
	for _,e := range less {
		s[i] = e
		i++
	}

	for _,e := range pivotList {
		s[i] = e
		i++
	}

	for _,e := range more {
		s[i] = e
		i++
	}
}

func main(){
	// FIXME no easier way to make an initialized slice?
	// eg. with arrays: s := [...]int{1,2,3,4}
	s := make([]int, 10)

	s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8], s[9] = 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
	fmt.Printf("Before sort: %v\n", s)
	sort(s)
	fmt.Printf("After: %v\n\n", s)

	s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8], s[9] = 10, 9, 8, 7, 6, 5, 4, 3, 2, 1
	fmt.Printf("Before sort: %v\n", s)
	sort(s)
	fmt.Printf("After: %v\n\n", s)

	for i,_ := range s {
		s[i] = rand.Intn(100)
	}
	fmt.Printf("Before sort: %v\n", s)
	sort(s)
	fmt.Printf("After: %v\n", s)

	fmt.Printf("Done!\n")
}
