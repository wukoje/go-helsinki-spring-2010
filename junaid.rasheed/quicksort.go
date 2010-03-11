package main

import "fmt"

func sort(s []int) {
	left := 0
	right := len(s) - 1
	quicksort(s, left, right)
}

func quicksort(s []int, left int, right int) {

	if right > left {
		pivotIndex := (left + right) / 2
		pivotNewIndex := partition(s, left, right, pivotIndex)
		quicksort(s, left, pivotNewIndex-1)
		quicksort(s, pivotNewIndex+1, right)
	}

}

func partition(s []int, left int, right int, pivotIndex int) int {

	pivotValue := s[pivotIndex]
	temp := s[pivotIndex]
	s[pivotIndex] = s[right] // Move pivot to end
	s[right] = temp
	storeIndex := left
	for i := left; i < right; i++ {
		if s[i] <= pivotValue {
			temp1 := s[i]
			s[i] = s[storeIndex]
			s[storeIndex] = temp1
			storeIndex = storeIndex + 1
		}
	}
	temp3 := s[storeIndex]
	s[storeIndex] = s[right] // Move pivot to its final place
	s[right] = temp3
	return storeIndex

}

func main() {

	s := []int{10, 9, 8, 5, 7, 2, 1, 6}

	fmt.Printf("Before sort: ")
	fmt.Print(s)
	sort(s)
	fmt.Printf("\nAfter sort: ")
	fmt.Print(s)
	fmt.Printf("\n")

	s1 := []int{2, 3, 1, 4, 8, 5, 33, 23, 88, 2, 4, 6, 9, 5, 3, 7, 8, 3, 1, 6, 7, 4, 7}

	fmt.Printf("Before sort: ")
	fmt.Print(s1)
	sort(s1)
	fmt.Printf("\nAfter sort: ")
	fmt.Print(s1)
	fmt.Printf("\n")

}

