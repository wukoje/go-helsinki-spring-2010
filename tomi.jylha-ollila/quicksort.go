package main


import (
	"fmt"
	"rand"
)


func partition(s []int) ([]int, []int) {
	pivot_index := rand.Intn(len(s))
	pivot := s[pivot_index]
	s[pivot_index], s[len(s)-1] = s[len(s)-1], s[pivot_index]
	border := 0
	for i, _ := range s {
		if s[i] < pivot {
			s[i], s[border] = s[border], s[i]
			border++
		}
	}
	s[border], s[len(s)-1] = s[len(s)-1], s[border]
	return s[0:border], s[border+1 : len(s)]
}


func sort(s []int) {
	for len(s) > 1 {
		left, right := partition(s)
		if len(left) < len(right) {
			sort(left)
			s = right
		} else {
			sort(right)
			s = left
		}
	}
}


func main() {
	s := make([]int, 10)
	for i, _ := range s {
		s[i] = rand.Intn(100)
	}
	fmt.Printf("%v\n", s)
	sort(s)
	fmt.Printf("%v\n", s)
	s = make([]int, 19)
	for i, _ := range s {
		s[i] = rand.Intn(100)
	}
	fmt.Printf("%v\n", s)
	sort(s)
	fmt.Printf("%v\n", s)
	s = make([]int, 30)
	for i, _ := range s {
		s[i] = 30 - i
	}
	fmt.Printf("%v\n", s)
	sort(s)
	fmt.Printf("%v\n", s)
}
