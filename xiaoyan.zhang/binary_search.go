/*
search returns the index of the element e in the sorted slice s.
If e does not exist in s, search returns -1.

Author: Xiaoyan ZHANG
*/

package main

import(
	"fmt"
)

func search(s []int, e int) int{
	first := 0
	last := len(s)-1

	for last >= first{
		middle := first+(last-first)/2
		current := s[middle]

		if current == e{
			return middle
		} else if current > e {
			last = middle - 1
		} else {
			first = middle + 1
		}
	}

	return -1
}

func main(){
	//Test case 1
	s_1 := []int{1,3,7,9,14,18}
	e_1 := 18

	fmt.Printf("Test case 1:\n\tfind %d in slice ", e_1)
	fmt.Println(s_1)
	fmt.Printf("Result: %d\n\n", search(s_1, e_1))

	//Test case 2
	s_2 := []int{-3,-1,0,8,19}
	e_2 := -1

	fmt.Printf("Test case 2:\n\tfind %d in slice ", e_2)
	fmt.Println(s_2)
	fmt.Printf("Result: %d\n\n", search(s_2, e_2))

	//Test case 3
	s_3 := []int{7,12,94}
	e_3 := 2

	fmt.Printf("Test case 3:\n\tfind %d in slice ", e_3)
	fmt.Println(s_3)
	fmt.Printf("Result: %d\n\n", search(s_3, e_3))

}



