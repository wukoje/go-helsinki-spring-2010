package main

import (
	"fmt"
	"rand"
	"sort"
)

// search returns the index of the element e in the
// sorted slice s. If e does not exist in s search returns -1
func search(s []int, e int) int {
	start := 0
	end := len(s)
	for start < end {
		i := start + (end - start)/2
		if e > s[i] {
			start = i+1
		} else if e < s[i] {
			end = i
		} else {
			return i
		}
	}
	return -1
}

func main() {
	ints := rand.Perm(1000)
	sort.SortInts(ints)

	for len(ints) > 0 {
		for i, c := range ints {
			res := search(ints, c)
			if res != i {
				fmt.Println("fail: got ", res, " expected ", i, " with ", ints)
				return
			}
		}
		ints = ints[0:len(ints)-1]
	}
}
