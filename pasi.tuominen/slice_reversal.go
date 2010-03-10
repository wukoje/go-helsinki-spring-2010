package main

import (
	"fmt"
)

func reverse(s []int) {
	sliceLength := len(s)
	for i := 0; i < sliceLength/2; i++ {
		s[i], s[sliceLength-1-i] = s[sliceLength-1-i],s[i]
	}
}

func printSlice(s []int) {
	for i:=0; i<len(s);i++ {
		fmt.Printf("%d ", s[i]);
	}
	fmt.Printf("\n");
}

func main() {

	slice_1 := []int{1,2,3,4,5,6,7,8,9}
	slice_2 := []int{1,2,3,4,5,6,7,8}
	slice_3 := []int{}
	
	fmt.Printf("Slice_1:\n")
	printSlice(slice_1)
	reverse(slice_1)
	printSlice(slice_1)
	
	fmt.Printf("Slice_2:\n")
	printSlice(slice_2)
	reverse(slice_2)
	printSlice(slice_2)
		
	fmt.Printf("Slice_3:\n")
	printSlice(slice_3)
	reverse(slice_3)
	printSlice(slice_3)
	
}
