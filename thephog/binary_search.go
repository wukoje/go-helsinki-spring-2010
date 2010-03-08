package main

import "fmt"

func search(s []int, e int) int {

	left := 0
	right := len(s) - 1

	for {
		middle := (left + right) / 2

		if s[middle] == e {
			return middle
		} else if e > s[middle] {
			left = middle + 1
		} else {
			right = middle - 1
		}

		if left > right {
			break
		}
	}

	return -1
}

func printSlice(s []int) {
	for i, value := range s {
		fmt.Printf("%d, ", value)
		if i%25 == 0 {
			fmt.Printf("\n")
		}

	}
	fmt.Printf("\n")
}

func initSlice(length int) []int {
	retSlice := make([]int, length)
	for i, _ := range retSlice {
		retSlice[i] = i
	}
	return retSlice
}

func main() {
	s := initSlice(66)
	fmt.Printf("Original slice of size 66:\n")
	printSlice(s)

	a, b, c, d := 5, 15, 33, 10000
	fmt.Printf("Searching for %d, %d, %d, %d\n", a, b, c, d)

	if i := search(s, a); i != -1 {
		fmt.Printf("Found %d at %d!\n", a, i)
	}

	if i := search(s, b); i != -1 {
		fmt.Printf("Found %d at %d!\n", b, i)
	}

	if i := search(s, c); i != -1 {
		fmt.Printf("Found %d at %d!\n", c, i)
	}

	if i := search(s, d); i != -1 {
		fmt.Printf("Found %d at %d!\n", d, i)
	}

	s = initSlice(1333)
	fmt.Printf("Original slice of size 1333:\n")
	printSlice(s)

	a, b, c, d = 667, 13, 99, 1337
	fmt.Printf("Searching for %d, %d, %d, %d\n", a, b, c, d)
	if i := search(s, a); i != -1 {
		fmt.Printf("Found %d at %d!\n", a, i)
	}

	if i := search(s, b); i != -1 {
		fmt.Printf("Found %d at %d!\n", b, i)
	}

	if i := search(s, c); i != -1 {
		fmt.Printf("Found %d at %d!\n", c, i)
	}

	if i := search(s, d); i != -1 {
		fmt.Printf("Found %d at %d!\n", d, i)
	}

}
