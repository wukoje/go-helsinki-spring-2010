package main

import "fmt"

func reverse(s []int) {
	last := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		tmp := s[i]
		s[i] = s[last-i]
		s[last-i] = tmp
	}
}

func testcase(s []int) {
	fmt.Printf("%v\n", s)
	reverse(s)
	fmt.Printf("%v\n", s)
}

func main() {
	testcase([]int{})
	testcase([]int{1})
	testcase([]int{1, 2})
	testcase([]int{1, 2, 3})
	testcase([]int{1, 2, 3, 4})
	testcase([]int{1, 2, 3, 4, 5})
}
