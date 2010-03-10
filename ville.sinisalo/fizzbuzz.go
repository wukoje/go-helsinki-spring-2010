package main

import "fmt"

func main() {
	for i:=1; i<=100; i++ {
		b := true
		if i%3 == 0 {
			fmt.Printf("Fizz")
			b = false
		}
		if i%5 == 0 {
			fmt.Printf("Buzz")
			b = false
		}
		if b {
			fmt.Printf("%d", i)
		}
		fmt.Printf("\n")
	}
}
