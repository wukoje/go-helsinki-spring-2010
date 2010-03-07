package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		divisible := false

		if i%3 == 0 {
			fmt.Printf("Fizz")
			divisible = true
		}
		if i%5 == 0 {
			fmt.Printf("Buzz")
			divisible = true
		}
		if divisible == false {
			fmt.Printf("%d", i)
		}
		fmt.Printf("\n")
	}
}
