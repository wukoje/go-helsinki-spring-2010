package main

import (
	"fmt"
	"rand"
)

func partition(s []int) ([]int, []int) {
	pi := rand.Intn(len(s))
	p := s[pi]
	s[pi], s[len(s)-1] = s[len(s)-1], s[pi]
	si := 0
	for i, _ := range s {
		if s[i] < p {
			s[i], s[si] = s[si], s[i]
			si++
		}
	}
	s[si], s[len(s)-1] = s[len(s)-1], s[si]
	return s[0:si], s[si+1 : len(s)]
}

// sort modifies the slice s so that the integers are sorted in
// place using quicksort
func sort(s []int) {
	if len(s) > 1 {
		l, r := partition(s)
		sort(l)
		sort(r)
	}
}

func main() {
	s := make([]int, 25)
	for i, _ := range s {
		s[i] = rand.Intn(100)
	}
	fmt.Println(s)
	sort(s)
	fmt.Printf(">> %v\n", s)

	s = make([]int, 1234567)
	for i, _ := range s {
		s[i] = rand.Intn(1234567)
	}
	sort(s)
}
