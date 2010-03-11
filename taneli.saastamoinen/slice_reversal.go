package main

import (
	"fmt"
)

func reverse(s []int) {
	for i,j := 0, len(s)-1; i < j; i,j = i+1, j-1 {
		s[i],s[j] = s[j],s[i]
	}
}

func main() {
	a1 := []int{1,2,3,4,5}
	a2 := []int{666,1911,348,123}
	s1 := make([]int, len(a1))
	s2 := make([]int, len(a2))
	copy(s1, a1)
	copy(s2, a2)

	fmt.Printf("%v -> ", s1)
	reverse(s1)
	fmt.Printf("%v\n", s1)

	fmt.Printf("%v -> ", s2)
	reverse(s2)
	fmt.Printf("%v\n", s2)
}

