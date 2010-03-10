package main

import (
	"fmt"
)

func main() {

	test := []int{2, 4, 7, 3, 5, 6}
	test2 := []int{546, 534, 54, 5, 7, 3}
	test3 := []int{1, 2, 3, 4, 5, 6, 7}
	test4 := []int{2, 3, 5, 7, 11, 13, 17, 19, 23}
	fmt.Println("Before: ", test)
	fmt.Println("After:", reverse(test))
	fmt.Println()
	fmt.Println("Before:", test2)
	fmt.Println("After:", reverse(test2))
	fmt.Println()
	fmt.Println("Before:", test3)
	fmt.Println("After:", reverse(test3))
	fmt.Println()
	fmt.Println("Before:", test4)
	fmt.Println("After:", reverse(test4))

}

// reverse reverses the contents of s in place
func reverse(s []int) []int {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}
