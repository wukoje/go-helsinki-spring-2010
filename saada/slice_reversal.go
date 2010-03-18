package main

import (
	"fmt"
)

func reverse(s []int) {
	n := len(s)-1
	for i := 0; i<n/2; i++ {
		x := s[i]
		s[i] = s[n-i]
		s[n-i] = x
	}
}

func main() {
	s1 := make([]int, 5)
	for i,_ := range s1 {
		s1[i] = i+1
	}
	fmt.Println(s1)
	reverse(s1)
	fmt.Println(s1)

	s2 := make([]int, 15)
	for i,_ := range s2 {
		s2[i] = i+4
	}
	for i,_ := range s2 {
		s2[i] = i%4
	}
	fmt.Println(s2)
	reverse(s2)
	fmt.Println(s2)
}
