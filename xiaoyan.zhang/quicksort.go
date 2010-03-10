/*
Ex 4
sort modifies the slice s
so that the integers are sorted in place using quicksort

Author: Xiaoyan ZHANG
*/

package main

import(
	"fmt"
)

func sort(s []int){
	low := 0
	high := len(s)-1
	pivot := s[0]

	for low != high {
		if s[high] < pivot{
			s[low] = s[high]
			low++
		} else {
			high--
			continue
		}
		
		for low != high {
			if s[low] >= pivot{
				s[high] = s[low]
				high--
				break
			} else {
				low++
			}
		}
	}

	s[low] = pivot

	if low > 1 {
		sort(s[0:low])
	}
	if high < len(s)-2 {
		sort(s[high+1:len(s)])
	}
}

func main(){
	//Test case 1 - normal case
	s := []int{3,2,16,4,12,8}
	fmt.Print("\nTest case 1: \noriginal slice: ")
	fmt.Println(s)
	fmt.Print("after quicksort: ")
	sort(s)
	fmt.Println(s)

	//Test case 2 - only one element
	s = []int{3}
	fmt.Print("\nTest case 2: \noriginal slice: ")
	fmt.Println(s)
	fmt.Print("after quicksort: ")
	sort(s)
	fmt.Println(s)

	//Test case 3 - case including duplicated element
	s = []int{28,23,8,34,7,9,23,4}
	fmt.Print("\nTest case 3: \noriginal slice: ")
	fmt.Println(s)
	fmt.Print("after quicksort: ")
	sort(s)
	fmt.Println(s)
}


