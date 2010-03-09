package main

import (
	"fmt"
	"rand"
	"time"
)


func partition(s []int) int {
	l := 0
	r := len(s) - 1

	pivot := s[0]
	for {
		for s[r] > pivot {
			r--
		}
		for s[l] < pivot {
			l++
		}

		if l < r {
			s[l], s[r] = s[r], s[l]
		} else {
			break
		}
	}
	return r
}

// a straight-forward implementation of quicksort from CLRS
func sort(s []int) {
	if len(s) > 0 {
		q := partition(s)
		sort(s[0:q])
		sort(s[q+1 : len(s)])
	}
}

func sorted(s []int) bool {
	for i, _ := range s {
		if i > 0 && s[i] < s[i-1] {
			return false
		}
	}
	return true
}

func main() {
	rand.Seed(time.Seconds())
	ints := rand.Perm(1000)

	sort(ints)

	if sorted(ints) {
		fmt.Println("Sorted :-)")
	} else {
		fmt.Println("Unsorted :-(")
	}
}
