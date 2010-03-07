package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 10)
	for i, _ := range slice {
		slice[i] = i + 1
	}
	fmt.Println(slice)
	reverse(slice)
	fmt.Println(slice)

	slice = make([]int, 5)
	for i, _ := range slice {
		slice[i] = i + 1
	}
	fmt.Println(slice)
	reverse(slice)
	fmt.Println(slice)

	slice = make([]int, 6)
	for i, _ := range slice {
		slice[i] = i + 1
	}
	fmt.Println(slice)
	reverse(slice)
	fmt.Println(slice)
}

func reverse(s []int) {
	head := 0
	tail := len(s) - 1
	for head < tail {
		s[head], s[tail] = s[tail], s[head]
		head++
		tail--
	}
}
