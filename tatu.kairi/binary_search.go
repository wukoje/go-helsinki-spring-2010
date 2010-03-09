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
	"rand"
	"sort"
)


// search returns the index of the element e in the
// sorted slice s. If e does not exist in s search returns -1
func search(s []int, e int) int {
	low := 0
	high := len(s) - 1

	for low <= high {
		mid := low + ((high - low) / 2)

		if s[mid] > e {
				high = mid - 1
		} else if s[mid] < e {
			low = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

// do 3 random slices and reverse them
// print the slice before and after reversing
func main(){

	for i := 0; i < 3; i++ {
		arr := randomSlice()

		searchable := arr[rand.Intn(len(arr))]

		printResults(arr, searchable, search(arr, searchable))
		fmt.Printf("---\n")

	}

	arr := randomSlice()

	printResults(arr, 99999, search(arr, 99999))

}

// print elements in the slice given as parameter and then a newline
func printResults(arr []int, searchable int, index int){
	fmt.Printf("searching: %d\nfound it in index: %d\n", searchable, index)
	fmt.Printf("array: ")
	for _,el := range(arr) {
		fmt.Printf("%d,", el)
	}
	fmt.Printf("\n")

}

// create slice with length/capacity between 1-51 populated with random numbers between 0-100
func randomSlice() []int {
	temp := make([]int, (rand.Intn(50) + 1)) // rand.Intn can return 0

	for i,_ := range(temp) {
		temp[i] = rand.Intn(100)
	}

	sort.SortInts(temp)
	return temp
}
