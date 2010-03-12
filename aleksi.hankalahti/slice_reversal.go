package main

import "fmt"

// reverse reverses the contents of s in place
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%v\n", s);
	reverse(s);
	fmt.Printf("%v\n", s);
}
