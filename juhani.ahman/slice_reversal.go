package main

import fmt "fmt"

// reverse reverses the contents of s in place
func reverse(s []int) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

func main() {
	s0 := []int{1, 2, 3}
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{99, 55, 33, 00, 11, 4, 5, 6, 7}
	fmt.Println(s0, s1, s2)

	//do reverse
	reverse(s0)
	reverse(s1)
	reverse(s2)

	fmt.Println(s0, s1, s2)
}
