package main

import "fmt"

func outputSlice(s []int) {
	for _, value := range s {
		fmt.Printf("%d, ", value)
	}
	fmt.Printf("\n")
}

func makeSlice(length int) []int {
	newSlice := make([]int, length)
	for i, _ := range newSlice {
		newSlice[i] = i
	}
	return newSlice
}
func swap(s []int, index1 int, index2 int) {
	temp := s[index1]
	s[index1] = s[index2]
	s[index2] = temp
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		swap(s, i, len(s)-i-1)
	}
}
func main(){

s := makeSlice(10)
	fmt.Printf("Original slice of size 10 :\n")
	outputSlice(s)
	reverse(s)
	fmt.Printf("reversed slice of size 10 :\n")
	outputSlice(s)
s1 := makeSlice(33)
	fmt.Printf("Original slice of size 33 :\n")
	outputSlice(s1)
	reverse(s1)
	fmt.Printf("reversed slice of size 33 :\n")
	outputSlice(s1)
}




