package main

import (
	"fmt"
)


func sort(s []int) {
	quicksort(s, 0, len(s)-1,nil)
}

func quicksort(s []int, low int, high int, c chan int) {

	if (high > low) {
		pivot := partition(s, low, high)
		
		ca := make(chan int)
		go quicksort(s, low, pivot-1, ca)
		quicksort(s, pivot+1, high, nil)
		<-ca
		
	}
	if (c != nil) {
		c <- 1
	}
}

func partition(s []int, low int, high int) int {
	
	pi := low + (high-low)/2 
	pv := s[pi]
	s[pi],s[high] = s[high],s[pi]
	si := low
	for i:=low;i < high;i++ {
	
		if (s[i] < pv) {
			s[i], s[si] = s[si], s[i]
			si++
		}
	}
	s[si],s[high] = s[high],s[si]
	return si
}

func printSlice(s []int) {
	for i:=0; i<len(s);i++ {
		fmt.Printf("%d ", s[i]);
	}
	fmt.Printf("\n");
}

func main() {

	slice1 := []int{2,7,3,8,2,7,4,6,8,8}
	slice2 := []int{1,2,3,4,5,6,7,8,9}
	slice3 := []int{2,1}
	slice4 := []int{}
	slice5 := []int{1}
	
	fmt.Printf("\nUnsorted: ")
	printSlice(slice1)
	sort(slice1)
	fmt.Printf("Sorted:   ")
	printSlice(slice1)
	
	fmt.Printf("\nUnsorted: ")
	printSlice(slice2)
	sort(slice2)
	fmt.Printf("Sorted:   ")
	printSlice(slice2)

	fmt.Printf("\nUnsorted: ")
	printSlice(slice3)
	sort(slice3)
	fmt.Printf("Sorted:   ")
	printSlice(slice3)
	
	fmt.Printf("\nUnsorted: ")
	printSlice(slice4)
	sort(slice4)
	fmt.Printf("Sorted:   ")
	printSlice(slice4)
	
	fmt.Printf("\nUnsorted: ")
	printSlice(slice5)
	sort(slice5)
	fmt.Printf("Sorted:   ")
	printSlice(slice5)
	
}
