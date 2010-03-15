
package main

import "fmt"

func search(s []int, e int) (i int) {
	left := 0
	right := len(s)

	for left < right {
		// watch out for integer overflow
		mid := left + (right - left) / 2 // =(left+right)/2
		switch {
		case e < s[mid]:
			right = mid
		case e == s[mid]:
			return mid
		case e > s[mid]:
			left = mid + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{ 5 }, 5))
	fmt.Println(search([]int{ 5 }, 1))
	fmt.Println(search([]int{ -1, 1 }, -1))
	fmt.Println(search([]int{ -1, 1 }, 1))
	fmt.Println(search([]int{ -1, 1 }, 0))
	fmt.Println(search([]int{ 0, 10, 20, 30, 40 }, 0))
	fmt.Println(search([]int{ 0, 10, 20, 30, 40 }, 40))
}
