package main

// Algorithm referred to http://en.wikipedia.org/wiki/Quicksort

import "fmt"

func main() {
	var l = [...]int {4,5,8,9,11,6,7,3,1,2,10}
	fmt.Printf("Before sorting:\t%v\n", l)
	sort(&l, 0, len(l)-1)
	fmt.Printf("After sorting:\t%v\n", l)
}

func sort(s []int, left int, right int) {
     if right > left {
         pivotIndex := (left+right)/2
         pivotNewIndex := partition(s, left, right, pivotIndex)
         sort(s, left, pivotNewIndex - 1)
         sort(s, pivotNewIndex + 1, right)
	}
}

func partition(s []int, left int, right int, pivotIndex int) int {
	pivotValue := s[pivotIndex]
	s[pivotIndex], s[right] = s[right], s[pivotIndex]
	storeIndex := left
	for i:=left; i<right; i++ {
         if s[i] <= pivotValue {
             s[i], s[storeIndex] = s[storeIndex], s[i]
             storeIndex = storeIndex + 1
		}
	}
	s[storeIndex], s[right] = s[right], s[storeIndex]
	return storeIndex
}

