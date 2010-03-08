package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fizz := i % 3 == 0
		buzz := i % 5 == 0
		if fizz || buzz {
			if fizz {
				fmt.Printf("Fizz")
			}
			if buzz {
				fmt.Printf("Buzz")
			}
			fmt.Printf("\n")
		} else {
			fmt.Printf("%v\n", i)
		}
	}
}
