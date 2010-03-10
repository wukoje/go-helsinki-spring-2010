package main

import "fmt"

func search(s []int, e int) int {
	l, r := 0, len(s)-1
	for p := r / 2; r >= l; p = (l + r) / 2 {
		switch {
		case s[p] == e:
			return p
		case s[p] < e:
			l = p + 1
		case s[p] > e:
			r = p - 1
		}
	}
	return -1
}

func main() {
	a := make([]int, 100000000)
	for i, _ := range a {
		a[i] = i * 3
	}

	//	fmt.Printf("Slice:\t%v\n", a)
	fmt.Printf("Index of %d:%d\n", 9, search(a, 9))
	fmt.Printf("Index of %d:%d\n", 45, search(a, 45))
	fmt.Printf("Index of %d:%d\n", 63, search(a, 63))
}
