/*
1 binary_search.go

// search returns the index of the element e in the
// sorted slice s. If e does not exist in s search returns -1
func search(s []int, e int) int

Generate some test data and output a few slices and search results.
*/

package main

import (
	"fmt"
)

func search(s []int, e int) int {
	beg := 0;
	end := len(s)-1;

	for beg <= end {
		mid := beg + ((end-beg) / 2); // overflow guard
		if s[mid] == e {
			return mid
		} else if s[mid] > e {
			end = mid - 1
		} else {
			beg = mid + 1
		}
	}
	return -1
}

func testHuge(shift uint) {
	max := int((1 << shift) - 1);
	fmt.Printf("Testing with a huuuge slice (sw: %d size: %v)...\n",
					shift, max);
	huge := make([]int, max);
	fmt.Printf("Managed to create huge slice\n");
	for i := 0; i < max; i++ {
		huge[i] = i;
	}
	fmt.Printf("Managed to fill huge slice\n");
	if ret := search(huge, max-1); ret < 0 {
		fmt.Printf("ERROR: %d not found\n", max-1)
	} else if huge[ret] != max-1 {
		fmt.Printf("ERROR: found %d from idx %d, but wanted %d\n",
					huge[ret], ret, max-1)
	} else {
		fmt.Printf("Huge slice - passed\n")
	}

}

func main() {
	s := []int{1, 3, 6, 9, 12, 15, 23, 45, 567, 666, 777, 1024};
	fmt.Printf("Search space: %v\n", s);
	for i := 0; i < len(s); i++ {
		if ret := search(s, s[i]); ret < 0 {
			fmt.Printf("ERROR: %d not found\n", s[i])
		} else if ret != i {
			fmt.Printf("Oops, %d from idx %d found from %d containing %d\n",
					s[i], i, ret, s[ret])
		} else {
			fmt.Printf("Index %d - passed\n", i);
		}
	}
	if ret := search(s, -666); ret >= 0 {
		fmt.Printf("ERROR: Found -666 from idx %d\n", ret)
	} else {
		fmt.Printf("Not found - passed\n")
	}
	testHuge(29);
	//testHuge(31); // out of mem (x86)
}

