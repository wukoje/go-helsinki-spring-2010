/*
Ex2
reverse the contents of s in place

Author: Xiaoyan ZHANG
*/

package main

import(
	"fmt"
)

func reverse(s []int){
	size := len(s)

	for i:=0; i<(size/2); i++ {
		s[i], s[size-i-1] = s[size-i-1], s[i]
	}
}

func main(){
	//test_case 1
	fmt.Println("Test case 1:")
	num_1 := []int{3,4,21,7,6,43,63,87,15}
	fmt.Print("Original slice:")
	fmt.Println(num_1)

	reverse(num_1)
	fmt.Print("After reversal:")
	fmt.Println(num_1)

	//test_case 2
	fmt.Println("Test case 2:")
	num_2 := []int{6}
	fmt.Print("Original slice:")
	fmt.Println(num_2)

	reverse(num_2)
	fmt.Print("After reversal:")
	fmt.Println(num_2)

	//test_case 3
	fmt.Println("Test case 3:")
	num_3 := []int{8,394,48,59,2,65}
	fmt.Print("Original slice:")
	fmt.Println(num_3)

	reverse(num_3)
	fmt.Print("After reversal:")
	fmt.Println(num_3)
}
