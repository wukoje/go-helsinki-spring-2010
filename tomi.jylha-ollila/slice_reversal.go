package main


import (
	"fmt"
	"rand"
)


func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}


func main() {
	s := make([]int, 10)
	for i, _ := range s {
		s[i] = i * i
	}
	fmt.Printf("%v\n", s)
	reverse(s)
	fmt.Printf("%v\n", s)
	s = make([]int, 17)
	for i, _ := range s {
		s[i] = int(rand.Float() * 100)
	}
	fmt.Printf("%v\n", s)
	reverse(s)
	fmt.Printf("%v\n", s)
}
