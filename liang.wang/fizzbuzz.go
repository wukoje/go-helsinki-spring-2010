package main

import "fmt"

func main() {
	bHit := false
	for i:=1; i<=100; i++ {
		if i%3 == 0 {
			fmt.Printf("Fizz")
			bHit = true
		}
		if i%5 == 0 {
			fmt.Printf("Buzz")
			bHit = true
		}
		if !bHit {
			fmt.Printf("%v", i)
		}
		bHit = false
		fmt.Printf("\n")
	}
}

