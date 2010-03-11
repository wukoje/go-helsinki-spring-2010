package main

import (
	"fmt"
)

func search(s []int, e int) int {
	return seek(s, 0, len(s)-1, e)
}

func seek(s []int, a int, b int, e int) int {
	if s == nil || len(s) < 1 || b-a < 0 {
		return -1
	}
	if a == b {
		if s[a] == e {
			return a
		} else {
			return -1
		}
	}
	p := a+(b-a)/2
	if s[p] > e {
		return seek(s, a, p-1, e)
	} else if s[p] < e {
		return seek(s, p+1, b, e)
	}
	return p
}

func main() {
	a := []int{1,2,3,4,5,6,7,8,9,10,11,67,92}
	fmt.Printf("a: %v, index of 1: %v, index of 8: %v, index of 666: %v\n", a, search(a, 1), search(a, 8), search(a, 666))
	b := []int{1,2,3,5,9,11,36,54,77,102,205,666,999,1024}
	fmt.Printf("b: %v, index of 8: %v, index of 666: %v, index of 103: %v\n", b, search(b, 8), search(b, 666), search(b, 103))
	c := []int{1,2,3,5,9,11,36,54,77,102,103,104,105,205,666}
	fmt.Printf("c: %v, index of 8: %v, index of 666: %v, index of 103: %v\n", c, search(c, 8), search(c, 666), search(c, 103))
}

