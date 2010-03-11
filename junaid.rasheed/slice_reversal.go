package main

import "fmt"

func main() {

	s := []int{1, 2, 3}
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	s2 := []int{1, 3, 5, 7, 9}

	fmt.Print(s)
	reverse(s)
	fmt.Print(s)

	fmt.Print(s1)
	reverse(s1)
	fmt.Print(s1)

	fmt.Print(s2)
	reverse(s2)
	fmt.Print(s2)

}

func reverse(s []int) {

	for i := len(s) - 1; i >= len(s)/2; i-- {
		temp := s[(len(s)-1)-i]
		s[(len(s)-1)-i] = s[i]
		s[i] = temp
	}

}

