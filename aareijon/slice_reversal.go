package main

import fmt "fmt"

func reverse(s []int) {
	last := len(s)
	for i := 0; i < len(s)/2; i++ {
		s[i],s[last-1-i] = s[last-1-i], s[i]
	}
}

func main() {
	var s = [...]int{1,2,3,4,5}

	fmt.Println(s)
	reverse(&s)
	fmt.Println(s)
}
