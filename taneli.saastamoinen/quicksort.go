package main

import (
	"fmt"
	"rand"
	"time"
)

func quicksort(s []int) {
	qsort(s, 0, len(s)-1)
}

func qsort(s []int, a int, b int) {
	if b-a < 1 {
		return
	}
	pivot := a+(b-a)/2
	s[pivot], s[b] = s[b], s[pivot]
	left := a
	for i := a; i < b; i++ {
		if s[i] < s[b] {
			s[i], s[left] = s[left], s[i]
			left++
		}
	}
	s[left], s[b] = s[b], s[left]
	qsort(s, a, left-1)
	qsort(s, left+1, b)
}

func main() {
	a1 := []int{1,3,6,21,4,2,99,35,14,22}
	s1 := make([]int, len(a1))
	copy(s1, a1)

	fmt.Printf("%v -> ", s1)
	quicksort(s1)
	fmt.Printf("%v\n", s1)

	rand.Seed(time.Nanoseconds())

	for i:=0; i<10; i++ {
		s2 := rand.Perm(11)
		fmt.Printf("%v -> ", s2)
		quicksort(s2)
		fmt.Printf("%v\n", s2)
	}
}

