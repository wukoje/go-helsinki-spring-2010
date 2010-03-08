package main

import "fmt"

func main() {
	var l = [...]int {1,2,3,4,5,6,7,8,9,10,11}
	fmt.Printf("Before reversing:\t%v\n", l)
	reverse(&l)
	fmt.Printf("After reversing:\t%v\n", l)
}

func reverse(s []int) {
	for i:=0; i<len(s)/2; i++ {
		s[i],s[len(s)-i-1] = s[len(s)-i-1],s[i]
	}
}

