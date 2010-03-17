package main

import (
	"fmt"
)


func main() {
	s := []int{105,1235,220,625,930,5,464,566,910,921,1200,1701,1756,124157,19,21}

	fmt.Printf("Test slice: ")
	for _, v := range s {
		fmt.Printf("%d ", v)
	} 

	sort(s)

	fmt.Printf("\nSorted slice: ")
	for _, v := range s {
		fmt.Printf("%d ", v)
	}
}


func sort (s []int) {
	quicksort(s, 0, len(s)-1)
}

func quicksort(s []int, low int, high int) {
	i := low
	j := high
	pivot := s[(low+high)/2]

	for i <= j {

		//look for swappable pair
		for s[i] < pivot {
			i++
		}
		for s[j] > pivot {
			j--
		}
		
		//swap, increase j & i		
		if i <= j {
			s[i], s[j] = s[j], s[i]	
			i++
			j--
		}
	}
	
	//recursion
	if low < j {
		quicksort(s, low, j)
	}
	if i < high {
		quicksort(s, i, high)
	}
}
