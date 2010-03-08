package main

import (
	"fmt"
)

func main() {
	for i := 1; i<=100; i++ {
		x := 0
		if  i%3 == 0 {
			fmt.Printf("Fizz")
			x++
		}
		if  i%5 == 0 {
			fmt.Printf("Buzz")
			x++
		}
		if x==0 {
			fmt.Printf("%d ", i)
		}

		fmt.Printf("\n")
	}

}
