package main

import (
	"fmt"
	"os"
	"rand"
)

// sort modifies the slice s so that the integers are sorted in
// place using quicksort
func sort(s []int) {
	first, middle, last := 0, (len(s)-1)/2, len(s)-1
	// Pivot is median of first, middle and last value in the slice
	var pivot int
	if s[first] <= s[middle] && s[middle] <= s[last] {
		pivot = middle
	} else if s[middle] <= s[first] && s[first] <= s[last] {
		pivot = first
	} else {
		pivot = last
	}
	pivotVal := s[pivot]
	s[pivot], s[last] = s[last], s[pivot]
	si := 0
	for i := 0; i < last; i++ {
		if s[i] <= pivotVal {
			s[i], s[si] = s[si], s[i]
			si += 1
		}
	}
	s[si], s[last] = s[last], s[si]
	if (si > 1) {
		sort(s[0:si])
	}
	if (si < last) {
		sort(s[si+1:len(s)])
	}
}

func main() {
	_, ms, _ := os.Time()
	r := rand.New(rand.NewSource(ms))
	const size = 8
	x := make([]int, size)
	for i := 0; i < len(x); i++ {
		x[i] = r.Intn(42)
	}
	fmt.Println("Sort this:", x)
	sort(x)
	fmt.Println("Result:", x)
}
