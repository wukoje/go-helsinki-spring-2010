package main

import (
	"fmt"
)

// reverse reverses the contents of s in place
func reverse(s []int) {
	for i := 0; 2*i < len(s); i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}

func main() {
	s := [...]int{1, 2, 3, 4, 5, 6}
	t := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(s)
	reverse(&s)
	fmt.Println(s)

	fmt.Println(t)
	reverse(t[0:3])
	reverse(t[3:6])
	reverse(t[6:9])
	fmt.Println(t)
}
