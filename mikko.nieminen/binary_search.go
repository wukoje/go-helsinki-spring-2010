package main

import (
	"fmt"
	)

const (
	SIZE = 50
	DEBUG = false
)

func main() {

	slice := make([]int, SIZE) // create a new slice and initialize it
	for i, _ := range slice {
		slice[i]=i*2;
	}

	print_slice(slice)

	test(slice, -2)
	test(slice, -1)
	test(slice, 0)
	test(slice, 1)
	test(slice, 2)
	test(slice, 15)
	test(slice, 23)
	test(slice, 40)
	test(slice, 43)
	test(slice, 54)
	test(slice, 82)
	test(slice, 98)
	test(slice, 99)
	test(slice, 2000)
	test(slice, 50000)

}

func test(slice []int, findMe int) {
	x := search(slice, findMe)
	if x<0 {
		fmt.Printf ("element %d not found from the given slice\n", findMe)
	} else {
		fmt.Printf ("element %d found from index %d\n", findMe, x)
	}
}

func search(s []int, e int) int {

	lowerIndex := 0
	upperIndex := len(s)-1
	var middle int

	for lowerIndex < upperIndex {
	        middle = (upperIndex+lowerIndex)/2
		if DEBUG {
			fmt.Printf ("lower=%d, middle=%d, upper=%d\n", lowerIndex, middle, upperIndex)
		}
	        if (e < s[middle]) {
	            upperIndex = middle
	        } else if (e > s[middle]) {
	            lowerIndex = middle + 1
	        } else {
	            return middle
	        }
    	}

	return -1 // element e not found from slice s
}

func print_slice(slice []int) {
	fmt.Println ("Slice elements:")
	for i, c := range slice {
		fmt.Printf ("[%d]=%v\n", i, c)
	}	
}
