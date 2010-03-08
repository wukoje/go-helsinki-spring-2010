package main

import "fmt"

func main() {
	test(make([]int, 0))

	s := make([]int, 5)
	for i, _ := range(s) {
		s[i] = i + 1
	}
	test(s)
    	
	test(fibs(10))
}

func test(s []int) {
	fmt.Printf("Before:\t%v\n", s)
	reverse(s)
	fmt.Printf("After:\t%v\n\n", s)
}

func reverse(s []int) {
	l := len(s)
	for i := l / 2 - 1 // the rightmost index to be swapped
		i >= 0;
		i-- {
		j := l - i - 1
		s[i], s[j] = s[j], s[i]
	}
}

func fibs(n int) []int {
	curr, next := 1, 1
	fs := make([]int, n)
	for i, _ := range fs {
		fs[i] = curr
		curr, next = next, curr + next
	}
	return fs
}
