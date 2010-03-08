package main

import "fmt"

func search(s []int, e int) int {
	if len(s) > 0 {
		i := len(s) / 2
		switch {
		case e == s[i]:
			return i
		case e < s[i]:
			return search(s[0:i], e)
		case e > s[i]:
			tmp := search(s[i+1:], e)
			if tmp != (-1) {
				tmp += i + 1
			}
			return tmp
		}
	}
	return -1
}

func main() {
	var allGood bool

	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Looking for things that we're supposed to find...")
	allGood = true
	for _, v := range a {
		x := search(a, v)
		if x != v {
			allGood = false
			break
		}
	}
	fmt.Println("success =", allGood)

	notThere := []int{-10, -54, 23, 77}
	fmt.Println("Looking for things that aren't there...")
	allGood = true
	for _, v := range notThere {
		x := search(a, v)
		if x != (-1) {
			allGood = false
			break
		}
	}
	fmt.Println("success =", allGood)
}
