package main

import (
	"fmt"
)


func main() {
	s := []int{10,15,20,25,30,35}

	fmt.Printf("Slice1 before: ")
	for _, v := range s {
		fmt.Printf("%d ", v)
	} 

	fmt.Printf("\n")

	reverse(s)

        fmt.Printf("Slice1 after: ")
        for _, v := range s {
                fmt.Printf("%d ", v)
        }

        fmt.Printf("\n")

        s = []int{33,90,11,22,10101,88,43,23}

        fmt.Printf("Slice2 before: ")
        for _, v := range s {
                fmt.Printf("%d ", v)
        }

        fmt.Printf("\n")

        reverse(s)

        fmt.Printf("Slice2 after: ")
        for _, v := range s {
                fmt.Printf("%d ", v)
        }

        fmt.Printf("\n")




}


func reverse (s []int) {
	for i, j := len(s)-1, 0; i>j; i, j = i-1, j+1 {
		s[i], s[j] = s[j], s[i]
	} 
}
