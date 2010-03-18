package main

import (
	"fmt"
)

func main() {

	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}
	s3 := []int{4, 8, 15, 16, 23, 42}
	fmt.Println("s1 = ", s1)
	fmt.Println("Searching 9", search(s1, 9))
	fmt.Println("Searching 4", search(s1, 4))
	fmt.Println("Searching 10", search(s1, 10))
	fmt.Println("s2 = ", s2)
	fmt.Println("Searching 19", search(s2, 19))
	fmt.Println("Searching 10", search(s2, 10))
	fmt.Println("Searching 23", search(s2, 23))
	fmt.Println("s3 = ", s3)
	fmt.Println("Searching 4", search(s3, 4))
	fmt.Println("Searching 42", search(s3, 42))
	fmt.Println("Searching 10", search(s3, 10))

}

func search(s []int, e int) int {

	max := len(s) - 1
	found := false
	a := s
	i := cap(a)
	for max+1 > 0 && found == false {

		half := max / 2
		//	fmt.Println(" ", half, " ", max, " ", a[half], " ",a, found, "a[half] == e ", a[half] == e )
		if a[half] == e {

			found = true
			return half + i - cap(a)

		} else if a[half] > e {

			a = a[0:half]
			max = len(a) - 1

		} else {

			a = a[half+1:]
			max = len(a) - 1
		}

	}
	return max

}
