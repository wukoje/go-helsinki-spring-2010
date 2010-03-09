package main

import "fmt"
import "rand"
import "time"

func main() {
	s := make([]int, 20)
	rand.Seed(time.Nanoseconds())
	for i, _ := range s {
		s[i] = rand.Intn(100)
	}
	
	fmt.Printf("%v\n", s)
	sort(s)
	fmt.Printf("%v\n", s)
}

func sort(s []int) {
	quicksort(s, 0, len(s) - 1)
}

func quicksort(s []int, start int, end int) {
	left := start
	right := end
	pivot := s[average(start, end)]
	
	for {
		for s[left] < pivot {
			left++
		}
		for pivot < s[right] {
			right--
		}
		if left <= right {
			s[left], s[right] = s[right], s[left]
			left++
			right--
		}
		if left >= right {
			break
		}
	}
	
	if start < right {
		quicksort(s, start, right)
	}
	if left < end {
		quicksort(s, left, end)
	}
}

func average(a int, b int) int {
	return a + (b - a) / 2
}
