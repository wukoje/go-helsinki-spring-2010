/*
1 slice_reversal.go

// reverse reverses the contents of s in place
func reverse(s []int)

Output some slices before and after reversal.
*/

package main

import (
	"fmt"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}


func main() {
	slice := []int{0,1,2,3,4,5,6,7,8,9,10,10,9,8,7,6,5,4,3,2,1,0};
	fmt.Printf("The whole slice: %v\n", slice);
	fmt.Printf("Let's reverse from 0 to 11 - %v\n", slice[0:11]);
	reverse(slice[0:11]);
	fmt.Printf("Slice: %v\n", slice);
	fmt.Printf("Then reverse the rest: %v\n", slice[11:]);
	reverse(slice[11:]);
	fmt.Printf("Slice: %v\n", slice);
	fmt.Printf("Reversing the whole slice should have no effect\n");
	reverse(slice);
	fmt.Printf("Slice: %v\n", slice);

	oddSlice := []int{1,2,3};
	fmt.Printf("\nNow we'll reverse a slice with odd number of elements..\n");
	fmt.Printf("Slice: %v\n", oddSlice);
	reverse(oddSlice);
	fmt.Printf("Reversed: %v\n", oddSlice);

	fmt.Printf("\nHow about an empty one\n");
	empty := []int{};
	fmt.Printf("Empty: %v\n", empty);
	reverse(empty);
	fmt.Printf("Reversed: %v\n", empty)
}
