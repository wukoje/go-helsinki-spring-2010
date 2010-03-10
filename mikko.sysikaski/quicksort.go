package main

import (
	"fmt"
	"rand"
	"time"
	. "sort"
)

// sort modifies the slice s so that the integers are sorted in
// place using quicksort
func sort(s []int) {
	n := len(s)
	if n < 2 {
		return
	}
	pivot := rand.Intn(n)
	p := s[pivot]
	s[pivot] = s[n-1]

	k := partition(s[0:n-1], p)
	s[n-1] = s[k]
	s[k] = p

	sort(s[0:k])
	sort(s[k+1 : n])
}

func partition(s []int, pivot int) int {
	start, end := 0, len(s)
	for {
		for {
			if start == end {
				return start
			} else if s[start] <= pivot {
				start++
			} else {
				break
			}
		}
		end--
		for {
			if start == end {
				return start
			} else if s[end] > pivot {
				end--
			} else {
				break
			}
		}
		s[start], s[end] = s[end], s[start]
		start++
	}
	panic("unreachable")
}


func main() {
	n := 1 << 20
	arr := make([]int, n)
	arr2 := make([]int, n)
	for i, _ := range arr {
		arr[i] = rand.Int()
		arr2[i] = arr[i]
	}

	t0 := time.Nanoseconds()
	sort(arr)
	fmt.Println("quicksort time:", float(time.Nanoseconds()-t0)/1e9)

	t0 = time.Nanoseconds()
	SortInts(arr2)
	fmt.Println("sort.SortInts time:", float(time.Nanoseconds()-t0)/1e9)

	// Check for correctness
	for i, _ := range arr {
		if arr[i] != arr2[i] {
			fmt.Println("Sorting failed!")
			break
		}
	}
}
