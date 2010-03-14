package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		d3 := i%3 == 0
		d5 := i%5 == 0
		if d3 {
			fmt.Print("Fizz")
		}
		if d5 {
			fmt.Print("Buzz")
		}
		if !d3 && !d5 {
			fmt.Printf("%d", i)
		}
		fmt.Println("")
	}
}
