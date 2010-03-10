/*
1 slice_reversal.go

// reverse reverses the contents of s in place
func reverse(s []int)

Output some slices before and after reversal.
*/

package main

import (
	"fmt"
	"sort"
	"rand"
)

// reverse reverses the contents of s in place
func reverse(s []int){
	j := len(s) - 1

  /* TODO: can't figure out how to pass two variables in the for-statement
			eg. for i := 0, j := 0; i < j; i++, j-- { ...
	*/
	for i := 0; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}

// do 10 random slices and reverse them
// print the slice before and after reversing
func main(){

	for i := 0; i < 10; i++ {
		arr := randomSlice()

		fmt.Printf("before: %v\n", arr)
		reverse(arr)
		fmt.Printf("\n")
		fmt.Printf("after: %v\n", arr)
		fmt.Printf("---\n")

	}
}

// create slice with length/capacity between 1-51 populated with random numbers between 0-100
func randomSlice() []int {
	temp := make([]int, (rand.Intn(50) + 1)) // rand.Intn can return 0

	for i,_ := range temp{
		temp[i] = rand.Intn(100)
	}

	sort.SortInts(temp)
	return temp
}
