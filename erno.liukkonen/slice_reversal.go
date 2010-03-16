package main

import fmt "fmt"

func main() {

	var test =[...]int{10,9,8,7,6,5,4,3,2,1,0}
	fmt.Println(test)
	reverse(&test)
	fmt.Println(test)
}

func reverse(s []int) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		s[i], s[length-i-1] = s[length-i-1], s[i]
	}
}


