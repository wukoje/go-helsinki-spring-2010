package main

import (
	"fmt"
)

func search(s []int, e int) int {
	return binarySearch(s,e,0,len(s)-1)
}

func binarySearch(s []int, e int, low int, high int) int {
	if (high < low) {
		return -1;
	}
	mid := low + ((high-low)/2)
	if (s[mid] > e) {
		return binarySearch(s,e,low,mid-1)
	} else if (s[mid] < e) {

		return binarySearch(s,e,mid+1,high)	
	}
	return mid
	
}
func printSlice(s []int) {
	for i:=0; i<len(s);i++ {
		fmt.Printf("%d ", s[i]);
	}
	fmt.Printf("\n");
}

func main() {

	slice := []int{1,2,3,4,5,6,7,8,9}

	fmt.Printf("The slice is:\n");
	printSlice(slice);
	
	e := 1
	fmt.Printf("Index of %d is: %d\n", e, search(slice,e));

	e = 5
	fmt.Printf("Index of %d is: %d\n", e, search(slice,e));

	e = 9
	fmt.Printf("Index of %d is: %d\n", e, search(slice,e))

	e = 15
	fmt.Printf("Index of %d is: %d\n", e, search(slice,e))	

	
}
