package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		printTheNumber := true
		if i % 3 == 0 {
			fmt.Printf("Fizz")
			printTheNumber = false
		}
		if i % 5 == 0 {
			fmt.Printf("Buzz")
			printTheNumber = false
		}
		if printTheNumber {
			fmt.Printf("%d\t", i)
		}
		fmt.Printf("\n")
	}
}

