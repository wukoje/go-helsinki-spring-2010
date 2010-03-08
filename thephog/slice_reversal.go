package main

import "fmt"

func swap(s []int, index1 int, index2 int) {
	temp := s[index1]
	s[index1] = s[index2]
	s[index2] = temp
}

func reverse(s []int) {
	for i := 0; i <= len(s)/2; i++ {
		swap(s, i, len(s)-i-1)
	}
}

func printSlice(s []int) {
	for _, value := range s {
		fmt.Printf("%d, ", value)
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
	s := initSlice(5)
	fmt.Printf("Original slice of size 5:\n")
	printSlice(s)
	reverse(s)
	fmt.Printf("Reversed slice of size 5:\n")
	printSlice(s)

	s = initSlice(21)
	fmt.Printf("Original slice of size 21:\n")
	printSlice(s)
	reverse(s)
	fmt.Printf("Reversed slice of size 21:\n")
	printSlice(s)
}
