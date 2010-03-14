/*
1 quicksort.go
// sort modifies the slice s so that the integers are sorted in
// place using quicksort
func sort(s []int)
*/
package main

import (
	"fmt";
	"rand"
)

func sort(s []int) {
	if len(s) > 1 {
		pivot := rand.Int() % (len(s))
		pivot = part(s, pivot)
		sort(s[0:pivot])
		sort(s[pivot:])
	}
}

func part(s []int, pivot int) int {
	last := len(s)-1
	pv := s[pivot]
	s[pivot], s[last] = s[last], s[pivot]
	low := 0
	for i:=0; i<last; i++ {
		if s[i] <= pv {
			s[i], s[low] = s[low], s[i]
			low++
		}
	}
	s[low], s[last] = s[last], s[low]

	return low
}

func isInOrder(s []int) bool {
	for i, j, l := 0, 1, len(s); j < l; i, j = i+1, j+1 {
		if s[i] > s[j] {
			return false
		}
	}
	return true
}

func test(size int, doPrint bool) {
	s := make([]int, size)

	for i := 0; i < size; i++ {
		s[i] = rand.Int() % 10000
	}
	fmt.Printf("Created slice with %d elements\n", size)
	fmt.Printf("Random slice in order? %v\n", isInOrder(s))
	if doPrint {
		fmt.Printf("Let's sort slice:\n%v\n", s)
	}
	sort(s)
	fmt.Printf("Done. In order? %v\n", isInOrder(s))
	if doPrint {
		fmt.Printf("Now it looks like:\n%v\n", s)
	}

	fmt.Printf("Sorting sorted slice...\n")
	sort(s)
	fmt.Printf("Done\n")
}

func main() {
	test(32, true)
	test(65536, false)
}
