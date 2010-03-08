package main

import "fmt"

func main() {

	// reverse slice with even number of elements.
	s := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Print(s, " reversed is ")
	reverse(s)
	fmt.Println(s)

	// reverse slice with odd number of elements.
	s = []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Print(s, " reversed is ")
	reverse(s)
	fmt.Println(s)
}

// reverse reverses the contents of s in place
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
